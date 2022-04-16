package vm

// Code generated automatically; DO NOT EDIT

const (
	opret                    bcop = 0
	opjz                     bcop = 1
	oploadk                  bcop = 2
	opsavek                  bcop = 3
	opxchgk                  bcop = 4
	oploadb                  bcop = 5
	opsaveb                  bcop = 6
	oploadv                  bcop = 7
	opsavev                  bcop = 8
	oploadzerov              bcop = 9
	opsavezerov              bcop = 10
	oploadpermzerov          bcop = 11
	opsaveblendv             bcop = 12
	oploads                  bcop = 13
	opsaves                  bcop = 14
	oploadzeros              bcop = 15
	opsavezeros              bcop = 16
	opfalse                  bcop = 17
	opandk                   bcop = 18
	opork                    bcop = 19
	opandnotk                bcop = 20
	opnandk                  bcop = 21
	opxork                   bcop = 22
	opnotk                   bcop = 23
	opxnork                  bcop = 24
	opbroadcastimmf          bcop = 25
	opbroadcastimmi          bcop = 26
	opabsf                   bcop = 27
	opabsi                   bcop = 28
	opnegf                   bcop = 29
	opnegi                   bcop = 30
	opsignf                  bcop = 31
	opsigni                  bcop = 32
	opsquaref                bcop = 33
	opsquarei                bcop = 34
	opsqrtf                  bcop = 35
	oproundf                 bcop = 36
	oproundevenf             bcop = 37
	optruncf                 bcop = 38
	opfloorf                 bcop = 39
	opceilf                  bcop = 40
	opaddf                   bcop = 41
	opaddimmf                bcop = 42
	opaddi                   bcop = 43
	opaddimmi                bcop = 44
	opsubf                   bcop = 45
	opsubimmf                bcop = 46
	opsubi                   bcop = 47
	opsubimmi                bcop = 48
	oprsubf                  bcop = 49
	oprsubimmf               bcop = 50
	oprsubi                  bcop = 51
	oprsubimmi               bcop = 52
	opmulf                   bcop = 53
	opmulimmf                bcop = 54
	opmuli                   bcop = 55
	opmulimmi                bcop = 56
	opdivf                   bcop = 57
	opdivimmf                bcop = 58
	oprdivf                  bcop = 59
	oprdivimmf               bcop = 60
	opdivi                   bcop = 61
	opdivimmi                bcop = 62
	oprdivi                  bcop = 63
	oprdivimmi               bcop = 64
	opmodf                   bcop = 65
	opmodimmf                bcop = 66
	oprmodf                  bcop = 67
	oprmodimmf               bcop = 68
	opmodi                   bcop = 69
	opmodimmi                bcop = 70
	oprmodi                  bcop = 71
	oprmodimmi               bcop = 72
	opaddmulimmi             bcop = 73
	opminvaluef              bcop = 74
	opminvalueimmf           bcop = 75
	opmaxvaluef              bcop = 76
	opmaxvalueimmf           bcop = 77
	opminvaluei              bcop = 78
	opminvalueimmi           bcop = 79
	opmaxvaluei              bcop = 80
	opmaxvalueimmi           bcop = 81
	opcvtktof64              bcop = 82
	opcvtktoi64              bcop = 83
	opcvti64tof64            bcop = 84
	opcvtf64toi64            bcop = 85
	opfproundu               bcop = 86
	opfproundd               bcop = 87
	opcvti64tostr            bcop = 88
	opcmpeqf                 bcop = 89
	opcmpeqi                 bcop = 90
	opcmpeqimmf              bcop = 91
	opcmpeqimmi              bcop = 92
	opcmpltf                 bcop = 93
	opcmplti                 bcop = 94
	opcmpltimmf              bcop = 95
	opcmpltimmi              bcop = 96
	opcmplef                 bcop = 97
	opcmplei                 bcop = 98
	opcmpleimmf              bcop = 99
	opcmpleimmi              bcop = 100
	opcmpgtf                 bcop = 101
	opcmpgti                 bcop = 102
	opcmpgtimmf              bcop = 103
	opcmpgtimmi              bcop = 104
	opcmpgef                 bcop = 105
	opcmpgei                 bcop = 106
	opcmpgeimmf              bcop = 107
	opcmpgeimmi              bcop = 108
	opisnanf                 bcop = 109
	opchecktag               bcop = 110
	opisnull                 bcop = 111
	opisnotnull              bcop = 112
	opistrue                 bcop = 113
	opisfalse                bcop = 114
	opeqslice                bcop = 115
	opequalv                 bcop = 116
	opeqv4mask               bcop = 117
	opeqv4maskplus           bcop = 118
	opeqv8                   bcop = 119
	opeqv8plus               bcop = 120
	opleneq                  bcop = 121
	opdateaddmonth           bcop = 122
	opdateaddmonthimm        bcop = 123
	opdateaddyear            bcop = 124
	opdatediffparam          bcop = 125
	opdatediffmonthyear      bcop = 126
	opdateextractmicrosecond bcop = 127
	opdateextractmillisecond bcop = 128
	opdateextractsecond      bcop = 129
	opdateextractminute      bcop = 130
	opdateextracthour        bcop = 131
	opdateextractday         bcop = 132
	opdateextractmonth       bcop = 133
	opdateextractyear        bcop = 134
	opdatetounixepoch        bcop = 135
	opdatetruncmillisecond   bcop = 136
	opdatetruncsecond        bcop = 137
	opdatetruncminute        bcop = 138
	opdatetrunchour          bcop = 139
	opdatetruncday           bcop = 140
	opdatetruncmonth         bcop = 141
	opdatetruncyear          bcop = 142
	opunboxts                bcop = 143
	opboxts                  bcop = 144
	optimelt                 bcop = 145
	optimegt                 bcop = 146
	opconsttm                bcop = 147
	optmextract              bcop = 148
	opwidthbucketf           bcop = 149
	opwidthbucketi           bcop = 150
	optimebucketts           bcop = 151
	opgeogridi               bcop = 152
	opgeogridimmi            bcop = 153
	opgeohash                bcop = 154
	opgeohashimm             bcop = 155
	opfindsym                bcop = 156
	opfindsym2               bcop = 157
	opfindsym2rev            bcop = 158
	opfindsym3               bcop = 159
	opblendv                 bcop = 160
	opblendrevv              bcop = 161
	opblendnum               bcop = 162
	opblendnumrev            bcop = 163
	opblendslice             bcop = 164
	opblendslicerev          bcop = 165
	opunpack                 bcop = 166
	optoint                  bcop = 167
	optof64                  bcop = 168
	opboxfloat               bcop = 169
	opboxint                 bcop = 170
	opboxmask                bcop = 171
	opboxmask2               bcop = 172
	opboxmask3               bcop = 173
	opboxstring              bcop = 174
	ophashvalue              bcop = 175
	ophashvalueplus          bcop = 176
	ophashmember             bcop = 177
	ophashlookup             bcop = 178
	opaggsumf                bcop = 179
	opaggsumi                bcop = 180
	opaggminf                bcop = 181
	opaggmini                bcop = 182
	opaggmaxf                bcop = 183
	opaggmaxi                bcop = 184
	opaggcount               bcop = 185
	opaggbucket              bcop = 186
	opaggslotaddf            bcop = 187
	opaggslotaddi            bcop = 188
	opaggslotavgf            bcop = 189
	opaggslotavgi            bcop = 190
	opaggslotminf            bcop = 191
	opaggslotmini            bcop = 192
	opaggslotmaxf            bcop = 193
	opaggslotmaxi            bcop = 194
	opaggslotcount           bcop = 195
	oplitref                 bcop = 196
	opsplit                  bcop = 197
	optuple                  bcop = 198
	opdupv                   bcop = 199
	opzerov                  bcop = 200
	opobjectsize             bcop = 201
	opCmpStrEqCs             bcop = 202
	opCmpStrEqCi             bcop = 203
	opCmpStrEqUTF8Ci         bcop = 204
	opSkip1charLeft          bcop = 205
	opSkip1charRight         bcop = 206
	opSkipNcharLeft          bcop = 207
	opSkipNcharRight         bcop = 208
	opTrimWsLeft             bcop = 209
	opTrimWsRight            bcop = 210
	opTrim4charLeft          bcop = 211
	opTrim4charRight         bcop = 212
	opTrimPrefixCs           bcop = 213
	opTrimPrefixCi           bcop = 214
	opTrimSuffixCs           bcop = 215
	opTrimSuffixCi           bcop = 216
	opContainsSubstrCs       bcop = 217
	opContainsSubstrCi       bcop = 218
	opContainsSuffixCs       bcop = 219
	opContainsSuffixCi       bcop = 220
	opContainsSuffixUTF8Ci   bcop = 221
	opContainsPrefixCs       bcop = 222
	opContainsPrefixCi       bcop = 223
	opContainsPrefixUTF8Ci   bcop = 224
	opLengthStr              bcop = 225
	opSubstr                 bcop = 226
	opSplitPart              bcop = 227
	opMatchpatCs             bcop = 228
	opMatchpatCi             bcop = 229
	opMatchpatUTF8Ci         bcop = 230
	opIsSubnetOfIP4          bcop = 231
	optrap                   bcop = 232
	_maxbcop                      = 233
)
