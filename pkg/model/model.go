package model

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/ghostiam/binstruct"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// Struct that represents TStatic from stealth client
// Tiles are useful to find lumberjacking, mining or fishing spots.
// There are also a lot of other useful ways to consume this struct
type StaticsXY struct {
	Tile, X, Y, Color uint16 `bin:"le"`
	Z                 byte   `bin:"le"`
}

// Multi structure that can represent wessel or house for example
type Multi struct {
	Id                                          uint32
	X, Y, XMax, XMin, YMax, YMin, Width, Height uint16 `bin:"le"`
	Z                                           byte   `bin:"le"`
}

// Determines that world cell which is a XYZ point is passable or not
type WorldCellPassable struct {
	Passable bool `bin:"len:1"`
	Z        int8 `bin:"len:2"`
}

// Point in UO world
type Point2D struct {
	X, Y uint16 `bin:"len:4"`
}

// Character buff or debuff
type BuffIcon struct {
	Attribute_ID uint16    `bin:"len:4"`
	TimeStart    time.Time `bin:"len:16"`
	Seconds      uint16    `bin:"len:4"`
	ClilocID1    uint32    `bin:"len:8"`
	ClilocID2    uint32    `bin:"len:8"`
}

// Bar info that contains all the buffs/debuffs of player's character
type BuffBarInfo struct {
	Count byte `bin:"len:1"`
	Buffs []BuffIcon
}

// Extended info about player's character
type ExtendedInfo struct {
	MaxWeight          uint16 `bin:"le"`
	Race               byte   `bin:"le"`
	StatCap            uint16 `bin:"le"`
	PetsCurrent        byte   `bin:"le"`
	PetsMax            byte   `bin:"le"`
	FireResist         uint16 `bin:"le"`
	ColdResist         uint16 `bin:"le"`
	PoisonResist       uint16 `bin:"le"`
	EnergyResist       uint16 `bin:"le"`
	Luck               int16  `bin:"le"`
	DamageMin          uint16 `bin:"le"`
	DamageMax          uint16 `bin:"le"`
	TithingPoints      uint32 `bin:"le"`
	HitChanceIncr      uint16 `bin:"le"`
	SwingSpeedIncr     uint16 `bin:"le"`
	DamageChanceIncr   uint16 `bin:"le"`
	LowerReagentCost   uint16 `bin:"le"`
	HpRegen            uint16 `bin:"le"`
	StamRegen          uint16 `bin:"le"`
	ManaRegen          uint16 `bin:"le"`
	ReflectPhysDamage  uint16 `bin:"le"`
	EnhancePotions     uint16 `bin:"le"`
	DefenseChanceIncr  uint16 `bin:"le"`
	SpellDamageIncr    uint16 `bin:"le"`
	FasterCastRecovery uint16 `bin:"le"`
	FasterCasting      uint16 `bin:"le"`
	LowerManaCost      uint16 `bin:"le"`
	StrengthIncr       uint16 `bin:"le"`
	DextIncr           uint16 `bin:"le"`
	IntIncr            uint16 `bin:"le"`
	HpIncr             uint16 `bin:"le"`
	StamIncr           uint16 `bin:"le"`
	ManaIncr           uint16 `bin:"le"`
	MaxHpIncr          uint16 `bin:"le"`
	MaxStamIncr        uint16 `bin:"le"`
	MaxManaIncrease    uint16 `bin:"le"`
}

// Represents X or Y coordinate in UO world
type Coordinate uint16

type MapCell struct {
	Tile uint16 `bin:"le"`
	Z    int8   `bin:"le"`
}

type FoundTile struct {
	Tile uint16     `bin:"le"`
	X    Coordinate `bin:"le"`
	Y    Coordinate `bin:"le"`
	Z    int8       `bin:"le"`
}

type TargetInfo struct {
	ID   uint32     `bin:"le"`
	Tile uint16     `bin:"le"`
	X    Coordinate `bin:"le"`
	Y    Coordinate `bin:"le"`
	Z    int8       `bin:"le"`
}

// TGumpInfo
/*
class _Group:
    args = [_int] * 3
    container = 'groups'
    keys = 'GroupNumber', 'Page', 'ElemNum'
*/

type Group struct {
	GroupNumber int32 `bin:"le"`
	Page        int32 `bin:"le"`
	ElemNum     int32 `bin:"le"`
}

/*
class _EndGroup(_Group):
    container = 'EndGroups'
*/

type GumpButton struct {
	X           int32 `bin:"le"`
	Y           int32 `bin:"le"`
	ReleasedID  int32 `bin:"le"`
	PressedID   int32 `bin:"le"`
	Quit        int32 `bin:"le"`
	PageID      int32 `bin:"le"`
	ReturnValue int32 `bin:"le"`
	Page        int32 `bin:"le"`
	ElemNum     int32 `bin:"le"`
}

/*
class _GumpButton:

    args = [_int] * 9
    container = 'GumpButtons'
    keys = ('X', 'Y', 'ReleasedID', 'PressedID', 'Quit', 'PageID',returnValue', 'Page', 'ElemNum')
*/
type ButtonTileArt struct {
	X           int32 `bin:"le"`
	Y           int32 `bin:"le"`
	ReleasedID  int32 `bin:"le"`
	PressedID   int32 `bin:"le"`
	Quit        int32 `bin:"le"`
	PageID      int32 `bin:"le"`
	ReturnValue int32 `bin:"le"`
	ArtID       int32 `bin:"le"`
	Hue         int32 `bin:"le"`
	ArtX        int32 `bin:"le"`
	ArtY        int32 `bin:"le"`
	ElemNum     int32 `bin:"le"`
}

/*
class _ButtonTileArt:
    args = [_int] * 12
    container = 'ButtonTileArts'
    keys = ('X', 'Y', 'ReleasedID', 'PressedID', 'Quit', 'PageID',returnValue', 'ArtID', 'Hue', 'ArtX', 'ArtY', 'ElemNum')
*/
type CheckBox struct {
	X           int32 `bin:"le"`
	Y           int32 `bin:"le"`
	ReleasedID  int32 `bin:"le"`
	PressedID   int32 `bin:"le"`
	Quit        int32 `bin:"le"`
	PageID      int32 `bin:"le"`
	ReturnValue int32 `bin:"le"`
	Page        int32 `bin:"le"`
	ElemNum     int32 `bin:"le"`
}

/*
class _CheckBox:
    args = [_int] * 8
    container = 'CheckBoxes'returnValue',
            'Page', 'ElemNum')
*/
type CheckerTrans struct {
	X       int32 `bin:"le"`
	Y       int32 `bin:"le"`
	Width   int32 `bin:"le"`
	Height  int32 `bin:"le"`
	Page    int32 `bin:"le"`
	ElemNum int32 `bin:"le"`
}

/*
class _ChekerTrans:
    args = [_int] * 6
    container = 'ChekerTrans'
    keys = 'X', 'Y', 'Width', 'Height', 'Page', 'ElemNum'
*/
type CroppedText struct {
	X       int32 `bin:"le"`
	Y       int32 `bin:"le"`
	Width   int32 `bin:"le"`
	Height  int32 `bin:"le"`
	Color   int32 `bin:"le"`
	TextId  int32 `bin:"le"`
	Page    int32 `bin:"le"`
	ElemNum int32 `bin:"le"`
}

/*
class _CroppedText:
    args = [_int] * 8
    container = 'CroppedText'
    keys = 'X', 'Y', 'Width', 'Height', 'Color', 'TextID', 'Page', 'ElemNum'
*/
type GumpPic struct {
	X       int32 `bin:"le"`
	Y       int32 `bin:"le"`
	Id      int32 `bin:"le"`
	Hue     int32 `bin:"le"`
	Page    int32 `bin:"le"`
	ElemNum int32 `bin:"le"`
}

/*
class _GumpPic:
    args = [_int] * 6
    container = 'GumpPics'
    keys = 'X', 'Y', 'ID', 'Hue', 'Page', 'ElemNum'
*/
type GumpPicTiled struct {
	X       int32 `bin:"le"`
	Y       int32 `bin:"le"`
	Width   int32 `bin:"le"`
	Height  int32 `bin:"le"`
	GumpId  int32 `bin:"le"`
	Page    int32 `bin:"le"`
	ElemNum int32 `bin:"le"`
}

/*
class _GumpPicTiled:
    fmt = '=7i'
    args = [_int] * 7
    container = 'GumpPicTiled'
    keys = 'X', 'Y', 'Width', 'Height', 'GumpID', 'Page', 'ElemNum'
*/
type RadioButton struct {
	X           int32 `bin:"le"`
	Y           int32 `bin:"le"`
	ReleasedID  int32 `bin:"le"`
	PressedID   int32 `bin:"le"`
	Status      int32 `bin:"le"`
	ReturnValue int32 `bin:"le"`
	Page        int32 `bin:"le"`
	ElemNum     int32 `bin:"le"`
}

/*
class _Radiobutton:
    args = [_int] * 8
    container = 'RadioButtons'returnValue',
            'Page', 'ElemNum')

*/
type ResizePic struct {
	X       int32 `bin:"le"`
	Y       int32 `bin:"le"`
	Width   int32 `bin:"le"`
	Height  int32 `bin:"le"`
	GumpId  int32 `bin:"le"`
	Page    int32 `bin:"le"`
	ElemNum int32 `bin:"le"`
}

/*
class _ResizePic:
    args = [_int] * 7
    container = 'ResizePics'
    keys = 'X', 'Y', 'GumpID', 'Width', 'Height', 'Page', 'ElemNum'
*/
type GumpText struct {
	X       int32 `bin:"le"`
	Y       int32 `bin:"le"`
	Color   int32 `bin:"le"`
	TextId  int32 `bin:"le"`
	Page    int32 `bin:"le"`
	ElemNum int32 `bin:"le"`
}

/*
class _GumpText:
    args = [_int] * 6
    container = 'GumpText'
    keys = 'X', 'Y', 'Color', 'TextID', 'Page', 'ElemNum'
*/
type TextEntry struct {
	X             int32 `bin:"le"`
	Y             int32 `bin:"le"`
	Width         int32 `bin:"le"`
	Height        int32 `bin:"le"`
	Color         int32 `bin:"le"`
	ReturnValue   int32 `bin:"le"`
	DefaultTextId int32 `bin:"le"`
	Page          int32 `bin:"le"`
	ElemNum       int32 `bin:"le"`
}

/*
class _TextEntry:
    args = [_int] * 7 + [_str, _int, _int]
    container = 'TextEntries'returnValue',
            'DefaultTextID', 'RealValue', 'Page', 'ElemNum')
*/
type Text struct {
	_ []string
}

/*
class _Text:
    args = [_str]
    container = 'Text'
    keys = None
*/
type TextEntryLimited struct {
	X             int32 `bin:"le"`
	Y             int32 `bin:"le"`
	Width         int32 `bin:"le"`
	Height        int32 `bin:"le"`
	Color         int32 `bin:"le"`
	ReturnValue   int32 `bin:"le"`
	DefaultTextId int32 `bin:"le"`
	Limit         int32 `bin:"le"`
	Page          int32 `bin:"le"`
	ElemNum       int32 `bin:"le"`
}

/*
class _TextEntryLimited:
    args = [_int] * 10
    container = 'TextEntriesLimited'returnValue',
            'DefaultTextID', 'Limit', 'Page', 'ElemNum')
*/
type TilePic struct {
	X       int32 `bin:"le"`
	Y       int32 `bin:"le"`
	Id      int32 `bin:"le"`
	Page    int32 `bin:"le"`
	ElemNum int32 `bin:"le"`
}

/*
class _TilePic:
    args = [_int] * 5
    container = 'TilePics'
    keys = 'X', 'Y', 'ID', 'Page', 'ElemNum'
*/
type TilePicHue struct {
	X       int32 `bin:"le"`
	Y       int32 `bin:"le"`
	Id      int32 `bin:"le"`
	Color   int32 `bin:"le"`
	Page    int32 `bin:"le"`
	ElemNum int32 `bin:"le"`
}

/*
class _TilePicHue:
    args = [_int] * 6
    container = 'TilePicHue'
    keys = 'X', 'Y', 'ID', 'Color', 'Page', 'ElemNum'
*/
type Tooltip struct {
	ClilocId int32 `bin:"le"`
	Page     int32 `bin:"le"`
	ElemNum  int32 `bin:"le"`
}

/*
class _Tooltip:
    args = [_uint, _str, _int, _int]
    container = 'Tooltips'
    keys = 'ClilocID', 'Arguments', 'Page', 'ElemNum'
*/
type HtmlGump struct {
	X          int32 `bin:"le"`
	Y          int32 `bin:"le"`
	Width      int32 `bin:"le"`
	Height     int32 `bin:"le"`
	TextId     int32 `bin:"le"`
	Background int32 `bin:"le"`
	Scrollbar  int32 `bin:"le"`
	Page       int32 `bin:"le"`
	ElemNum    int32 `bin:"le"`
}

/*
class _HtmlGump:
    args = [_int] * 9
    container = 'HtmlGump'
    keys = ('X', 'Y', 'Width', 'Height', 'TextID', 'Background', 'Scrollbar',
            'Page', 'ElemNum')
*/
type XmfHtmlGump struct {
	X          int32 `bin:"le"`
	Y          int32 `bin:"le"`
	Width      int32 `bin:"le"`
	Height     int32 `bin:"le"`
	ClilocID   int32 `bin:"le"`
	Background int32 `bin:"le"`
	Scrollbar  int32 `bin:"le"`
	Page       int32 `bin:"le"`
	ElemNum    int32 `bin:"le"`
}

/*
class _XmfHtmlGump:
    args = [_int] * 4 + [_uint] + [_int] * 4
    container = 'XmfHtmlGump'
    keys = ('X', 'Y', 'Width', 'Height', 'ClilocID', 'Background', 'Scrollbar',
            'Page', 'ElemNum')
*/
type XmfHTMLGumpColor struct {
	X          int32 `bin:"le"`
	Y          int32 `bin:"le"`
	Width      int32 `bin:"le"`
	Height     int32 `bin:"le"`
	ClilocID   int32 `bin:"le"`
	Background int32 `bin:"le"`
	Scrollbar  int32 `bin:"le"`
	Hue        int32 `bin:"le"`
	Page       int32 `bin:"le"`
	ElemNum    int32 `bin:"le"`
}

/*
class _XmfHTMLGumpColor:
    args = [_int] * 4 + [_uint] + [_int] * 5
    container = 'XmfHTMLGumpColor'
    keys = ('X', 'Y', 'Width', 'Height', 'ClilocID', 'Background', 'Scrollbar',
            'Hue', 'Page', 'ElemNum')
*/
type XmfHTMLTok struct {
	X          int32 `bin:"le"`
	Y          int32 `bin:"le"`
	Width      int32 `bin:"le"`
	Height     int32 `bin:"le"`
	Background int32 `bin:"le"`
	Scrollbar  int32 `bin:"le"`
	Color      int32 `bin:"le"`
	ClilocID   int32 `bin:"le"`
	Page       int32 `bin:"le"`
	ElemNum    int32 `bin:"le"`
}

/*
class _XmfHTMLTok:
    args = [_int] * 7 + [_uint, _str, _int, _int]
    container = 'XmfHTMLTok'
    keys = ('X', 'Y', 'Width', 'Height', 'Background', 'Scrollbar', 'Color',
            'ClilocID', 'Arguments', 'Page', 'ElemNum')
*/
type ItemProperty struct {
	Prop    uint32 `bin:"le"`
	ElemNum int32  `bin:"le"`
}

/*
class _ItemProperty:
    args = [_uint, _int]
    container = 'ItemProperties'
    keys = 'Prop', 'ElemNum'
*/
type Gump struct {
	Serial    uint32 `bin:"le"`
	GumpId    uint32 `bin:"le"`
	X         uint16 `bin:"le"`
	Y         uint16 `bin:"le"`
	Pages     int32  `bin:"le"`
	NoMove    bool   `bin:"le"`
	NoResize  bool   `bin:"le"`
	NoDispose bool   `bin:"le"`
	NoClose   bool   `bin:"le"`
	ExtInfo   ExtGumpInfo
}

type ExtGumpInfo struct {
	GroupsLen           uint32             `bin:"le"`
	Groups              []Group            `bin:"len:GroupsLen"`
	EndGroupsLen        uint32             `bin:"le"`
	EndGroups           []Group            `bin:"len:EndGroupsLen"`
	GumpButtonsLen      uint32             `bin:"le"`
	GumpButtons         []GumpButton       `bin:"len:GumpButtonsLen"`
	ButtonTileArtLen    uint32             `bin:"le"`
	ButtonTileArt       []ButtonTileArt    `bin:"len:ButtonTileArtLen"`
	CheckBoxLen         uint32             `bin:"le"`
	CheckBox            []CheckBox         `bin:"len:CheckBoxLen"`
	CheckerTransLen     uint32             `bin:"le"`
	CheckerTransparency []CheckerTrans     `bin:"len:CheckerTransLen"`
	CroppedTextLen      uint32             `bin:"le"`
	CroppedText         []CroppedText      `bin:"len:CroppedTextLen"`
	GumpPicLen          uint32             `bin:"le"`
	GumpPic             []GumpPic          `bin:"len:GumpPicLen"`
	GumpPicTiledLen     uint32             `bin:"le"`
	GumpPicTiled        []GumpPicTiled     `bin:"len:GumpPicTiledLen"`
	RadioButtonLen      uint32             `bin:"le"`
	RadioButton         []RadioButton      `bin:"len:RadioButtonLen"`
	ResizePicLen        uint32             `bin:"le"`
	ResizePics          []ResizePic        `bin:"len:ResizePicLen"`
	GumpTextLen         uint32             `bin:"le"`
	GumpText            []GumpText         `bin:"len:GumpTextLen"`
	TextEntryLen        uint32             `bin:"le"`
	TextEntry           []TextEntry        `bin:"len:TextEntryLen"`
	TextLen             uint32             `bin:"le"`
	Text                []string           `bin:"len:TextLen,[StringFunc]"`
	TextEntryLimitedLen uint32             `bin:"le"`
	TextEntryLimited    []TextEntryLimited `bin:"len:TextEntryLimitedLen"`
	TilePicLen          uint32             `bin:"le"`
	TilePic             []TilePic          `bin:"len:TilePicLen"`
	TilePicHueLen       uint32             `bin:"le"`
	TilePicture         []TilePicHue       `bin:"len:TilePicHueLen"`
	TooltipLen          uint32             `bin:"le"`
	Tooltip             []Tooltip          `bin:"len:TooltipLen"`
	HtmlGumpLen         uint32             `bin:"le"`
	HtmlGump            []HtmlGump         `bin:"len:HtmlGumpLen"`
	XmfHtmlGumpLen      uint32             `bin:"le"`
	XmfHTMLGump         []XmfHtmlGump      `bin:"len:XmfHtmlGumpLen"`
	XmfHTMLGumpColorLen uint32             `bin:"le"`
	XmfHTMLGumpColor    []XmfHTMLGumpColor `bin:"len:XmfHTMLGumpColorLen"`
	XmfHTMLTokLen       uint32             `bin:"le"`
	XmfHTMLTok          []XmfHTMLTok       `bin:"len:XmfHTMLTokLen"`
	ItemPropertyLen     uint32             `bin:"le"`
	ItemProperty        []ItemProperty     `bin:"len:GroupsLen"`
}

func (d *ExtGumpInfo) StringFunc(r binstruct.Reader) (string, error) {
	lenStr, err := r.ReadUint32()
	if err != nil {
		return "", err
	}

	_, str, err := r.ReadBytes(int(lenStr))
	if err != nil {
		return "", err
	}

	return decodeUtf16(str), nil
}

//TODO: needs refactoring to own package!!
func decodeUtf16(inputBytes []byte) string {
	win16be := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	// Make a transformer that is like win16be, but abides by BOM:
	utf16bom := unicode.BOMOverride(win16be.NewDecoder())

	// Make a Reader that uses utf16bom:
	unicodeReader := transform.NewReader(bytes.NewReader(inputBytes), utf16bom)

	// decode and print:
	decoded, err := ioutil.ReadAll(unicodeReader)
	if err != nil {
		log.Fatal("Failed to parse string from packet...")
	}

	return strings.Replace(string(decoded), "\r\n", "\n", -1)
}

/*
class _Gump:
    fmt = '<2I2hi4?'
    args = [_uint, _uint, _short, _short, _int] + [_bool] * 4
    keys = ('Serial', 'GumpID', 'X', 'Y', 'Pages', 'NoMove', 'NoResize',
            'NoDispose', 'NoClose')
*/

/**
Serial: 1069244
GumpID: B3423A1D
X: 0000
Y: 0000
Pages: 4
Gump Options:

ResizePics: X   Y   ID   Width   Height   Page   ElemNum
0:        50  25  2600  540  430  0  0

HTMLGumps: X   Y   Width   Height   TextID   Background   scrollbar   Page   ElemNum
0:        110  90  450  74  0  1  1  1  6
1:        110  170  450  74  1  1  1  1  8
2:        110  250  450  74  2  1  1  1  10
3:        110  330  450  74  3  1  1  1  12
4:        110  90  450  74  4  1  1  2  15
5:        110  170  450  74  5  1  1  2  17
6:        110  250  450  74  6  1  1  2  19
7:        110  330  450  74  7  1  1  2  21

XmfHtmlGump: X   Y   Width   Height   ClilocID   Background   scrollbar   Page   ElemNum   ClilocText
0:        150  50  360  40  1001002  0  0  0  2  <CENTER><U>Ultima Online Help Menu</U></CENTER>
1:        110  90  450  145  1062572  1  1  3  24  <U><CENTER>Another player is harassing me (or Exploiting).</CENTER></U><BR>VERBAL HARASSMENT<BR>Use this option when another player is verbally harassing your character. Verbal harassment behaviors include but are not limited to, using bad language, threats etc.. Before you submit a complaint be sure you understand what constitutes harassment <A HREF="http://uo.custhelp.com/cgi-bin/uo.cfg/php/enduser/std_adp.php?p_faqid=40">– what is verbal harassment? -</A> and that you have followed these steps:<BR>1. You have asked the player to stop and they have continued.<BR>2. You have tried to remove yourself from the situation.<BR>3. You have done nothing to instigate or further encourage the harassment.<BR>4. You have added the player to your ignore list. <A HREF="http://uo.custhelp.com/cgi-bin/uo.cfg/php/enduser/std_adp.php?p_faqid=138">- How do I ignore a player?</A><BR>5. You have read and understand Origin’s definition of harassment.<BR>6. Your account information is up to date. (Including a current email address)<BR>*If these steps have not been taken, GMs may be unable to take action against the offending player.<BR>**A chat log will be review by a GM to assess the validity of this complaint. Abuse of this system is a violation of the Rules of Conduct.<BR>EXPLOITING<BR>Use this option to report someone who may be exploiting or cheating. <A HREF="http://uo.custhelp.com/cgi-bin/uo.cfg/php/enduser/std_adp.php?p_faqid=41">– What constitutes an exploit?</a>
2:        110  240  450  145  1062573  1  1  3  26  <U><CENTER>Another player is harassing me using game mechanics.</CENTER></U><BR><BR>PHYSICAL HARASSMENT<BR>Use this option when another player is harassing your character using game mechanics. Physical harassment includes but is not limited to luring and any act that causes a players death in Trammel. Before you submit a complaint be sure you understand what constitutes harassment <A HREF="http://uo.custhelp.com/cgi-bin/uo.cfg/php/enduser/std_adp.php?p_faqid=59">- What is physical harassment? -</A> and that you have followed these steps:<BR>1. You have asked the player to stop and they have continued.<BR>2. You have tried to remove yourself from the situation.<BR>3. You have done nothing to instigate or further encourage the harassment.<BR>4. You have added the player to your ignore list. <A HREF="http://uo.custhelp.com/cgi-bin/uo.cfg/php/enduser/std_adp.php?p_faqid=138">- How do I ignore a player? -</A><BR>5. You have read and understand Origin’s definition of harassment.<BR>6. Your account information is up to date. (Including a current email address)<BR>*If these steps have not been taken, GMs may be unable to take action against the offending player.<BR>**This issue will be reviewed by a GM to assess the validity of this complaint. Abuse of this system is a violation of the Rules of Conduct.
3:        180  390  335  40  1001015  0  0  3  28  NO  - I meant to ask for help with another matter.

Text Lines:
<u>General question about Ultima Online.</u> Select this option if you have a general gameplay question, need help learning to use a skill, or if you would like to search the UO Knowledge Base.
<u>My character is physically stuck in the game.</u> This choice only covers cases where your character is physically stuck in a location they cannot move out of. This option will only work two times in 24 hours.
<u>Another player is harassing me.</u> Another player is verbally harassing your character. When you select this option you will be sending a text log to Origin Systems. To see what constitutes harassment please visit http://support.uo.com/gm_9.html.
<u>Other.</u> If you are experiencing a problem in the game that does not fall into one of the other categories or is not addressed on the Support web page (located at http://support.uo.com), please use this option.
<u>Report a bug or contact Origin.</u> Use this option to launch your web browser and mail in a bug report. Your report will be read by our Quality Assurance Staff. We apologize for not being able to reply to individual reports.
<u>Suggestion for the Game.</u> If you'd like to make a suggestion for the game, it should be directed to the Development Team Members who participate in the discussion forums on the UO.Com web site. Choosing this option will take you to the Discussion Forums.
<u>Account Management</u> For questions regarding your account such as forgotten passwords, payment options, account activation, and account transfer, please choose this option.
<u>Other.</u> If you are experiencing a problem in the game that does not fall into one of the other categories or is not addressed on the Support web page (located at http://support.uo.com), and requires in-game assistance, use this option.

GumpButtons: X   Y   Released_ID  Pressed_ID   Quit   Page_ID   Return_value   Page   ElemNum
0:        425  415  2073  2072  1  0  0  0  3
1:        80  90  5540  5541  1  2  1  1  5
2:        80  170  5540  5541  1  0  2  1  7
3:        80  250  5540  5541  0  3  0  1  9
4:        80  330  5540  5541  0  2  0  1  11
5:        80  90  5540  5541  1  0  3  2  14
6:        80  170  5540  5541  1  0  4  2  16
7:        80  250  5540  5541  1  0  5  2  18
8:        80  330  5540  5541  1  0  6  2  20
9:        80  90  5540  5541  1  0  7  3  23
10:        80  240  5540  5541  1  0  8  3  25
11:        150  390  5540  5541  0  1  0  3  27

Previous gump reply:
  Button id: 0
*/
