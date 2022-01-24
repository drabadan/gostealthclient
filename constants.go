package gostealthclient

var (
	debug = false
)

var increment uint16 = 0

const READ_DELAY = 0 // time.Microsecond * 50

const (
	SCZero                             = 0
	SCReturnValue                      = 1
	SCScriptDLLTerminate               = 2
	SCExecEventProc                    = 3
	SCPauseResumeScript                = 4
	SCLangVersion                      = 5
	SCErrorReport                      = 6
	SCClearEventProc                   = 7
	SCGetProfileName                   = 8
	SCGetConnectedStatus               = 9
	SCAddToSystemJournal               = 10
	SCSetEventProc                     = 11
	SCGetStealthInfo                   = 12
	SCGetCell                          = 13
	SCGetSelfID                        = 14
	SCGetX                             = 15
	SCGetY                             = 16
	SCGetZ                             = 17
	SCGetWorldNum                      = 18
	SCGetCharName                      = 19
	SCChangeProfile                    = 20
	SCGetARStatus                      = 21
	SCSetARStatus                      = 22
	SCGetPauseScriptOnDisconnectStatus = 23
	SCSetPauseScriptOnDisconnectStatus = 24
	SCGetSelfSex                       = 25
	SCGetCharTitle                     = 26
	SCGetSelfGold                      = 27
	SCGetSelfArmor                     = 28
	SCGetSelfWeight                    = 29
	SCGetSelfMaxWeight                 = 30
	SCGetSelfRace                      = 31
	SCGetSelfPetsMax                   = 32
	SCGetSelfPetsCurrent               = 33
	SCGetSelfFireResist                = 34
	SCGetSelfColdResist                = 35
	SCGetSelfPoisonResist              = 36
	SCGetSelfEnergyResist              = 37
	SCGetConnectedTime                 = 38
	SCGetDisconnectedTime              = 39
	SCGetLastContainer                 = 40
	SCGetLastTarget                    = 41
	SCGetLastAttack                    = 42
	SCGetLastStatus                    = 43
	SCGetLastObject                    = 44
	SCConnect                          = 45
	SCDisconnect                       = 46
	SCGetShardName                     = 47
	SCGetBackpackID                    = 48
	SCGetSelfStr                       = 49
	SCGetSelfInt                       = 50
	SCGetSelfDex                       = 51
	SCGetSelfLife                      = 52
	SCGetSelfMana                      = 53
	SCGetSelfStam                      = 54
	SCGetSelfMaxLife                   = 55
	SCGetSelfMaxMana                   = 56
	SCGetSelfMaxStam                   = 57
	SCGetSelfLuck                      = 58
	SCGetExtInfo                       = 59
	SCGetProxyIP                       = 60
	SCGetProxyPort                     = 61
	SCGetUseProxy                      = 62
	SCGetHiddenStatus                  = 63
	SCGetPoisonedStatus                = 64
	SCGetParalyzedStatus               = 65
	SCGetDeadStatus                    = 66
	SCGetWarTarget                     = 67
	SCSetWarMode                       = 68
	SCAttack                           = 69
	SCUseSelfPaperdollScroll           = 70
	SCUseOtherPaperdollScroll          = 71
	SCGetTargetID                      = 72
	SCCancelTarget                     = 73
	SCTargetToObject                   = 74
	SCTargetToXYZ                      = 75
	SCTargetToTile                     = 76
	SCWaitTargetObject                 = 77
	SCWaitTargetTile                   = 78
	SCWaitTargetXYZ                    = 79
	SCWaitTargetSelf                   = 80
	SCWaitTargetType                   = 81
	SCCancelWaitTarget                 = 82
	SCWaitTargetGround                 = 83
	SCWaitTargetLast                   = 84
	SCUsePrimaryAbility                = 85
	SCUseSecondaryAbility              = 86
	SCGetActiveAbility                 = 87
	SCToggleFly                        = 88
	SCGetSkillID                       = 89
	SCUseSkill                         = 90
	SCChangeSkillLockState             = 91
	SCGetSkillCap                      = 92
	SCSkillValue                       = 93
	SCReqVirtuesGump                   = 94
	SCUseVirtue                        = 95
	SCCastSpell                        = 96
	SCIsActiveSpellAbility             = 98
	SCSetCatchBag                      = 99
	SCUnsetCatchBag                    = 100
	SCUseObject                        = 101
	SCUseType                          = 102
	SCUseFromGround                    = 103
	SCClickOnObject                    = 104
	SCGetFoundedParamID                = 105
	SCGetLineID                        = 106
	SCGetLineType                      = 107
	SCGetLineTime                      = 108
	SCGetLineMsgType                   = 109
	SCGetLineTextColor                 = 110
	SCGetLineTextFont                  = 111
	SCGetLineIndex                     = 112
	SCGetLineCount                     = 113
	SCGetLineName                      = 114
	SCAddJournalIgnore                 = 115
	SCClearJournalIgnore               = 116
	SCUAddChatUserIgnore               = 117
	SCClearChatUserIgnore              = 118
	SCClearJournal                     = 119
	SCLastJournalMessage               = 120
	SCInJournal                        = 121
	SCInJournalBetweenTimes            = 122
	SCJournal                          = 123
	SCSetJournalLine                   = 124
	SCLowJournal                       = 125
	SCHighJournal                      = 126
	SCAddToJournal                     = 304
	SCSetFindDistance                  = 127
	SCGetFindDistance                  = 128
	SCSetFindVertical                  = 129
	SCGetFindVertical                  = 130
	SCFindTypeEx                       = 131
	SCFindNotoriety                    = 132
	SCFindAtCoord                      = 133
	SCIgnore                           = 134
	SCIgnoreOff                        = 135
	SCIgnoreReset                      = 136
	SCGetIgnoreList                    = 137
	SCGetFindedList                    = 138
	SCGetFindItem                      = 139
	SCGetFindCount                     = 140
	SCGetFindQuantity                  = 141
	SCGetFindFullQuantity              = 142
	SCPredictedX                       = 143
	SCPredictedY                       = 144
	SCPredictedZ                       = 145
	SCPredictedDirection               = 146
	SCGetName                          = 147
	SCGetAltName                       = 148
	SCGetTitle                         = 149
	SCGetCliloc                        = 150
	SCGetType                          = 151
	SCGetToolTipRec                    = 152
	SCGetClilocByID                    = 153
	SCGetQuantity                      = 154
	SCIsObjectExists                   = 155
	SCGetPrice                         = 156
	SCGetDirection                     = 157
	SCGetDistance                      = 158
	SCGetColor                         = 159
	SCGetStr                           = 160
	SCGetInt                           = 161
	SCGetDex                           = 162
	SCGetHP                            = 163
	SCGetMaxHP                         = 164
	SCGetMana                          = 165
	SCGetMaxMana                       = 166
	SCGetStam                          = 167
	SCGetMaxStam                       = 168
	SCGetNotoriety                     = 169
	SCGetParent                        = 170
	SCIsWarMode                        = 171
	SCIsNPC                            = 172
	SCIsDead                           = 173
	SCIsRunning                        = 174
	SCIsContainer                      = 175
	SCIsHidden                         = 176
	SCIsMovable                        = 177
	SCIsYellowHits                     = 178
	SCIsPoisoned                       = 179
	SCIsParalyzed                      = 180
	SCIsFemale                         = 181
	SCOpenDoor                         = 182
	SCBow                              = 183
	SCSalute                           = 184
	SCGetPickupedItem                  = 185
	SCSetPickupedItem                  = 186
	SCGetDropCheckCoord                = 187
	SCSetDropCheckCoord                = 188
	SCGetDropDelay                     = 189
	SCSetDropDelay                     = 190
	SCDragItem                         = 191
	SCDropItem                         = 192
	SCRequestContextMenu               = 193
	SCSetContextMenuHook               = 194
	SCGetContextMenu                   = 195
	SCClearContextMenu                 = 196
	SCCheckTradeState                  = 197
	SCGetTradeContainer                = 198
	SCGetTradeOpponent                 = 199
	SCGetTradeCount                    = 200
	SCGetTradeOpponentName             = 201
	SCTradeCheck                       = 202
	SCConfirmTrade                     = 203
	SCCancelTrade                      = 204
	SCWaitMenu                         = 205
	SCAutoMenu                         = 206
	SCMenuHookPresent                  = 207
	SCMenuPresent                      = 208
	SCCancelMenu                       = 209
	SCCloseMenu                        = 210
	SCWaitGumpInt                      = 211
	SCWaitGumpTextEntry                = 212
	SCGumpAutoTextEntry                = 213
	SCGumpAutoRadiobutton              = 214
	SCGumpAutoCheckBox                 = 215
	SCNumGumpButton                    = 216
	SCNumGumpTextEntry                 = 217
	SCNumGumpRadiobutton               = 218
	SCNumGumpCheckBox                  = 219
	SCGetGumpsCount                    = 220
	SCCloseSimpleGump                  = 221
	SCGetGumpSerial                    = 222
	SCGetGumpID                        = 223
	SCGetGumpNoClose                   = 224
	SCGetGumpTextLines                 = 225
	SCGetGumpFullLines                 = 226
	SCGetGumpShortLines                = 227
	SCGetGumpButtonsDescription        = 228
	SCGetGumpInfo                      = 229
	SCAddGumpIgnoreByID                = 230
	SCAddGumpIgnoreBySerial            = 231
	SCClearGumpsIgnore                 = 232
	SCObjAtLayerEx                     = 233
	SCGetLayer                         = 234
	SCWearItem                         = 235
	SCGetDressSpeed                    = 236
	SCSetDressSpeed                    = 237
	SCSetDress                         = 238
	SCGetDressSet                      = 239
	SCAutoBuy                          = 240
	SCGetShopList                      = 241
	SCClearShopList                    = 242
	SCAutoBuyEx                        = 243
	SCGetAutoBuyDelay                  = 244
	SCSetAutoBuyDelay                  = 245
	SCGetAutoSellDelay                 = 246
	SCSetAutoSellDelay                 = 247
	SCAutoSell                         = 248
	SCRequestStats                     = 249
	SCHelpRequest                      = 250
	SCQuestRequest                     = 251
	SCRenameMobile                     = 252
	SCMobileCanBeRenamed               = 253
	SCChangeStatLockState              = 254
	SCGetStaticArtBitmap               = 255
	// 256 - Reserved
	SCSetAlarm = 257
	// 258-261 HTTP
	SCInviteToParty      = 262
	SCRemoveFromParty    = 263
	SCPartyMessageTo     = 264
	SCPartySay           = 265
	SCPartyCanLootMe     = 266
	SCPartyAcceptInvite  = 267
	SCPartyDeclineInvite = 268
	SCPartyLeave         = 269
	SCPartyMembersList   = 270
	SCInParty            = 271
	//272-277 ICQ Obsolete
	SCGetTileFlags                = 278
	SCGetLandTileData             = 280
	SCGetStaticTileData           = 281
	SCGetLayerCount               = 282
	SCReadStaticsXY               = 283
	SCGetSurfaceZ                 = 284
	SCIsWorldCellPassable         = 285
	SCGetStaticTilesArray         = 286
	SCGetLandTilesArray           = 287
	SCClientPrint                 = 289
	SCClientPrintEx               = 290
	SCCloseClientUIWindow         = 291
	SCClientRequestObjectTarget   = 292
	SCClientRequestTileTarget     = 293
	SCClientTargetResponsePresent = 294
	SCClientTargetResponse        = 295
	SCCheckLagBegin               = 297
	SCCheckLagEnd                 = 298
	SCIsCheckLagEnd               = 299
	SCGetQuestArrow               = 300
	SCSetSilentMode               = 301
	SCGetSilentMode               = 302
	SCFillNewWindow               = 303 //304 busy SCAddToJournal
	SCGetStealthPath              = 305
	SCGetStealthProfilePath       = 306
	SCGetShardPath                = 307
	SCSendTextToUO                = 308
	SCSendTextToUOColor           = 309
	SCSetGlobal                   = 310
	SCGetGlobal                   = 311
	SCConsoleEntryReply           = 312
	SCConsoleEntryUnicodeReply    = 313
	SCSetRunUnmountTimer          = 316
	SCSetWalkMountTimer           = 317
	SCSetRunMountTimer            = 318
	SCSetWalkUnmountTimer         = 319
	SCGetRunMountTimer            = 320
	SCGetWalkMountTimer           = 321
	SCGetRunUnmountTimer          = 322
	SCGetWalkUnmountTimer         = 323
	SCStep                        = 324
	SCStepQ                       = 325
	SCMoveXYZ                     = 326
	SCMoveXY                      = 327
	SCSetBadLocation              = 328
	SCSetGoodLocation             = 329
	SCClearBadLocationList        = 330
	SCSetBadObject                = 331
	SCClearBadObjectList          = 332
	SCCheckLOS                    = 333
	SCGetPathArray                = 334
	SCGetPathArray3D              = 335
	SCSetFindInNulPoint           = 336
	SCGetFindInNulPoint           = 337
	SCGetMenuItems                = 338
	SCGetLastMenuItems            = 339
	SCFindTypesArrayEx            = 340
	SCGameServerIPString          = 341
	SCCloseClientGump             = 342
	SCGetProfileShardName         = 343
	SCGetLastStepQUsedDoor        = 344
	SCGetContextMenuRec           = 345
	SCClearSystemJournal          = 346
	SCGetMultis                   = 347
	SCClearInfoWindow             = 348
	SCGetBuffBarInfo              = 349
	SCConvertIntegerToFlags       = 350
	SCSkillCurrentValue           = 351
	SCChangeProfileEx             = 352
	SCMoverStop                   = 353
	SCSetARExtParams              = 354
	SCGetClientVersionInt         = 355
	SCUnequipItemsSetMacro        = 356
	SCEquipItemsSetMacro          = 357
	SCGetMenuItemsEx              = 358
	SCUseItemOnMobile             = 359
	SCGlobalChatJoinChannel       = 361
	SCGlobalChatLeaveChannel      = 362
	SCGlobalChatSendMsg           = 363
	SCGlobalChatActiveChannel     = 364
	SCGlobalChatChannelsList      = 365
	SCGetNextStepZ                = 366
	SCClientHide                  = 368
	SCGetSkillLockState           = 369
	SCEquipLastWeapon             = 370
	SCCreateChar                  = 371
	SCGetStatLockState            = 372
	SCBookGetPageText             = 373
	SCBookSetText                 = 374
	SCBookSetPageText             = 375
	SCBookClearText               = 376
	SCBookSetHeader               = 377
	//new mothods add HERE ^^^
	//mover vars!
	SCSetMoveOpenDoor          = 400
	SCGetMoveOpenDoor          = 401
	SCSetMoveThroughNPC        = 402
	SCGetMoveThroughNPC        = 403
	SCSetMoveThroughCorner     = 404
	SCGetMoveThroughCorner     = 405
	SCSetMoveHeuristicMult     = 406
	SCGetMoveHeuristicMult     = 407
	SCSetMoveCheckStamina      = 408
	SCGetMoveCheckStamina      = 409
	SCSetMoveTurnCost          = 410
	SCGetMoveTurnCost          = 411
	SCSetMoveBetweenTwoCorners = 412
	SCGetMoveBetweenTwoCorners = 413

	//mover vars finish nothing to add here!

	SCGetScriptsCount      = 450
	SCGetScriptPath        = 451
	SCGetScriptName        = 452
	SCGetScriptState       = 453
	SCStartScript          = 454
	SCStopScript           = 455
	SCPauseResumeSelScript = 456
	SCStopAllScripts       = 457

	SCMessenger_GetConnected = 501
	SCMessenger_SetConnected = 502
	SCSMessenger_GetToken    = 503
	SCSMessenger_SetToken    = 504
	SCMessenger_GetName      = 505
	SCMessenger_SendMessage  = 506

	SCAddFigure    = 550
	SCRemoveFigure = 551
	SCUpdateFigure = 552
	SCClearFigures = 553
)

var _SPELLS = map[string]uint16{
	// # 1st circle
	"clumsy":         1,
	"create food":    2,
	"feeblemind":     3,
	"heal":           4,
	"magic arrow":    5,
	"night sight":    6,
	"reactive armor": 7,
	"weaken":         8,
	//# 2nd circle
	"agility":      9,
	"cunning":      10,
	"cure":         11,
	"harm":         12,
	"magic trap":   13,
	"magic untrap": 14,
	"protection":   15,
	"strength":     16,
	//# 3rd circle
	"bless":         17,
	"fireball":      18,
	"magic lock":    19,
	"poison":        20,
	"telekinesis":   21,
	"teleport":      22,
	"unlock":        23,
	"wall of stone": 24,
	//# 4th circle
	"arch cure":       25,
	"arch protection": 26,
	"curse":           27,
	"fire field":      28,
	"greater heal":    29,
	"lightning":       30,
	"mana drain":      31,
	"recall":          32,
	//# 5th circle
	"blade spirit":     33,
	"dispel field":     34,
	"incognito":        35,
	"magic reflection": 36,
	"spell reflection": 36,
	"mind blast":       37,
	"paralyze":         38,
	"poison field":     39,
	"summon creature":  40,
	//# 6th circle
	"dispel":         41,
	"energy bolt":    42,
	"explosion":      43,
	"invisibility":   44,
	"mark":           45,
	"mass curse":     46,
	"paralyze field": 47,
	"reveal":         48,
	//# 7th circle
	"chain lightning": 49,
	"energy field":    50,
	"flame strike":    51,
	"gate travel":     52,
	"mana vampire":    53,
	"mass dispel":     54,
	"meteor swarm":    55,
	"polymorph":       56,
	//# 8th circle
	"earthquake":             57,
	"energy vortex":          58,
	"resurrection":           59,
	"summon air elemental":   60,
	"summon daemon":          61,
	"summon earth elemental": 62,
	"summon fire elemental":  63,
	"summon water elemental": 64,
	//# Necromancy
	"animate dead":     101,
	"blood oath":       102,
	"corpse skin":      103,
	"curse weapon":     104,
	"evil omen":        105,
	"horrific beast":   106,
	"lich form":        107,
	"mind rot":         108,
	"pain spike":       109,
	"poison strike":    110,
	"strangle":         111,
	"summon familiar":  112,
	"vampiric embrace": 113,
	"vengeful spirit":  114,
	"wither":           115,
	"wraith form":      116,
	"exorcism":         117,
	//# Paladin spells
	"cleanse by fire":   201,
	"close wounds":      202,
	"consecrate weapon": 203,
	"dispel evil":       204,
	"divine fury":       205,
	"enemy of one":      206,
	"holy light":        207,
	"noble sacrifice":   208,
	"remove curse":      209,
	"sacred journey":    210,
	//# Bushido spells
	"honorable execution": 401,
	"confidence":          402,
	"evasion":             403,
	"counter attack":      404,
	"lightning strike":    405,
	"momentum strike":     406,
	//# Ninjitsu spells
	"focus attack":    501,
	"death strike":    502,
	"animal form":     503,
	"ki attack":       504,
	"surprise attack": 505,
	"backstab":        506,
	"shadow jump":     507,
	"mirror image":    508,
	//# Spellweaving spells
	"arcane circle":      601,
	"gift of renewal":    602,
	"immolating weapon":  603,
	"attunement":         604,
	"thunderstorm":       605,
	"nature fury":        606,
	"summon fey":         607,
	"summon fiend":       608,
	"reaper form":        609,
	"wildfire":           610,
	"essence of wind":    611,
	"dryad allure":       612,
	"ethereal voyage":    613,
	"word of death":      614,
	"gift of life":       615,
	"arcane empowerment": 616,
	//# Mysticism spells
	"nether bolt":     678,
	"healing stone":   679,
	"pure magic":      680,
	"enchant":         681,
	"sleep":           682,
	"eagle strike":    683,
	"animated weapon": 684,
	"stone form":      685,
	"spell trigger":   686,
	"mass sleep":      687,
	"cleansing winds": 688,
	"bombard":         689,
	"spell plague":    690,
	"hail storm":      691,
	"nether cyclone":  692,
	"rising colossus": 693,
	//# Shared Passives
	"enchanted summoning": 715,
	"enchanted_summoning": 715,
	"intuition":           718,
	"warriors gifts":      733,
	"warriors_gifts":      733,
	"warrior's gifts":     733,
	//# Provocation
	"inspire":    701,
	"invigorate": 702,
	//# Peacemaking
	"resilience":   703,
	"perseverance": 704,
	//# Discordance
	"tribulation": 705,
	"despair":     706,
	//# Magery
	"death_ray":      707,
	"death ray":      707,
	"ethereal_burst": 708,
	"ethereal burst": 708,
	"ethereal_blast": 708,
	"ethereal blast": 708,
	//# Mysticism
	"nether_blast":  709,
	"nether blast":  709,
	"mystic_weapon": 710,
	"mystic weapon": 710,
	//# Necromancy
	"command_undead": 711,
	"command undead": 711,
	"conduit":        712,
	//# Spellweaving
	"mana_shield":   713,
	"mana shield":   713,
	"summon_reaper": 714,
	"summon reaper": 714,
	//# Bushido
	"anticipate_hit": 716,
	"anticipate hit": 716,
	"warcry":         717,
	//# Chivalry
	"rejuvenate": 719,
	"holy_fist":  720,
	"holy fist":  720,
	//# Ninjitsu
	"shadow":           721,
	"white_tiger_form": 722,
	"white tiger form": 722,
	//# Archery
	"flaming_shot":     723,
	"flaming shot":     723,
	"playing_the_odds": 724,
	"playing the odds": 724,
	//# Fencing
	"thrust": 725,
	"pierce": 726,
	//# Mace Fighting
	"stagger":   727,
	"toughness": 728,
	//# Swordsmanship
	"onslaught":   729,
	"focused_eye": 730,
	"focused eye": 730,
	//# Throwing
	"elemental_fury": 731,
	"elemental fury": 731,
	"called_shot":    732,
	"called shot":    732,
	//# Parrying
	"shield_bash":     734,
	"shield bash":     734,
	"bodyguard":       735,
	"heighten_senses": 736,
	"heighten senses": 736,
	//# Poisoning
	"tolerance":       737,
	"injected_strike": 738,
	"injected strike": 738,
	"potency":         739,
	//# Wrestling
	"rampage":       740,
	"fists_of_fury": 741,
	"fists of fury": 741,
	"knockout":      742,
	//# Animal Taming
	"whispering":      743,
	"boarding":        745,
	"combat_training": 744,
	"combat training": 744,
}

var _VIRTUES = map[string]uint32{
	"compassion":   0x69,
	"honesty":      0x6A,
	"honor":        0x6B,
	"humility":     0x6C,
	"justice":      0x6D,
	"sacrifice":    0x6E,
	"spirituality": 0x6F,
	"valor":        0x70,
}
