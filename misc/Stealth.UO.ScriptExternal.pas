unit Stealth.UO.ScriptExternal;

interface

uses Stealth.Base.ScriptExternal, Stealth.Base.ScriptMethodsConstants,
     Stealth.Base.NetwIPC, Stealth.Base.BytePackets, Stealth.Base.Types,
     Classes,SysUtils, Types,
     TypInfo, IdContext;

type
  TUOScriptExternalThread = class(TBaseScriptExternalThread)
  protected
    procedure ProceedOtherCMD(ScriptData : TScriptData; MethodNum, PacketLen : Word;PacketReader : TPacketReader;
                              ScriptMethodAnswer : TStealthCommPacket; AContext: TIdContext); override;
  end;


implementation

uses
  ClassCharacter,
  ScriptTypes,
  BasicTypedefs,
  Stealth.UO.PacketsCreation;

procedure TUOScriptExternalThread.ProceedOtherCMD(ScriptData : TScriptData; MethodNum,PacketLen : Word;PacketReader : TPacketReader;
                                                  ScriptMethodAnswer : TStealthCommPacket; AContext: TIdContext);
var
  TempContextMenu : TContextMenu;
  ClientTargetResponse : TTargetInfo;
  _XDrop, _YDrop, _ZDrop : Integer;
  _X, _Y, _X2, _Y2, _TileModel,_ObjType : Word;
  _Z, _Z2 : ShortInt;
  _WorldNum : Byte;
  LandTileData : TLandTileData;
  StaticTileData : TStaticTileData;
  DressSet : TLayersObjectsList;
  SkillID : Integer;
  Page : Word;
  ExtInfo : TExtendedInfo;
  ClilocRec : TClilocRec;
  GumpInfo : TGumpInfo;
  BitmapStream: TMemoryStream;
  TileDataFlagSet : TTileDataFlagSet;
  MapCell : TMapCell;
  StaticCell : TStaticCell;
  StaticItem : TStaticItemRealXY;
  FoundTilesArray : TFoundTilesArray;
  FoundTilesArrayDyn : TFoundTilesArrayDyn;
  point : Tpoint;
  PathArray : TPathArray;
  PathArrayDyn : TPathArrayDyn;
  MultiItems : TMultiItems;
  BuffBarInfo : TBuffBarInfo;
  flag : TTileDataFlags;
  ScriptIndex : Word;
  ScriptState : TScriptState;
  TempCardinal,ID, i : Cardinal;
  TempBool,TempBool2: Boolean;
  TempWord,TempWord2: Word;
  TempByte, TempByte2: Byte;
  TempStr,TempStr2,TempStr3: String;
  TempInt,TempInt2 : Integer;
  TempStrArr : TStringDynArray;
  TempDouble : Double;
  TypesArray,ColorsArray : TArray<Word>;
  ContArray : TArray<Cardinal>;
  MenuResponses : TMenuResponses;
  Figure : TMapFigure;
  cCRace: TCharRace;
  cCStr, cCDex, cCInt : Byte;
  cCProfileName, cShardName, cCName,
  cCskill1Name, cCskill2Name, cCskill3Name, cCskill4Name : String;
  cCSk1Value, cCSk2Value, cCSk3Value, cCSk4Value : Integer;
  cCStartingCity : Byte;
  cCFreeSlot : Cardinal;
  cCIsMale : Boolean;
begin

case MethodNum of
  SCGetBackpackID:
  begin
    ScriptMethodAnswer.AddParam(Script_GetBackpackID);
  end;
  SCGetSelfSex:
  begin
    ScriptMethodAnswer.AddParam(Script_GetSelfSex);
  end;
  SCGetCharTitle:
  begin
    ScriptMethodAnswer.AddStringParam(Script_GetCharTitle);
  end;
  SCGetSelfArmor:
  begin
    ScriptMethodAnswer.AddParam(Script_GetSelfArmor);
  end;
  SCGetSelfWeight:
  begin
    ScriptMethodAnswer.AddParam(Script_GetSelfWeight);
  end;
  SCGetSelfMaxWeight:
  begin
    ScriptMethodAnswer.AddParam(Script_GetSelfExtInfo.MaxWeight);
  end;
  SCGetSelfRace:
  begin
    ScriptMethodAnswer.AddParam(Script_GetSelfExtInfo.Race);
  end;
  SCGetSelfLuck:
  begin
    ScriptMethodAnswer.AddParam(Script_GetSelfExtInfo.Luck);
  end;
  SCGetSelfPetsMax:
  begin
    ScriptMethodAnswer.AddParam(Script_GetSelfExtInfo.PetsMax);
  end;
  SCGetSelfPetsCurrent:
  begin
    ScriptMethodAnswer.AddParam(Script_GetSelfExtInfo.PetsCurrent);
  end;
  SCGetSelfFireResist:
  begin
    ScriptMethodAnswer.AddParam(Script_GetSelfExtInfo.FireResist);
  end;
  SCGetSelfColdResist:
  begin
    ScriptMethodAnswer.AddParam(Script_GetSelfExtInfo.ColdResist);
  end;
  SCGetSelfPoisonResist:
  begin
    ScriptMethodAnswer.AddParam(Script_GetSelfExtInfo.PoisonResist);
  end;
  SCGetSelfEnergyResist:
  begin
    ScriptMethodAnswer.AddParam(Script_GetSelfExtInfo.EnergyResist);
  end;
  SCGetExtInfo:
  begin
    ScriptMethodAnswer.AddCustomType<TExtendedInfo>(Script_GetSelfExtInfo);
  end;
  SCGetName:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddStringParam(Script_GetName(ID));
  end;
  SCGetParent:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_GetParent(ID));
   end;
  SCObjAtLayerEx:
  begin
    TempByte := PacketReader.ReadByte;
    TempCardinal := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_ObjAtLayerEx(TempByte,TempCardinal));
  end;
  SCGetLayer:
  begin
    TempCardinal := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_GetLayer(TempCardinal));
  end;
  SCGetLastContainer:
  begin
    ScriptMethodAnswer.AddParam(Script_GetLastContainer);
  end;
  SCGetLastTarget:
  begin
    ScriptMethodAnswer.AddParam(Script_GetLastTarget);
  end;
  SCGetLastAttack:
  begin
    ScriptMethodAnswer.AddParam(Script_GetLastAttack);
  end;
  SCGetLastStatus:
  begin
    ScriptMethodAnswer.AddParam(Script_GetLastStatus);
  end;
  SCGetLastObject:
  begin
    ScriptMethodAnswer.AddParam(Script_GetLastObject);
  end;
  SCGetSelfStr:
  begin
    ScriptMethodAnswer.AddParam(Script_GetSelfStr);
  end;
  SCGetSelfInt:
  begin
    ScriptMethodAnswer.AddParam(Script_GetSelfInt);
  end;
  SCGetSelfDex:
  begin
    ScriptMethodAnswer.AddParam(Script_GetSelfDex);
  end;
  SCGetSelfLife:
  begin
    ScriptMethodAnswer.AddParam(Script_GetSelfLife);
  end;
  SCGetSelfMana:
  begin
    ScriptMethodAnswer.AddParam(Script_GetSelfMana);
  end;
  SCGetSelfStam:
  begin
    ScriptMethodAnswer.AddParam(Script_GetSelfStam);
  end;
  SCGetSelfMaxLife:
  begin
    ScriptMethodAnswer.AddParam(Script_GetSelfMaxLife);
  end;
  SCGetSelfMaxMana:
  begin
    ScriptMethodAnswer.AddParam(Script_GetSelfMaxMana);
  end;
  SCGetSelfMaxStam:
  begin
    ScriptMethodAnswer.AddParam(Script_GetSelfMaxStam);
  end;
  SCSetFindDistance:
  begin
    TempCardinal := PacketReader.ReadDWord;
    fFindDistance := TempCardinal;
  end;
  SCGetFindDistance:
  begin
    ScriptMethodAnswer.AddParam(fFindDistance);
  end;
  SCSetFindVertical:
  begin
    TempCardinal := PacketReader.ReadDWord;
    fFindVertical := TempCardinal;
  end;
  SCGetFindVertical:
  begin
    ScriptMethodAnswer.AddParam(fFindVertical);
  end;
  SCSetFindInNulPoint:
  begin
    TempBool := PacketReader.ReadBool;
    fFindInNulPoint := TempBool;
  end;
  SCGetFindInNulPoint:
  begin
    ScriptMethodAnswer.AddParam(fFindInNulPoint);
  end;
  SCFindTypeEx:
  begin
    _ObjType := PacketReader.ReadWord;
    TempWord := PacketReader.ReadWord;
    TempCardinal := PacketReader.ReadDWord;
    TempBool := PacketReader.ReadBool;

    TempCardinal := Script_FindTypeEx(_ObjType,TempWord,TempCardinal,TempBool, True);
    ScriptMethodAnswer.AddParam(TempCardinal);
    SetDLLContextfindFields(AContext);
  end;
  SCFindTypesArrayEx:
  begin
    TypesArray := PacketReader.ReadCustomType<TArray<Word>>;
    ColorsArray := PacketReader.ReadCustomType<TArray<Word>>;
    ContArray := PacketReader.ReadCustomType<TArray<Cardinal>>;
    TempBool := PacketReader.ReadBool;

    TempCardinal := Script_FindTypesArrayEx(TypesArray,ColorsArray,ContArray,TempBool);
    ScriptMethodAnswer.AddParam(TempCardinal);
    SetDLLContextfindFields(AContext);
  end;
  SCGetMultis:
  begin
    Script_GetMultis(MultiItems);
    ScriptMethodAnswer.AddCustomType<TMultiItems>(MultiItems);
  end;
  SCFindNotoriety:
  begin
    _ObjType := PacketReader.ReadWord;
    TempByte := PacketReader.ReadByte;

    TempCardinal := Script_FindNotoriety(_ObjType,TempByte);
    ScriptMethodAnswer.AddParam(TempCardinal);
    SetDLLContextfindFields(AContext);
  end;
  SCFindAtCoord:
  begin
    _X := PacketReader.ReadWord;
    _Y := PacketReader.ReadWord;

    TempCardinal := Script_FindAtCoord(_X,_Y);
    ScriptMethodAnswer.AddParam(TempCardinal);
    SetDLLContextfindFields(AContext);
  end;
  SCIgnore:
  begin
    TempCardinal := PacketReader.ReadDWord;
    Script_Ignore(TempCardinal);
    SetDLLContextfindFields(AContext);
  end;
  SCIgnoreOff:
  begin
    TempCardinal := PacketReader.ReadDWord;
    Script_IgnoreOff(TempCardinal);
    SetDLLContextfindFields(AContext);
  end;
  SCIgnoreReset:
  begin
    Script_IgnoreReset;
    SetDLLContextfindFields(AContext);
  end;
  SCGetIgnoreList:
  begin
    ScriptMethodAnswer.AddCustomType<TArray<Cardinal>>(ScriptData.fIgnoreList);
  end;
  SCGetFindedList:
  begin
    ScriptMethodAnswer.AddCustomType<TArray<Cardinal>>(ScriptData.fFindedList);
  end;
  SCGetFindItem:
  begin
    ScriptMethodAnswer.AddParam(ScriptData.fFindItem);
  end;
  SCGetFindCount:
  begin
    ScriptMethodAnswer.AddParam(ScriptData.fFindCount);
  end;
  SCGetFindQuantity:
  begin
    ScriptMethodAnswer.AddParam(ScriptData.fFindQuantity);
  end;
  SCGetFindFullQuantity:
  begin
    ScriptMethodAnswer.AddParam(ScriptData.fFindFullQuantity);
  end;
  SCPredictedX:
  begin
    ScriptMethodAnswer.AddParam(Script_PredictedX);
  end;
  SCPredictedY:
  begin
    ScriptMethodAnswer.AddParam(Script_PredictedY);
  end;
  SCPredictedZ:
  begin
    ScriptMethodAnswer.AddParam(Script_PredictedZ);
  end;
  SCPredictedDirection:
  begin
    ScriptMethodAnswer.AddParam(Script_PredictedDirection);
  end;
  SCGetX:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_GetX(ID));
  end;
  SCGetY:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_GetY(ID));
  end;
  SCGetZ:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_GetZ(ID));
  end;
  SCGetAltName:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddStringParam(Script_GetAltName(ID));
  end;
  SCGetTitle:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddStringParam(Script_GetTitle(ID));
  end;
  SCGetCliloc:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddStringParam(Script_GetCliloc(ID));
  end;
  SCGetType:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_GetType(ID));
  end;
  SCGetToolTipRec:
  begin
    ID := PacketReader.ReadDWord;
    ClilocRec := Script_GetToolTipRec(ID);
    ScriptMethodAnswer.AddCustomType<TClilocItemRecArray>(ClilocRec.Items);
  end;
  SCGetClilocByID:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddStringParam(Script_GetClilocByID(ID));
  end;
  SCGetQuantity:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_GetQuantity(ID));
  end;
  SCGetPrice:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_GetPrice(ID));
  end;
  SCGetDirection:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_GetDirection(ID));
  end;
  SCGetDistance:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_GetDistance(ID));
  end;
  SCGetColor:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_GetColor(ID));
  end;
  SCGetStr:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_GetStr(ID));
  end;
  SCGetInt:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_GetInt(ID));
  end;
  SCGetDex:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_GetDex(ID));
  end;
  SCGetHP:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_GetHP(ID));
  end;
  SCGetMaxHP:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_GetMaxHP(ID));
  end;
  SCGetMana:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_GetMana(ID));
  end;
  SCGetMaxMana:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_GetMaxMana(ID));
  end;
  SCGetStam:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_GetStam(ID));
  end;
  SCGetMaxStam:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_GetMaxStam(ID));
  end;
  SCGetNotoriety:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_GetNotoriety(ID));
  end;
  SCIsWarMode:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_IsWarMode(ID));
  end;
  SCIsNPC:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_IsNPC(ID));
  end;
  SCIsDead:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_IsDead(ID));
  end;
  SCIsRunning:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_IsRunning(ID));
  end;
  SCIsContainer:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_IsContainer(ID));
  end;
  SCIsHidden:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_IsHidden(ID));
  end;
  SCIsMovable:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_IsMovable(ID));
  end;
  SCIsYellowHits:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_IsYellowHits(ID));
  end;
  SCIsPoisoned:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_IsPoisoned(ID));
  end;
  SCIsParalyzed:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_IsParalyzed(ID));
  end;
  SCIsFemale:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_IsFemale(ID));
  end;
  SCOpenDoor:
  begin
    Script_OpenDoor;
  end;
  SCBow:
  begin
    Script_Bow;
  end;
  SCSalute:
  begin
    Script_Salute;
  end;
  SCGetPickupedItem:
  begin
    ScriptMethodAnswer.AddParam(TCharacter(CharObj).PickupedItem);
  end;
  SCSetPickupedItem:
  begin
    ID := PacketReader.ReadDWord;
    TCharacter(CharObj).PickupedItem := ID;
    TCharacter(CharObj).PickupedItemType := Script_GetType(ID);
  end;
  SCGetDropCheckCoord:
  begin
    ScriptMethodAnswer.AddParam(fDropCheckCoord);
  end;
  SCSetDropCheckCoord:
  begin
    TempBool := PacketReader.ReadBool;
    fDropCheckCoord := TempBool;
  end;
  SCDragItem:
  begin
    ID := PacketReader.ReadDWord;
    TempInt := PacketReader.ReadInt; //rescount
    TempBool := Script_DragItem(ID,TempInt);
    ScriptMethodAnswer.AddParam(TempBool);
  end;
  SCDropItem:
  begin
    ID := PacketReader.ReadDWord;
    _XDrop := PacketReader.ReadInt;
    _YDrop := PacketReader.ReadInt;
    _ZDrop := PacketReader.ReadInt;
    TempBool := Script_DropItem(ID,_XDrop,_YDrop,_ZDrop);
    ScriptMethodAnswer.AddParam(TempBool);
  end;
  SCGetDropDelay:
  begin
    ScriptMethodAnswer.AddParam(fDropDelay);
  end;
  SCSetDropDelay:
  begin
    TempCardinal := PacketReader.ReadDWord;
    fDropDelay := TempCardinal;
  end;
  SCRequestContextMenu:
  begin
    ID := PacketReader.ReadDWord;
    Script_RequestContextMenu(ID);
  end;
  SCSetContextMenuHook:
  begin
    ID := PacketReader.ReadDWord;
    TempByte := PacketReader.ReadByte;
    Script_SetContextMenuHook(ID,TempByte);
  end;
  SCGetContextMenu:
  begin
    TempStrArr := TCharacter(CharObj).GetContextMenuStrings;
    ScriptMethodAnswer.AddCustomType<TStringDynArray>(TempStrArr);
  end;
  SCClearContextMenu: Script_ClearContextMenu;
  SCCheckTradeState:
  begin
    ScriptMethodAnswer.AddParam(Script_CheckTradeState);
  end;
  SCGetTradeContainer:
  begin
    TempByte := PacketReader.ReadByte;
    TempByte2 := PacketReader.ReadByte;
    if (TempByte2 = 0) or (TempByte2 > 2) then Exit;
    ScriptMethodAnswer.AddParam(Script_GetTradeContainer(TempByte, TempByte2));
  end;
  SCGetTradeOpponent:
  begin
    TempByte := PacketReader.ReadByte;
    ScriptMethodAnswer.AddParam(Script_GetTradeOpponent(TempByte));
  end;
  SCGetTradeCount:
  begin
    ScriptMethodAnswer.AddParam(Script_GetTradeCount);
  end;
  SCGetTradeOpponentName:
  begin
    TempByte := PacketReader.ReadByte;
    ScriptMethodAnswer.AddStringParam(Script_GetTradeOpponentName(TempByte));
  end;
  SCTradeCheck:
  begin
    TempByte := PacketReader.ReadByte;
    TempByte2 := PacketReader.ReadByte;
    if (TempByte2 = 0) or (TempByte2 > 2) then Exit;
    ScriptMethodAnswer.AddParam(Script_TradeCheck(TempByte,TempByte2));
  end;
  SCConfirmTrade:
  begin
    TempByte := PacketReader.ReadByte;
    Script_ConfirmTrade(TempByte);
  end;
  SCCancelTrade:
  begin
    TempByte := PacketReader.ReadByte;
    ScriptMethodAnswer.AddParam(Script_CancelTrade(TempByte));
  end;
  SCWaitMenu:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    TempStr2 := PacketReader.ReadUString(PacketReader.ReadDWord);
    Script_WaitMenu(TempStr, TempStr2);
  end;
  SCAutoMenu:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    TempStr2 := PacketReader.ReadUString(PacketReader.ReadDWord);
    Script_AutoMenu(TempStr, TempStr2);
  end;
  SCMenuHookPresent:
  begin
    ScriptMethodAnswer.AddParam(Script_MenuHookPresent);
  end;
  SCMenuPresent:
  begin
    ScriptMethodAnswer.AddParam(Script_MenuPresent);
  end;
  SCCancelMenu: Script_CancelMenu;
  SCCloseMenu : Script_CloseMenu;
  SCWaitGumpInt:
  begin
    TempInt := PacketReader.ReadInt;
    Script_WaitGumpInt(TempInt);
  end;
  SCWaitGumpTextEntry:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    Script_WaitGumpTextEntry(TempStr);
  end;
  SCGumpAutoTextEntry:
  begin
    TempInt := PacketReader.ReadInt;
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    Script_GumpAutoTextEntry(TempInt,TempStr);
  end;
  SCGumpAutoRadiobutton:
  begin
    TempInt := PacketReader.ReadInt;
    TempInt2 := PacketReader.ReadInt;
    Script_GumpAutoRadiobutton(TempInt,TempInt2);
  end;
  SCGumpAutoCheckBox:
  begin
    TempInt := PacketReader.ReadInt;
    TempInt2 := PacketReader.ReadInt;
    Script_GumpAutoCheckBox(TempInt,TempInt2);
  end;
  SCNumGumpButton:
  begin
    TempWord := PacketReader.ReadWord;
    TempInt := PacketReader.ReadInt;
    ScriptMethodAnswer.AddParam(Script_NumGumpButton(TempWord,TempInt));
  end;
  SCNumGumpTextEntry:
  begin
    TempWord := PacketReader.ReadWord;
    TempInt := PacketReader.ReadInt;
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    ScriptMethodAnswer.AddParam(Script_NumGumpTextEntry(TempWord,TempInt,TempStr));
  end;
  SCNumGumpRadiobutton:
  begin
    TempWord := PacketReader.ReadWord;
    TempInt := PacketReader.ReadInt;
    TempInt2 := PacketReader.ReadInt;
    ScriptMethodAnswer.AddParam(Script_NumGumpRadiobutton(TempWord,TempInt,TempInt2));
  end;
  SCNumGumpCheckBox:
  begin
    TempWord := PacketReader.ReadWord;
    TempInt := PacketReader.ReadInt;
    TempInt2 := PacketReader.ReadInt;
    ScriptMethodAnswer.AddParam(Script_NumGumpCheckBox(TempWord,TempInt,TempInt2));
  end;
  SCGetGumpsCount:
  begin
    TempWord := Word(Script_GetGumpsCount);
    ScriptMethodAnswer.AddParam(TempWord);
  end;
  SCCloseSimpleGump:
  begin
    TempWord := PacketReader.ReadWord;
    Script_CloseSimpleGump(TempWord);
  end;
  SCGetGumpSerial:
  begin
    TempWord := PacketReader.ReadWord;
    ScriptMethodAnswer.AddParam(Script_GetGumpSerial(TempWord));
  end;
  SCGetGumpID:
  begin
    TempWord := PacketReader.ReadWord;
    ScriptMethodAnswer.AddParam(Script_GetGumpID(TempWord));
  end;
  SCGetGumpNoClose:
  begin
    TempWord := PacketReader.ReadWord;
    ScriptMethodAnswer.AddParam(Script_GetGumpNoClose(TempWord));
  end;
  SCGetGumpTextLines:
  begin
    TempWord := PacketReader.ReadWord;
    TCharacter(CharObj).Gumps.LockGumps;
    TempStrArr := TCharacter(CharObj).Gumps.GetGumpTextLines(TempWord);
    TCharacter(CharObj).Gumps.UnlockGumps;
    ScriptMethodAnswer.AddCustomType<TStringDynArray>(TempStrArr);
  end;
  SCGetGumpFullLines:
  begin
    TempWord := PacketReader.ReadWord;
    TCharacter(CharObj).Gumps.LockGumps;
    TempStrArr := TCharacter(CharObj).Gumps.GetGumpFullLines(TempWord);
    TCharacter(CharObj).Gumps.UnlockGumps;
    ScriptMethodAnswer.AddCustomType<TStringDynArray>(TempStrArr);
  end;
  SCGetGumpShortLines:
  begin
    TempWord := PacketReader.ReadWord;
    TCharacter(CharObj).Gumps.LockGumps;
    TempStrArr := TCharacter(CharObj).Gumps.GetGumpShortLines(TempWord);
    TCharacter(CharObj).Gumps.UnlockGumps;
    ScriptMethodAnswer.AddCustomType<TStringDynArray>(TempStrArr);
  end;
  SCGetGumpButtonsDescription:
  begin
    TempWord := PacketReader.ReadWord;
    TCharacter(CharObj).Gumps.LockGumps;
    TempStrArr := TCharacter(CharObj).Gumps.GetGumpButtonsDescription(TempWord);
    TCharacter(CharObj).Gumps.UnlockGumps;
    ScriptMethodAnswer.AddCustomType<TStringDynArray>(TempStrArr);
  end;
  SCGetGumpInfo:
  begin
    TempWord := PacketReader.ReadWord;
    Script_GetGumpInfo(TempWord, GumpInfo);
    ScriptMethodAnswer.AddCustomType<TGumpInfo>(GumpInfo);
  end;
  SCAddGumpIgnoreByID:
  begin
    TempInt := PacketReader.ReadInt;
    Script_AddGumpIgnoreByID(TempInt);
  end;
  SCAddGumpIgnoreBySerial:
  begin
    TempInt := PacketReader.ReadInt;
    Script_AddGumpIgnoreBySerial(TempInt);
  end;
  SCClearGumpsIgnore: Script_ClearGumpsIgnore;
  SCWearItem:
  begin
    TempByte := PacketReader.ReadByte;
    TempCardinal := PacketReader.ReadDWord;
    Script_WearItem(TempByte,TempCardinal);
  end;
  SCGetDressSpeed:
  begin
    ScriptMethodAnswer.AddParam(fDressSpeed);
  end;
  SCSetDressSpeed:
  begin
    TempWord := PacketReader.ReadWord;
    fDressSpeed := TempWord;
  end;
  SCSetDress:
  begin
    Script_SetDress;
  end;
  SCGetDressSet:
  begin
    DressSet := TCharacter(CharObj).DressSet;
    ScriptMethodAnswer.AddCustomType<TLayersObjectsList>(DressSet);
  end;
  SCAutoBuy:
  begin
    _ObjType := PacketReader.ReadWord;
    TempWord := PacketReader.ReadWord;
    TempWord2 := PacketReader.ReadWord;
    Script_AutoBuy(_ObjType, TempWord, TempWord2);
  end;
  SCGetShopList:
  begin
    TempStrArr := TCharacter(CharObj).GetCharShopList;
    ScriptMethodAnswer.AddCustomType<TStringDynArray>(TempStrArr);
  end;
  SCClearShopList: Script_ClearShopList;
  SCAutoBuyEx:
  begin
    _ObjType := PacketReader.ReadWord;
    TempWord := PacketReader.ReadWord;
    TempWord2 := PacketReader.ReadWord;
    TempCardinal := PacketReader.ReadDWord;
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    Script_AutoBuyEx(_ObjType, TempWord, TempWord2,TempCardinal,TempStr);
  end;
  SCGetAutoBuyDelay:
  begin
    ScriptMethodAnswer.AddParam(Script_GetAutoBuyDelay);
  end;
  SCSetAutoBuyDelay:
  begin
    TempWord := PacketReader.ReadWord;
    Script_SetAutoBuyDelay(TempWord);
  end;
  SCGetAutoSellDelay:
  begin
    ScriptMethodAnswer.AddParam(Script_GetAutoSellDelay);
  end;
  SCSetAutoSellDelay:
  begin
    TempWord := PacketReader.ReadWord;
    Script_SetAutoSellDelay(TempWord);
  end;
  SCAutoSell:
  begin
    _ObjType := PacketReader.ReadWord;
    TempWord := PacketReader.ReadWord;
    TempWord2 := PacketReader.ReadWord;
    Script_AutoSell(_ObjType, TempWord, TempWord2);
  end;
  SCRequestStats:
  begin
    ID := PacketReader.ReadDWord;
    Script_RequestStats(ID);
  end;
  SCHelpRequest: Script_HelpRequest;
  SCQuestRequest: Script_QuestRequest;
  SCRenameMobile:
  begin
    ID := PacketReader.ReadDWord;
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    Script_RenameMobile(ID, TempStr);
  end;
  SCMobileCanBeRenamed:
  begin
    ID := PacketReader.ReadDWord;
    ScriptMethodAnswer.AddParam(Script_MobileCanBeRenamed(ID));
  end;
  SCChangeStatLockState:
  begin
    TempByte := PacketReader.ReadByte;
    TempByte2 := PacketReader.ReadByte;
    Script_ChangeStatLockState(TempByte,TempByte2);
  end;
  SCGetStaticArtBitmap:
  begin
    ID := PacketReader.ReadDWord;
    TempWord := PacketReader.ReadWord;
    BitmapStream := TMemoryStream.Create;
    Script_GetStaticArtBitmapStream(Id, TempWord, BitmapStream);
    ScriptMethodAnswer.WriteBufer(BitmapStream.Memory^,BitmapStream.Size);
    BitmapStream.DisposeOf;
  end;
  SCCheckLagBegin:
  begin
    SinCheckLagBegin;
  end;
  SCCheckLagEnd: TCharacter(CharObj).CheckLagBegin := False;
  SCIsCheckLagEnd:
  begin
    TempBool := TCharacter(CharObj).CheckLagEnd;
    if TempBool then
      TCharacter(CharObj).CheckLagBegin := False;
    ScriptMethodAnswer.AddParam(TempBool);
  end;
  SCInviteToParty:
  begin
    ID := PacketReader.ReadDWord;
    Script_InviteToParty(ID);
  end;
  SCRemoveFromParty:
  begin
    ID := PacketReader.ReadDWord;
    Script_RemoveFromParty(ID);
  end;
  SCPartyMessageTo:
  begin
    ID := PacketReader.ReadDWord;
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    Script_PartyMessageTo(ID, TempStr);
  end;
  SCPartySay:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    Script_PartySay(TempStr);
  end;
  SCPartyCanLootMe:
  begin
    TempBool := PacketReader.ReadBool;
    Script_PartyCanLootMe(TempBool);
  end;
  SCPartyAcceptInvite:
  begin
    Script_PartyAcceptInvite;
  end;
  SCPartyDeclineInvite:
  begin
    Script_PartyDeclineInvite;
  end;
  SCPartyLeave:
  begin
    Script_PartyLeave;
  end;
  SCInParty:
  begin
    ScriptMethodAnswer.AddParam(Script_InParty);
  end;
  SCPartyMembersList:
  begin
    ScriptMethodAnswer.AddCustomType<TArray<Cardinal>>(TCharacter(CharObj).PartyMembersList);
  end;
  SCGetWorldNum:
  begin
    ScriptMethodAnswer.AddParam(Script_GetWorldNum);
  end;
  SCAddJournalIgnore:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    Script_AddJournalIgnore(TempStr);
  end;
  SCClearJournalIgnore: Script_ClearJournalIgnore;
  SCUAddChatUserIgnore:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    Script_AddChatUserIgnore(TempStr);
  end;
  SCClearChatUserIgnore: Script_ClearChatUserIgnore;
  SCGetHiddenStatus:
  begin
    ScriptMethodAnswer.AddParam(Script_GetHiddenStatus);
  end;
  SCGetPoisonedStatus:
  begin
    ScriptMethodAnswer.AddParam(Script_GetPoisonedStatus);
  end;
  SCGetParalyzedStatus:
  begin
    ScriptMethodAnswer.AddParam(Script_GetParalyzedStatus);
  end;
  SCGetDeadStatus:
  begin
    ScriptMethodAnswer.AddParam(Script_GetDeadStatus);
  end;
  SCGetWarTarget:
  begin
    ScriptMethodAnswer.AddParam(Script_GetWarTarget);
  end;
  SCSetWarMode:
  begin
    TempBool := PacketReader.ReadBool;
    Script_SetWarMode(TempBool);
  end;
  SCAttack:
  begin
    TempCardinal := PacketReader.ReadDWord;
    Script_Attack(TempCardinal);
  end;
  SCUseSelfPaperdollScroll:
  begin
    Script_UseSelfPaperdollScroll;
  end;
  SCUseOtherPaperdollScroll:
  begin
    ID := PacketReader.ReadDWord;
    Script_UseOtherPaperdollScroll(ID);
  end;
  SCGetTargetID:
  begin
    ScriptMethodAnswer.AddParam(Script_GetTargetID);
  end;
  SCCancelTarget:
  begin
    Script_CancelTarget;
  end;
  SCTargetToObject:
  begin
    TempCardinal := PacketReader.ReadDWord;
    Script_TargetToObject(TempCardinal);
  end;
  SCTargetToXYZ:
  begin
    _X := PacketReader.ReadWord;
    _Y := PacketReader.ReadWord;
    _Z := PacketReader.ReadShortInt;
    Script_TargetToXYZ(_X, _Y, _Z);
  end;
  SCTargetToTile:
  begin
    _TileModel := PacketReader.ReadWord;
    _X := PacketReader.ReadWord;
    _Y := PacketReader.ReadWord;
    _Z := PacketReader.ReadShortInt;
    Script_TargetToTile(_TileModel, _X, _Y, _Z);
  end;
  SCWaitTargetObject:
  begin
    TempCardinal := PacketReader.ReadDWord;
    Script_WaitTargetObject(TempCardinal);
  end;
  SCWaitTargetTile:
  begin
    _TileModel := PacketReader.ReadWord;
    _X := PacketReader.ReadWord;
    _Y := PacketReader.ReadWord;
    _Z := PacketReader.ReadShortInt;
    Script_WaitTargetTile(_TileModel,_X,_Y,_Z);
  end;
  SCWaitTargetXYZ:
  begin
    _X := PacketReader.ReadWord;
    _Y := PacketReader.ReadWord;
    _Z := PacketReader.ReadShortInt;
    Script_WaitTargetXYZ(_X,_Y,_Z);
  end;
  SCWaitTargetSelf:
  begin
    Script_WaitTargetObject(Script_GetSelfID);
  end;
  SCWaitTargetType:
  begin
    _ObjType := PacketReader.ReadWord;
    Script_WaitTargetType(_ObjType);
  end;
  SCCancelWaitTarget:
  begin
    Script_CancelWaitTarget;
  end;
  SCWaitTargetGround:
  begin
    _ObjType := PacketReader.ReadWord;
    Script_WaitTargetGround(_ObjType);
  end;
  SCWaitTargetLast:
  begin
    Script_WaitTargetLast;
  end;
  SCUsePrimaryAbility:
  begin
    Script_UsePrimaryAbility;
  end;
  SCUseSecondaryAbility:
  begin
    Script_UseSecondaryAbility;
  end;
  SCGetActiveAbility:
  begin
    ScriptMethodAnswer.AddStringParam(Script_GetAbility);
  end;
  SCToggleFly:
  begin
    Script_ToggleFly;
  end;
  SCGetSkillID:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    GetSkillID(TempStr, SkillID);
    ScriptMethodAnswer.AddParam(SkillID);
  end;
  SCUseSkill:
  begin
    SkillID := PacketReader.ReadInt;
    if (SkillID >= 1) and (SkillID < 250) then
      ScriptSendPacket(TPacketCreation.UseSkill(SkillID - 1));
  end;
  SCChangeSkillLockState:
  begin
    SkillID := PacketReader.ReadInt;
    TempByte := PacketReader.ReadByte;
    if (TempByte <= 2)  and (SkillID >= 1) and (SkillID < 250)
      and TCharacter(CharObj).Connected then
    begin
      TCharacter(CharObj).SetSkillLock(SkillID, TempByte);
      ScriptSendPacket(TPacketCreation.ChangeSkillLockState(SkillID - 1,TempByte));
    end;
  end;
  SCGetSkillCap:
  begin
    SkillID := PacketReader.ReadInt;
    TempDouble := 0;
    if (SkillID >= 1) and (SkillID < 250) then
      TempDouble := TCharacter(CharObj).GetSkillCapValue(SkillID);
    ScriptMethodAnswer.AddCustomType<Double>(TempDouble);
  end;
  SCSkillValue:
  begin
    SkillID := PacketReader.ReadInt;
    TempDouble := 0;
    if (SkillID >= 1) and (SkillID < 250) then
      TempDouble := TCharacter(CharObj).GetSkillUnmodifiedValue(SkillID);
    ScriptMethodAnswer.AddCustomType<Double>(TempDouble);
  end;
  SCSkillCurrentValue:
  begin
    SkillID := PacketReader.ReadInt;
    TempDouble := 0;
    if (SkillID >= 1) and (SkillID < 250) then
      TempDouble := TCharacter(CharObj).GetSkillValue(SkillID);
    ScriptMethodAnswer.AddCustomType<Double>(TempDouble);
  end;
  SCReqVirtuesGump: Script_ReqVirtuesGump;
  SCUseVirtue:
  begin
    TempCardinal := PacketReader.ReadDWord;
    ScriptSendPacket(TPacketCreation.UseVirtue(Script_GetSelfID,TempCardinal));
  end;
  SCCastSpell:
  begin
    TempInt := PacketReader.ReadInt;
    if TCharacter(CharObj).ClientVersionInt < 5000000 then // Old Format
      ScriptSendPacket(TPacketCreation.OldCastSpell(TempInt))
    else //AOS-format
      ScriptSendPacket(TPacketCreation.NewCastSpell(TempInt));
  end;
  SCIsActiveSpellAbility:
  begin
    TempInt := PacketReader.ReadInt;
    TempBool := TCharacter(CharObj).ActiveSpellAbilityList.Contains(Word(TempInt));
    ScriptMethodAnswer.AddParam(TempBool);
  end;
  SCSetCatchBag:
  begin
    TempCardinal := PacketReader.ReadDWord;
    Script_SetCatchBag(TempCardinal);
  end;
  SCUnsetCatchBag: Script_UnSetCatchBag;
  SCUseObject:
  begin
    TempCardinal := PacketReader.ReadDWord;
    if Script_IsNPC(TempCardinal) and (TCharacter(CharObj).ClientVersionInt > 05000000) then
      TCharacter(CharObj).RequestStats(TempCardinal);
    ScriptSendPacket(TPacketCreation.UseObject(TempCardinal));
  end;
  SCUseType:
  begin
    TempWord := PacketReader.ReadWord;
    TempWord2 := PacketReader.ReadWord;
    TempCardinal := Script_UseType(TempWord,TempWord2);
    ScriptMethodAnswer.AddParam(TempCardinal);
  end;
  SCUseFromGround:
  begin
    TempWord := PacketReader.ReadWord;
    TempWord2 := PacketReader.ReadWord;
    TempCardinal := Script_UseFromGround(TempWord,TempWord2);
    ScriptMethodAnswer.AddParam(TempCardinal);
  end;
  SCClickOnObject:
  begin
    TempCardinal := PacketReader.ReadDWord;
    Script_ClickOnObject(TempCardinal);
  end;
  SCGetTileFlags:
  begin
    TempByte := PacketReader.ReadByte;
    TempWord := PacketReader.ReadWord;
    TempCardinal := Script_GetTileFlags(TempByte + 1, TempWord);
    ScriptMethodAnswer.AddParam(TempCardinal);
  end;
  SCGetLandTileData:
  begin
    TempWord := PacketReader.ReadWord;
    LandTileData := Script_GetLandTileData(TempWord);
    ScriptMethodAnswer.AddCustomType<TLandTileData>(LandTileData);
  end;
  SCGetStaticTileData:
  begin
    TempWord := PacketReader.ReadWord;
    StaticTileData := Script_GetStaticTileData(TempWord);
    ScriptMethodAnswer.AddCustomType<TStaticTileData>(StaticTileData);
  end;
  SCGetCell:
  begin
    _X := PacketReader.ReadWord;
    _Y := PacketReader.ReadWord;
    _WorldNum := PacketReader.ReadByte;
    MapCell := Script_GetCell(_X,_Y,_WorldNum);
    ScriptMethodAnswer.AddCustomType<TMapCell>(MapCell);
  end;
  SCGetLayerCount:
  begin
    _X := PacketReader.ReadWord;
    _Y := PacketReader.ReadWord;
    _WorldNum := PacketReader.ReadByte;
    TempByte := Script_GetLayerCount(_X,_Y,_WorldNum);
    ScriptMethodAnswer.AddParam(TempByte);
  end;
  SCReadStaticsXY:
  begin
    _X := PacketReader.ReadWord;
    _Y := PacketReader.ReadWord;
    _WorldNum := PacketReader.ReadByte;
    StaticCell := Script_ReadStaticsXY(_X,_Y,_WorldNum);
    ScriptMethodAnswer.AddCustomType<TStaticCell>(StaticCell);
  end;
  SCGetSurfaceZ:
  begin
    _X := PacketReader.ReadWord;
    _Y := PacketReader.ReadWord;
    _WorldNum := PacketReader.ReadByte;
    TempByte := Byte(Script_GetSurfaceZ(_X,_Y,_WorldNum));
    ScriptMethodAnswer.AddParam(TempByte);
  end;
  SCIsWorldCellPassable:
  begin
    _X := PacketReader.ReadWord;
    _Y := PacketReader.ReadWord;
    _Z := PacketReader.ReadShortInt;
    _X2 := PacketReader.ReadWord;
    _Y2 := PacketReader.ReadWord;
    _WorldNum := PacketReader.ReadByte;
    TempBool := Script_IsWorldCellPassable(_X, _Y, _Z, _X2, _Y2, _Z2, _WorldNum);
    ScriptMethodAnswer.AddParam(TempBool).AddParam(_Z2);
  end;
  SCGetStaticTilesArray:
  begin
    _X := PacketReader.ReadWord;
    _Y := PacketReader.ReadWord;
    _X2 := PacketReader.ReadWord;
    _Y2 := PacketReader.ReadWord;
    _WorldNum := PacketReader.ReadByte;
    TypesArray := PacketReader.ReadCustomType<TArray<Word>>;

    TempWord := Script_GetStaticTilesArrayEx(_X, _Y, _X2, _Y2, _WorldNum,TypesArray,FoundTilesArray);
    SetLength(FoundTilesArrayDyn, TempWord);
    if TempWord > 0 then
      Move(FoundTilesArray[0],FoundTilesArrayDyn[0],TempWord*SizeOf(TFoundTile));
    ScriptMethodAnswer.AddCustomType<TFoundTilesArrayDyn>(FoundTilesArrayDyn);
  end;
  SCGetLandTilesArray:
  begin
    _X := PacketReader.ReadWord;
    _Y := PacketReader.ReadWord;
    _X2 := PacketReader.ReadWord;
    _Y2 := PacketReader.ReadWord;
    _WorldNum := PacketReader.ReadByte;
    TypesArray := PacketReader.ReadCustomType<TArray<Word>>;

    TempWord := Script_GetLandTilesArrayEx(_X, _Y, _X2, _Y2, _WorldNum,TypesArray,FoundTilesArray);
    SetLength(FoundTilesArrayDyn, TempWord);
    if TempWord > 0 then
      Move(FoundTilesArray[0],FoundTilesArrayDyn[0],TempWord*SizeOf(TFoundTile));
    ScriptMethodAnswer.AddCustomType<TFoundTilesArrayDyn>(FoundTilesArrayDyn);
  end;
  SCClientPrint:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    Script_ClientPrint(TempStr);
  end;
  SCClientPrintEx:
  begin
    TempCardinal := PacketReader.ReadDWord;
    TempWord := PacketReader.ReadWord;
    TempWord2 := PacketReader.ReadWord;
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    Script_ClientPrintEx(TempCardinal,TempWord,TempWord2,TempStr);
  end;
  SCCloseClientUIWindow:
  begin
    TempByte := PacketReader.ReadByte;
    TempCardinal := PacketReader.ReadDWord;
    Script_CloseClientUIWindow(TUIWindowType(TempByte), TempCardinal);
  end;
  SCCloseClientGump:
  begin
    TempCardinal := PacketReader.ReadDWord;
    Script_CloseClientGump(TempCardinal);
  end;
  SCClientRequestObjectTarget:
  begin
    Script_ClientRequestObjectTarget;
  end;
  SCClientRequestTileTarget:
  begin
    Script_ClientRequestTileTarget;
  end;
  SCClientTargetResponsePresent:
  begin
    ScriptMethodAnswer.AddParam(Script_ClientTargetResponsePresent);
  end;
  SCClientTargetResponse:
  begin
    ClientTargetResponse := Script_ClientTargetResponse;
    ScriptMethodAnswer.AddCustomType<TTargetInfo>(ClientTargetResponse);
  end;
  SCGetQuestArrow:
  begin
    TCharacter(CharObj).GetQuestArrow(point);
    ScriptMethodAnswer.AddCustomType<Tpoint>(point);
  end;
  SCGetSilentMode:
  begin
    ScriptMethodAnswer.AddParam(TCharacter(CharObj).SilentMode);
  end;
  SCSetSilentMode:
  begin
    TCharacter(CharObj).SilentMode := PacketReader.ReadBool;
  end;
  SCFillNewWindow:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    if TempStr.Length > 0 then
      Script_FillInfoWindow(TempStr);
  end;
  SCClearInfoWindow:
  begin
    Script_ClearInfoWindow;
  end;
  SCSendTextToUO:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    Script_SendTextToUO(TempStr);
  end;
  SCSendTextToUOColor:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    TempCardinal := PacketReader.ReadWord;
    Script_SendTextToUOColor(TempStr, TempCardinal);
  end;
  SCConsoleEntryReply:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    Script_ConsoleEntryReply(TempStr);
  end;
  SCConsoleEntryUnicodeReply:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    Script_ConsoleEntryUnicodeReply(TempStr);
  end;
  SCStep:
  begin
    TempByte := PacketReader.ReadByte;
    TempBool := PacketReader.ReadBool;
    ScriptMethodAnswer.AddParam(Script_Step(TempByte,TempBool));
  end;
  SCStepQ:
  begin
    TempByte := PacketReader.ReadByte;
    TempBool := PacketReader.ReadBool;
    ScriptMethodAnswer.AddParam(Script_Stepq(TempByte,TempBool));
  end;
  SCMoveXYZ:
  begin
    _X := PacketReader.ReadWord;
    _Y := PacketReader.ReadWord;
    _Z := PacketReader.ReadShortInt;
    TempInt := PacketReader.ReadInt;
    TempInt2 := PacketReader.ReadInt;
    TempBool := PacketReader.ReadBool;
    TempBool2 := Script_MoveXYZ(_X,_Y,_Z,TempInt,TempInt2,TempBool);
    ScriptMethodAnswer.AddParam(TempBool2);
  end;
  SCMoveXY:
  begin
    _X := PacketReader.ReadWord;
    _Y := PacketReader.ReadWord;
    TempBool := PacketReader.ReadBool;
    TempInt := PacketReader.ReadInt;
    TempBool2 := PacketReader.ReadBool;
    TempBool2 := Script_MoveXY(_X,_Y,TempBool,TempInt,TempBool2);
    ScriptMethodAnswer.AddParam(TempBool2);
  end;
  SCSetBadLocation:
  begin
    _X := PacketReader.ReadWord;
    _Y := PacketReader.ReadWord;
    Script_SetBadLocation(_X,_Y);
  end;
  SCSetGoodLocation:
  begin
    _X := PacketReader.ReadWord;
    _Y := PacketReader.ReadWord;
    Script_SetGoodLocation(_X,_Y);
  end;
  SCClearBadLocationList: Script_ClearBadLocationList;
  SCSetBadObject:
  begin
    _ObjType := PacketReader.ReadWord;
    TempWord := PacketReader.ReadWord;
    TempByte := PacketReader.ReadShortInt;
    Script_SetBadObject(_ObjType,TempWord,TempByte);
  end;
  SCClearBadObjectList: Script_ClearBadObjectList;
  SCCheckLOS:
  begin
    _X := PacketReader.ReadWord;
    _Y := PacketReader.ReadWord;
    _Z := PacketReader.ReadShortInt;
    _X2 := PacketReader.ReadWord;
    _Y2 := PacketReader.ReadWord;
    _Z2 := PacketReader.ReadShortInt;
    TempByte := PacketReader.ReadByte;
    TempByte2 := PacketReader.ReadByte;
    TempCardinal := PacketReader.ReadDWord;
    TCharacter(CharObj).World.LOSOptions := TempByte2 OR TempCardinal;

    TempBool2 := Script_CheckLOS(_X,_Y,_Z,_X2,_Y2,_Z2,TempByte);
    ScriptMethodAnswer.AddParam(TempBool2);
  end;
  SCGetPathArray:
  begin
    _X := PacketReader.ReadWord;
    _Y := PacketReader.ReadWord;
    TempBool := PacketReader.ReadBool;
    TempInt := PacketReader.ReadInt;
    TempInt := Script_GetPathArray(_X,_Y,TempBool,TempInt,PathArray );

    if TempInt < 0 then
      TempInt := 0;

    SetLength(PathArrayDyn, TempInt);
    if TempInt > 0 then
      Move(PathArray[0],PathArrayDyn[0],TempInt*SizeOf(TMyPoint));
    ScriptMethodAnswer.AddCustomType<TPathArrayDyn>(PathArrayDyn);
  end;
  SCGetPathArray3D:
  begin
    _X := PacketReader.ReadWord;
    _Y := PacketReader.ReadWord;
    _Z := PacketReader.ReadShortInt;
    _X2 := PacketReader.ReadWord;
    _Y2 := PacketReader.ReadWord;
    _Z2 := PacketReader.ReadShortInt;
    TempByte := PacketReader.ReadByte;
    TempInt := PacketReader.ReadInt;
    TempInt2 := PacketReader.ReadInt;
    TempBool := PacketReader.ReadBool;

    TempInt := Script_GetPathArray3D(_X,_Y,_Z,_X2,_Y2,_Z2, TempByte, TempInt,TempInt2,TempBool,PathArray );

    if TempInt < 0 then
      TempInt := 0;

    SetLength(PathArrayDyn, TempInt);
    if TempInt > 0 then
      Move(PathArray[0],PathArrayDyn[0],TempInt*SizeOf(TMyPoint));
    ScriptMethodAnswer.AddCustomType<TPathArrayDyn>(PathArrayDyn);
  end;
  SCGetRunUnmountTimer:
  begin
    ScriptMethodAnswer.AddParam(Script_GetRunUnmountTimer);
  end;
  SCSetRunUnmountTimer:
  begin
    TempWord := PacketReader.ReadWord;
    Script_SetRunUnmountTimer(TempWord);
  end;
  SCGetWalkMountTimer:
  begin
    ScriptMethodAnswer.AddParam(Script_GetWalkMountTimer);
  end;
  SCSetWalkMountTimer:
  begin
    TempWord := PacketReader.ReadWord;
    Script_SetWalkMountTimer(TempWord);
  end;
  SCGetRunMountTimer:
  begin
    ScriptMethodAnswer.AddParam(Script_GetRunMountTimer);
  end;
  SCSetRunMountTimer:
  begin
    TempWord := PacketReader.ReadWord;
    Script_SetRunMountTimer(TempWord);
  end;
  SCGetWalkUnmountTimer:
  begin
    ScriptMethodAnswer.AddParam(Script_GetWalkUnmountTimer);
  end;
  SCSetWalkUnmountTimer:
  begin
    TempWord := PacketReader.ReadWord;
    Script_SetWalkUnmountTimer(TempWord);
  end;
  SCGetMenuItems:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    TempStrArr := GetMenuItems2Arr(TempStr);
    ScriptMethodAnswer.AddCustomType<TStringDynArray>(TempStrArr);
  end;
  SCGetLastMenuItems:
  begin
    TempStrArr := GetLastMenuItems2Arr;
    ScriptMethodAnswer.AddCustomType<TStringDynArray>(TempStrArr);
  end;
  SCGameServerIPString:
  begin
    ScriptMethodAnswer.AddStringParam(Script_GameServerIPString);
  end;
  SCGetLastStepQUsedDoor:
  begin
    ScriptMethodAnswer.AddParam(fLastStepQUsedDoor);
  end;
  SCGetContextMenuRec:
  begin
    TempContextMenu := TCharacter(CharObj).ContextMenu;
    ScriptMethodAnswer.AddCustomType<TContextMenu>(TempContextMenu);
  end;
  SCGetBuffBarInfo:
  begin
    BuffBarInfo := Script_GetBuffBarInfo;
    ScriptMethodAnswer.AddCustomType<TBuffArr>(BuffBarInfo.Buffs);
  end;
  SCConvertIntegerToFlags:
  begin
    TempByte := PacketReader.ReadByte;
    TempCardinal := PacketReader.ReadDWord;
    TileDataFlagSet := Script_ConvertIntegerToFlags(TempByte+1, TempCardinal);

    for flag := Low(TTileDataFlags) to High(TTileDataFlags) do
      if flag in TileDataFlagSet then
        TempStrArr := TempStrArr + [GetEnumName(TypeInfo(TTileDataFlags), integer(flag))];

    ScriptMethodAnswer.AddCustomType<TStringDynArray>(TempStrArr);

  end;
  SCChangeProfileEx:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    TempStr2 := PacketReader.ReadUString(PacketReader.ReadDWord);
    TempStr3 := PacketReader.ReadUString(PacketReader.ReadDWord);

    TempInt := Script_ExtChangeProfile(TempStr,TempStr2,TempStr3);
    ScriptMethodAnswer.AddParam(TempInt);
  end;
  SCSetARExtParams:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    TempStr2 := PacketReader.ReadUString(PacketReader.ReadDWord);

    TempBool := PacketReader.ReadBool;
    Script_SetARExtParams(TempStr,TempStr2,TempBool);
  end;
  SCMoverStop:
  begin
    Script_MoverStop;
  end;
  SCSetMoveOpenDoor:
  begin
    TempBool := PacketReader.ReadBool;
    Script_SetMoveOpenDoor(TempBool);
  end;
  SCGetMoveOpenDoor:
  begin
    ScriptMethodAnswer.AddParam(Script_GetMoveOpenDoor);
  end;
  SCSetMoveThroughNPC:
  begin
    TempWord := PacketReader.ReadWord;
    Script_SetMoveThroughNPC(TempWord);
  end;
  SCGetMoveThroughNPC:
  begin
    ScriptMethodAnswer.AddParam(Script_GetMoveThroughNPC);
  end;
  SCSetMoveThroughCorner:
  begin
    TempBool := PacketReader.ReadBool;
    Script_SetMoveThroughCorner(TempBool);
  end;
  SCGetMoveThroughCorner:
  begin
    ScriptMethodAnswer.AddParam(Script_GetMoveThroughCorner);
  end;
  SCSetMoveHeuristicMult:
  begin
    TempInt := PacketReader.ReadInt;
    Script_SetMoveHeuristicMult(TempInt);
  end;
  SCGetMoveHeuristicMult:
  begin
    ScriptMethodAnswer.AddParam(Script_GetMoveHeuristicMult);
  end;
  SCSetMoveCheckStamina:
  begin
    TempWord := PacketReader.ReadWord;
    TCharacter(CharObj).World.CheckStamina := TempWord;
  end;
  SCGetMoveCheckStamina:
  begin
    ScriptMethodAnswer.AddParam(TCharacter(CharObj).World.CheckStamina);
  end;
  SCSetMoveTurnCost:
  begin
    TempInt := PacketReader.ReadInt;
    TCharacter(CharObj).World.TurnCost := TempInt;
  end;
  SCGetMoveTurnCost:
  begin
    ScriptMethodAnswer.AddParam(TCharacter(CharObj).World.TurnCost);
  end;
  SCSetMoveBetweenTwoCorners:
  begin
    TempBool := PacketReader.ReadBool;
    TCharacter(CharObj).World.WalkBetweenTwoCorners := TempBool;
  end;
  SCGetMoveBetweenTwoCorners:
  begin
    ScriptMethodAnswer.AddParam(TCharacter(CharObj).World.WalkBetweenTwoCorners);
  end;
  SCGetClientVersionInt:
  begin
    ScriptMethodAnswer.AddParam(TCharacter(CharObj).ClientVersionInt);
  end;
  SCUnequipItemsSetMacro:
  begin
    Script_undress;
  end;
  SCEquipItemsSetMacro:
  begin
    Script_EquipDressSet;
  end;
  SCGetMenuItemsEx:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);

    MenuResponses := Script_GetMenuItemsEx(TempStr);
    ScriptMethodAnswer.AddCustomType<TMenuResponses>(MenuResponses);
  end;
  SCUseItemOnMobile:
  begin
    TempCardinal := PacketReader.ReadDWord;
    ID := PacketReader.ReadDWord;
    Script_UseItemOnMobile(TempCardinal, ID);
  end;
  SCGlobalChatJoinChannel:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    Script_GlobalChatJoinChannel(TempStr);
  end;
  SCGlobalChatLeaveChannel:
  begin
    Script_GlobalChatLeaveChannel;
  end;
  SCGlobalChatSendMsg:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    Script_GlobalChatSendMsg(TempStr);
  end;
  SCGlobalChatActiveChannel:
  begin
    TempStr := Script_GlobalChatActiveChannel;
    ScriptMethodAnswer.AddStringParam(TempStr);
  end;
  SCGlobalChatChannelsList:
  begin
    TempStrArr := Script_GlobalChatChannelsList;
    ScriptMethodAnswer.AddCustomType<TStringDynArray>(TempStrArr);
  end;
  SCAddFigure:
  begin
    Figure := PacketReader.ReadCustomType<TMapFigure>;
    TCharacter(CharObj).MapFigures.AddFigure(Figure);
  end;
  SCRemoveFigure:
  begin
    ID := PacketReader.ReadDWord;
    TCharacter(CharObj).MapFigures.RemoveFigure(ID);
  end;
  SCUpdateFigure:
  begin
    ID := PacketReader.ReadDWord;

    Figure := PacketReader.ReadCustomType<TMapFigure>;

    TCharacter(CharObj).MapFigures.UpdateFigure(ID, Figure);
  end;
  SCClearFigures:
  begin
    TCharacter(CharObj).MapFigures.ClearFigures;
  end;
  SCGetNextStepZ:
  begin
    _X := PacketReader.ReadWord;
    _Y := PacketReader.ReadWord;
    _X2 := PacketReader.ReadWord;
    _Y2 := PacketReader.ReadWord;
    TempByte := PacketReader.ReadByte;
    _Z := PacketReader.ReadShortInt;

    _Z2 := Script_GetNextStepZ(_X,_Y,_X2,_Y2,TempByte,_Z );
    ScriptMethodAnswer.AddParam(_Z2);
  end;
  SCClientHide:
  begin
    ID := PacketReader.ReadDWord;
    Script_ClientHide(ID);
  end;
  SCGetSkillLockState:
  begin
    SkillID := PacketReader.ReadInt;
    TempByte := TCharacter(CharObj).GetSkillLock(SkillID);
    ScriptMethodAnswer.AddParam(TempByte);
  end;
  SCEquipLastWeapon:
  begin
    Script_EquipLastWeapon;
  end;
  SCGetStatLockState:
  begin
    TempByte := PacketReader.ReadByte;
    _Z := Script_GetStatLockState(TempByte);
    ScriptMethodAnswer.AddParam(_Z);
  end;

  SCBookGetPageText:
  begin
    Page := PacketReader.ReadWord;
    TempStr := Script_BookGetPageText(Page);
    ScriptMethodAnswer.AddStringParam(TempStr);
  end;
  SCBookSetText:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    Script_BookSetText(TempStr);
  end;
  SCBookSetPageText:
  begin
    Page := PacketReader.ReadWord;
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    Script_BookSetPageText(Page,TempStr);
  end;
  SCBookClearText:
  begin
    Script_BookClearText;
  end;
  SCBookSetHeader:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    TempStr2 := PacketReader.ReadUString(PacketReader.ReadDWord);
    Script_BookSetHeader(TempStr,TempStr2);
  end;
  SCDumpObjectsCache:
  begin
    Script_Dump;
  end;
  SCCreateChar:
  begin
{    procedure Script_CreateChar(ProfileName, ShardName, NewCharName : String; IsMale : Boolean; Race: TCharRace; Str, Dex, Int : Byte;
                                   skill1Name, skill2Name, skill3Name, skill4Name : String; Sk1Value, Sk2Value, Sk3Value, Sk4Value : Integer;
                                   StartingCity : Byte; FreeSlot : Cardinal);
}
    cCProfileName := PacketReader.ReadUString(PacketReader.ReadDWord);
    cShardName := PacketReader.ReadUString(PacketReader.ReadDWord);
    cCName := PacketReader.ReadUString(PacketReader.ReadDWord);
    cCIsMale := PacketReader.ReadBool;
    cCRace := TCharRace(PacketReader.ReadByte);
    cCStr := PacketReader.ReadByte;
    cCDex := PacketReader.ReadByte;
    cCInt := PacketReader.ReadByte;
    cCskill1Name := PacketReader.ReadUString(PacketReader.ReadDWord);
    cCskill2Name := PacketReader.ReadUString(PacketReader.ReadDWord);
    cCskill3Name := PacketReader.ReadUString(PacketReader.ReadDWord);
    cCskill4Name := PacketReader.ReadUString(PacketReader.ReadDWord);
    cCSk1Value := PacketReader.ReadInt;
    cCSk2Value := PacketReader.ReadInt;
    cCSk3Value := PacketReader.ReadInt;
    cCSk4Value := PacketReader.ReadInt;
    cCStartingCity := PacketReader.ReadByte;
    cCFreeSlot := PacketReader.ReadDWord;
    Script_CreateChar(cCProfileName, cShardName, cCName, cCIsMale, cCRace, cCStr, cCDex, cCInt,
                      cCskill1Name, cCskill2Name, cCskill3Name, cCskill4Name,
                      cCSk1Value, cCSk2Value, cCSk3Value, cCSk4Value,
                      cCStartingCity, cCFreeSlot);
  end;
  SCGetScriptsCount:
  begin
    TempWord := Script_GetScriptsCount;
    ScriptMethodAnswer.AddParam(TempWord);
  end;
  SCGetScriptPath:
  begin
    ScriptIndex := PacketReader.ReadWord;
    TempStr := Script_GetScriptPath(ScriptIndex);
    ScriptMethodAnswer.AddStringParam(TempStr);
  end;
  SCGetScriptName:
  begin
    ScriptIndex := PacketReader.ReadWord;
    TempStr := Script_GetScriptName(ScriptIndex);
    ScriptMethodAnswer.AddStringParam(TempStr);
  end;
  SCGetScriptState:
  begin
    ScriptIndex := PacketReader.ReadWord;
    ScriptState := Script_GetScriptState(ScriptIndex);
    ScriptMethodAnswer.AddParam(ScriptState);
  end;
  SCStartScript:
  begin
    TempStr := PacketReader.ReadUString(PacketReader.ReadDWord);
    TempWord := Script_StartScript(TempStr);
    ScriptMethodAnswer.AddParam(TempWord);
  end;
  SCStopScript:
  begin
    ScriptIndex := PacketReader.ReadWord;
    Script_StopScript(ScriptIndex);
  end;
  SCPauseResumeSelScript:
  begin
    ScriptIndex := PacketReader.ReadWord;
    Script_PauseResumeScript(ScriptIndex);
  end;
  SCStopAllScripts:
  begin
    Script_StopAllScripts;
  end

  else
  begin
    TCharacter(CharObj).AddToSystemJournal('Unknown packet coming from external script, id = ' + MethodNum.ToString);
    TStealthCommPacket.Create(SCErrorReport).AddReturnID(00).AddParam(8).Send(AContext);
    PacketReader.IncPacketPos(PacketLen);
  end;
end;

end;

end.
