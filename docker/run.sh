#!/bin/bash
echo -n "Checking for AVX-512 support..."
if grep -q avx512 /proc/cpuinfo; then
    echo "supported"
else
    echo "not supported"
    exit 1
fi

echo -n "Checking for CURL..."
if [ -x "$(command -v curl)" ]; then
    echo "found"
else
    echo "not found (please install CURL)"
    exit 2
fi

echo -n "Checking for Docker installation..."
if [ -x "$(command -v docker)" ]; then
    echo "found"
else
    echo "not found (please install Docker)"
    exit 3
fi

echo -n "Checking for Docker Compose installation..."
if [ -x "$(command -v docker-compose)" ]; then
    echo "found"
else
    echo "not found (please install Docker Compose)"
    exit 4
fi

echo -n "Generating secrets..."
./generate-env.sh > /dev/null
. ./.env
echo "done"

docker-compose up -d

# Create the test bucket
if docker run \
    --rm --net sneller-network --env-file .env \
    amazon/aws-cli --endpoint http://minio:9100 \
    s3 mb "s3://test"
then
    echo "Test bucket created"
else
    exit 5
fi

# Download the customer data
TEMPFILE=$(mktemp)
curl --output $TEMPFILE 'https://sneller-example-data.s3.amazonaws.com/docker/sf1/customer.json'

# Copy the customer data to the test bucket
if docker run \
    --rm --net sneller-network --env-file .env -v "$TEMPFILE:/data/customer.json" \
    amazon/aws-cli --endpoint http://minio:9100 \
    s3 cp /data/customer.json s3://test/sf1/customer.json
then
    rm $TEMPFILE
else
    rm $TEMPFILE
    exit 6
fi

# Create table definition in Minio bucket
echo '{"name": "customer", "input": [{"pattern": "s3://test/sf1/*.json","format": "json"}]}' > $TEMPFILE
if docker run \
    --rm --net sneller-network --env-file .env -v "$TEMPFILE:/data/definition.json" \
    amazon/aws-cli --endpoint http://minio:9100 \
    s3 cp /data/definition.json s3://test/db/sf1/customer/definition.json
then
    rm $TEMPFILE
else
    rm $TEMPFILE
    exit 7
fi

# Ingest data
if docker run \
   --rm --net sneller-network --env-file .env \
   ${SNELLER_REPO}sneller/sdb \
   -v sync sf1 customer
then
    echo "Data ingested"
else
    exit 8
fi

# Query the number of items
curl -G -H "Authorization: Bearer $SNELLER_TOKEN" \
    --data-urlencode "database=sf1" \
    --data-urlencode 'json' \
    --data-urlencode 'query=SELECT COUNT(*) FROM customer' \
    'http://localhost:9180/executeQuery'
