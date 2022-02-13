
unit ScriptTypes;
interface

uses

  {$IFNDEF HAS_FMX}
   Graphics,
  {$ENDIF}
  System.UITypes;

type
{$TYPEINFO ON}
// enums
TTileDataFlags = (tsfBackground, tsfWeapon, tsfTransparent, tsfTranslucent, tsfWall, tsfDamaging, tsfImpassable, tsfWet, tsfUnknown, tsfSurface, tsfBridge, tsfGeneric, tsfWindow, tsfNoShoot, tsfPrefixA, tsfPrefixAn, tsfInternal, tsfFoliage, tsfPartialHue, tsfUnknown1, tsfMap, tsfContainer, tsfWearable, tsfLightSource, tsfAnimated, tsfNoDiagonal, tsfUnknown2, tsfArmor, tsfRoof, tsfDoor, tsfStairBack, tsfStairRight, tlfTranslucent, tlfWall, tlfDamaging, tlfImpassable, tlfWet, tlfSurface, tlfBridge, tlfPrefixA, tlfPrefixAn, tlfInternal, tlfMap, tlfUnknown3);
TFloatFormat = (ffGeneral, ffExponent, ffFixed, ffNumber, ffCurrency);
TFigureCoord = (fcWorld, fcScreen);
TFigureKind = (fkLine, fkEllipse, fkRectangle, fkDirection, fkText);
TMsgDlgType = (mtWarning, mtError, mtInformation, mtConfirmation, mtCustom);
TMsgDlgBtn = (mbYes, mbNo, mbOK, mbCancel, mbAbort, mbRetry, mbAll, mbNoToAll, mbYesToAll, mbHelp);
TUIWindowType = (wtPaperdoll, wtStatus, wtCharProfile, wtContainer);

TCharRace = (crHuman, crElf, crGargoyle);

type OPENARRAYOFU16 = array of Word;
     OPENARRAYOFU32 = array of Cardinal;

type
  TCardinalDynArray_ = array of Cardinal;
  TCardinalDynArray = type TCardinalDynArray_;
  PCardinalDynArray = ^TCardinalDynArray;
// records
TAboutStealth = packed record
  StealthVersion : String;
  Build : String;
  BuildDate : String;
  GITRevNumber : Word;
end;
PAboutStealth = ^TAboutStealth;

TTileName = array[0..19] of Byte;

TLandTileData = packed record
  Flags : Cardinal;
  Flags2 : Cardinal;
  TextureID : Word;
  Name : TTileName;
end;
PLandTileData = ^TLandTileData;

TStaticTileData = packed record
  Flags : Int64;
  Weight : Word;
  AnimID : Word;
  Height : Integer;
  RadarColorRGBA : Cardinal;
  Name : TTileName;
end;
PStaticTileData = ^TStaticTileData;

TStaticItemRealXY = packed record
  Tile : Word;
  X : Word;
  Y : Word;
  Z : ShortInt;
  Color : Word;
end;
PStaticItemRealXY = ^TStaticItemRealXY;

TStaticCell = packed record
  Statics : array of TStaticItemRealXY;
  StaticCount : Byte;
end;
PStaticCell = ^TStaticCell;

TMapCell = packed record
  Tile : Word;
  Z : ShortInt;
end;
PMapCell = ^TMapCell;

TMapCell2 = packed record
  Tile : Word;
  Z : Word;
//  Tile2 : Word;
end;

TMyPoint = packed record
  X : Word;
  Y : Word;
  Z : ShortInt;
end;
PMyPoint = ^TMyPoint;

TFoundTile = packed record
  Tile : Word;
  X : SmallInt;
  Y : SmallInt;
  Z : ShortInt;
end;
PFoundTile = ^TFoundTile;

{$IFDEF HAS_FMX}
TBrushStyle = (bsSolid, bsClear, bsHorizontal, bsVertical,
  bsFDiagonal, bsBDiagonal, bsCross, bsDiagCross);
{$ENDIF}

TMapFigure = packed record
  kind : TFigureKind;
  coord : TFigureCoord;
  x1 : Integer;
  y1 : Integer;
  x2 : Integer;
  y2 : Integer;
  brushColor : TColor;
  brushStyle : TBrushStyle;
  color : TColor;
  worldNum : Byte;
  text : String;
end;
PMapFigure = ^TMapFigure;

TTargetInfo = packed record
  ID : Cardinal;
  Tile : Word;
  X : Word;
  Y : Word;
  Z : ShortInt;
end;
PTargetInfo = ^TTargetInfo;

TExtendedInfo = packed record
  MaxWeight : Word;
  Race : Byte;
  StatCap : Word;
  PetsCurrent : Byte;
  PetsMax : Byte;
  FireResist : Word;
  ColdResist : Word;
  PoisonResist : Word;
  EnergyResist : Word;
  Luck : SmallInt;
  DamageMin : Word;
  DamageMax : Word;
  Tithing_points : Cardinal;


  ArmorMax : Word;
  fireresistMax : Word;
  coldresistMax : Word;
  poisonresistMax : Word;
  energyresistMax : Word;
  DefenseChance : Word;
  DefensceChanceMax : Word;
  Hit_Chance_Incr : Word;
  Damage_Incr : Word;
  Swing_Speed_Incr : Word;
  Lower_Reagent_Cost : Word;
  Spell_Damage_Incr : Word;
  Faster_Cast_Recovery : Word;
  Faster_Casting : Word;
  Lower_Mana_Cost : Word;
  HP_Regen : Word;
  Stam_Regen : Word;
  Mana_Regen : Word;
  Reflect_Phys_Damage : Word;
  Enhance_Potions : Word;
  Strength_Incr : Word;
  Dext_Incr : Word;
  Int_Incr : Word;
  HP_Incr : Word;
  Stam_Incr : Word;
  Mana_Incr : Word;
end;
PExtendedInfo = ^TExtendedInfo;

TClilocItemRec = packed record
  ClilocID : Cardinal;
  Params : array of String;
end;
PClilocItemRec = ^TClilocItemRec;

TClilocItemRecArray = array of TClilocItemRec;

TClilocRec = packed record
  Count : Cardinal;
  Items : TClilocItemRecArray;
end;
PClilocRec = ^TClilocRec;

TMenuResponse = packed record
  Model : Word;
  Color : Word;
  Text :  String;
end;

TMenuResponses = array of TMenuResponse;

TGroup = packed record
  groupnumber : Integer;
  Page : Integer;
  ElemNum : Integer;
end;

TEndGroup = packed record
  groupnumber : Integer;
  Page : Integer;
  ElemNum : Integer;
end;

TPage = packed record
  Page : Integer;
  ElemNum : Integer;
end;

TMasterGump = packed record
  ID : Cardinal;
  ElemNum : Integer;
end;

TGumpButton = packed record
  x : Integer;
  y : Integer;
  released_id : Integer;
  pressed_id : Integer;
  quit : Integer;
  page_id : Integer;
  return_value : Integer;
  Page : Integer;
  ElemNum : Integer;
end;

TButtonTileArt = packed record
  x : Integer;
  y : Integer;
  released_id : Integer;
  pressed_id : Integer;
  quit : Integer;
  page_id : Integer;
  return_value : Integer;
  art_id : Integer;
  Hue : Integer;
  art_x : Integer;
  art_y : Integer;
  ElemNum : Integer;
end;

TCheckerTrans = packed record
  x : Integer;
  y : Integer;
  width : Integer;
  height : Integer;
  Page : Integer;
  ElemNum : Integer;
end;

TCroppedText = packed record
  x : Integer;
  y : Integer;
  width : Integer;
  height : Integer;
  color : Integer;
  text_id : Integer;
  Page : Integer;
  ElemNum : Integer;
end;

TGumpPic = packed record
  x : Integer;
  y : Integer;
  id : Integer;
  Hue : Integer;
  Page : Integer;
  ElemNum : Integer;
end;

TGumpPicTiled = packed record
  x : Integer;
  y : Integer;
  width : Integer;
  height : Integer;
  gump_id : Integer;
  Page : Integer;
  ElemNum : Integer;
end;

TRadio = packed record
  x : Integer;
  y : Integer;
  released_id : Integer;
  pressed_id : Integer;
  status : Integer;
  return_value : Integer;
  Page : Integer;
  ElemNum : Integer;
end;

TCheckBox = packed record
  x : Integer;
  y : Integer;
  released_id : Integer;
  pressed_id : Integer;
  status : Integer;
  return_value : Integer;
  Page : Integer;
  ElemNum : Integer;
end;

TResizePic = packed record
  x : Integer;
  y : Integer;
  gump_id : Integer;
  width : Integer;
  height : Integer;
  Page : Integer;
  ElemNum : Integer;
end;

TGumpText = packed record
  x : Integer;
  y : Integer;
  color : Integer;
  text_id : Integer;
  Page : Integer;
  ElemNum : Integer;
end;

TTextEntry = packed record
  x : Integer;
  y : Integer;
  width : Integer;
  height : Integer;
  color : Integer;
  return_value : Integer;
  default_text_id : Integer;
  Value : String;
  Page : Integer;
  ElemNum : Integer;
end;

TTilePic = packed record
  x : Integer;
  y : Integer;
  id : Integer;
  Page : Integer;
  ElemNum : Integer;
end;

TTilePicHue = packed record
  x : Integer;
  y : Integer;
  id : Integer;
  color : Integer;
  Page : Integer;
  ElemNum : Integer;
end;

TTooltip = packed record
  Cliloc_ID : Cardinal;
  Arguments : String;
  Page : Integer;
  ElemNum : Integer;
end;

THtmlGump = packed record
  x : Integer;
  y : Integer;
  width : Integer;
  height : Integer;
  text_id : Integer;
  background : Integer;
  scrollbar : Integer;
  Page : Integer;
  ElemNum : Integer;
end;

TXmfHTMLGump = packed record
  x : Integer;
  y : Integer;
  width : Integer;
  height : Integer;
  Cliloc_id : Cardinal;
  background : Integer;
  scrollbar : Integer;
  Page : Integer;
  ElemNum : Integer;
end;

TXmfHTMLGumpColor = packed record
  x : Integer;
  y : Integer;
  width : Integer;
  height : Integer;
  Cliloc_id : Cardinal;
  background : Integer;
  scrollbar : Integer;
  Hue : Integer;
  Page : Integer;
  ElemNum : Integer;
end;

TXmfHTMLTok = packed record
  x : Integer;
  y : Integer;
  width : Integer;
  height : Integer;
  background : Integer;
  scrollbar : Integer;
  Color : Integer;
  Cliloc_id : Cardinal;
  Arguments : String;
  Page : Integer;
  ElemNum : Integer;
end;

TItemProperty = packed record
  Prop : Cardinal;
  ElemNum : Integer;
end;

TUnknownItem = packed record
  CmdName : String;
  Arguments : String;
  ElemNum : Integer;
end;

TGumpInfo = packed record
  Serial : Cardinal;
  GumpID : Cardinal;
  X : SmallInt;
  Y : SmallInt;
  Pages : Integer;
  NoMove : Boolean;
  NoResize : Boolean;
  NoDispose : Boolean;
  NoClose : Boolean;
  Groups : array of TGroup;
  EndGroups : array of TEndGroup;
  GumpButtons : array of TGumpButton;
  ButtonTileArts : array of TButtonTileArt;
  CheckBoxes : array of TCheckBox;
  CheckerTrans : array of TCheckerTrans;
  CroppedText : array of TCroppedText;
  GumpPics : array of TGumpPic;
  GumpPicTiled : array of TGumpPicTiled;
  RadioButtons : array of TRadio;
  ResizePics : array of TResizePic;
  GumpText : array of TGumpText;
  TextEntries : array of TTextEntry;
  Text : array of string;
  TextEntriesLimited : array of TTextEntry;
  TilePics : array of TTilePic;
  TilePicHue : array of TTilePicHue;
  Tooltips : array of TTooltip;
  HtmlGump : array of THtmlGump;
  XmfHtmlGump : array of TXmfHtmlGump;
  XmfHTMLGumpColor : array of TXmfHTMLGumpColor;
  XmfHTMLTok : array of TXmfHTMLTok;
  ItemProperties : array of TItemProperty;
end;

TBuffIcon = packed record
  Attribute_ID : Word;
  TimeStart : TDateTime;
  Seconds : Word;
  ClilocID1 : Cardinal;
  ClilocID2 : Cardinal;
end;

TBuffArr = array of TBuffIcon;

TBuffBarInfo = packed record
  Count : Byte;
  Buffs : TBuffArr;
end;


TMultiItem = packed record
   ID : Cardinal;
   X : Word;
   Y : Word;
   Z : ShortInt;

   XMin : Word;
   XMax : Word;
   YMin : Word;
   YMax : Word;
   Width : Word;
   Height : Word;
end;
TMultiItems = array of TMultiItem;


// aliases
TPathArray = Array[0..999] of TMyPoint;
TPathArrayDyn = Array of TMyPoint;
TFoundTilesArray = Array[0..999] of TFoundTile;
TFoundTilesArrayDyn = Array of TFoundTile;
TTileDataFlagSet = set of TTileDataFlags;
{$TYPEINFO OFF}


implementation
end.