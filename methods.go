package gostealthclient

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	n "github.com/drabadan/gostealthclient/internal/network"
	c "github.com/drabadan/gostealthclient/pkg/constants"
	m "github.com/drabadan/gostealthclient/pkg/model"
)

func Connected() <-chan bool {
	p := n.NewBoolPacket(c.SCGetConnectedStatus)
	return p.Out
}

func AddToSystemJournal(text string) {
	n.NewVoidPacket(c.SCAddToSystemJournal, text)
}

/**
 TODO: Fix

func GetStealthInfo() <-chan StealthClientInfo {
	p := n.NewStealthClientInfoPacket()
	go p.transform()
return p.Out
}
*/

func Connect() {
	n.NewVoidPacket(c.SCConnect)
}
func Disconnect() {
	n.NewVoidPacket(c.SCDisconnect)
}
func SetPauseScriptOnDisconnectStatus(value bool) {
	n.NewVoidPacket(c.SCSetPauseScriptOnDisconnectStatus, value)
}
func GetPauseScriptOnDisconnectStatus() <-chan bool {
	p := n.NewBoolPacket(c.SCGetPauseScriptOnDisconnectStatus)
	return p.Out
}
func SetARStatus(value bool) {
	n.NewVoidPacket(c.SCSetARStatus, value)
}
func GetARStatus() <-chan bool {
	p := n.NewBoolPacket(c.SCGetARStatus)
	return p.Out
}
func CharName() <-chan string {
	p := n.NewStringPacket(c.SCGetCharName)
	return p.Out
}
func ChangeProfile(profileName string) <-chan int32 {
	p := n.NewIntPacket(c.SCChangeProfile, profileName)
	return p.Out
}
func ChangeProfileEx(pName, shardName, charName string) <-chan uint16 {
	p := n.NewUint16Packet(c.SCChangeProfileEx, pName, shardName, charName)
	return p.Out
}
func ProfileName() <-chan string {
	p := n.NewStringPacket(c.SCGetProfileName)
	return p.Out
}
func Self() <-chan uint32 {
	p := n.NewUint32Packet(c.SCGetSelfID)
	return p.Out
}
func Sex() <-chan byte {
	p := n.NewBytePacket(c.SCGetSelfSex)
	return p.Out
}
func GetCharTitle() <-chan string {
	p := n.NewStringPacket(c.SCGetCharTitle)
	return p.Out
}
func Gold() <-chan uint16 {
	p := n.NewUint16Packet(c.SCGetSelfGold)
	return p.Out
}
func Armor() <-chan uint16 {
	p := n.NewUint16Packet(c.SCGetSelfArmor)
	return p.Out
}
func Weight() <-chan uint16 {
	p := n.NewUint16Packet(c.SCGetSelfWeight)
	return p.Out
}
func MaxWeight() <-chan uint16 {
	p := n.NewUint16Packet(c.SCGetSelfMaxWeight)
	return p.Out
}
func WorldNum() <-chan byte {
	p := n.NewBytePacket(c.SCGetWorldNum)
	return p.Out
}
func Race() <-chan byte {
	p := n.NewBytePacket(c.SCGetSelfRace)
	return p.Out
}
func MaxPets() <-chan byte {
	p := n.NewBytePacket(c.SCGetSelfPetsMax)
	return p.Out
}
func PetsCurrent() <-chan byte {
	p := n.NewBytePacket(c.SCGetSelfPetsCurrent)
	return p.Out
}
func FireResist() <-chan uint16 {
	p := n.NewUint16Packet(c.SCGetSelfFireResist)
	return p.Out
}
func ColdResist() <-chan uint16 {
	p := n.NewUint16Packet(c.SCGetSelfColdResist)
	return p.Out
}
func PoisonResist() <-chan uint16 {
	p := n.NewUint16Packet(c.SCGetSelfPoisonResist)
	return p.Out
}
func EnergyResist() <-chan uint16 {
	p := n.NewUint16Packet(c.SCGetSelfEnergyResist)
	return p.Out
}

func ConnectedTime() <-chan time.Time {
	p := n.NewTimePacket(c.SCGetConnectedTime)
	return p.Out
}

func DisconnectedTime() <-chan time.Time {
	p := n.NewTimePacket(c.SCGetDisconnectedTime)
	return p.Out
}

func LastContainer() <-chan uint32 {
	p := n.NewUint32Packet(c.SCGetLastContainer)
	return p.Out
}

func LastTarget() <-chan uint32 {
	p := n.NewUint32Packet(c.SCGetLastTarget)
	return p.Out
}

func LastAttack() <-chan uint32 {
	p := n.NewUint32Packet(c.SCGetLastAttack)
	return p.Out
}

func LastStatus() <-chan uint32 {
	p := n.NewUint32Packet(43)
	return p.Out
}

func LastObject() <-chan uint32 {
	p := n.NewUint32Packet(44)
	return p.Out
}

func GetBuffBarInfo() <-chan m.BuffBarInfo {
	p := n.NewBuffBarInfo()
	return p.Out
}

func ShardName() <-chan string {
	p := n.NewStringPacket(47)
	return p.Out
}

func ProfileShardName() <-chan string {
	p := n.NewStringPacket(343)
	return p.Out
}

func ProxyIP() <-chan string {
	p := n.NewStringPacket(60)
	return p.Out
}

func ProxyPort() <-chan uint16 {
	p := n.NewUint16Packet(c.SCGetProxyPort)
	return p.Out
}

func UseProxy() <-chan bool {
	p := n.NewBoolPacket(62)
	return p.Out
}
func Backpack() <-chan uint32 {
	p := n.NewUint32Packet(48)
	return p.Out
}

func Str() <-chan uint32 {
	p := n.NewUint32Packet(49)
	return p.Out
}

func Int() <-chan uint32 {
	p := n.NewUint32Packet(50)
	return p.Out
}

func Dex() <-chan uint32 {
	p := n.NewUint32Packet(51)
	return p.Out
}

func Life() <-chan uint32 {
	p := n.NewUint32Packet(52)
	return p.Out
}

func HP() <-chan uint32 {
	return Life()
}

func Mana() <-chan uint32 {
	p := n.NewUint32Packet(53)
	return p.Out
}

func Stam() <-chan uint32 {
	p := n.NewUint32Packet(54)
	return p.Out
}

func MaxLife() <-chan uint32 {
	p := n.NewUint32Packet(55)
	return p.Out
}

func MaxHP() <-chan uint32 {
	return MaxLife()
}

func MaxMana() <-chan uint32 {
	p := n.NewUint32Packet(56)
	return p.Out
}

func MaxStam() <-chan uint32 {
	p := n.NewUint32Packet(57)
	return p.Out
}

func Luck() <-chan uint16 {
	p := n.NewUint16Packet(c.SCGetSelfLuck)
	return p.Out
}

func GetExtInfo() <-chan m.ExtendedInfo {
	p := n.NewGetExtInfoPacket()
	return p.Out
}

func Hidden() <-chan bool {
	p := n.NewBoolPacket(63)
	return p.Out
}

func Poisoned() <-chan bool {
	p := n.NewBoolPacket(64)
	return p.Out
}

func Paralyzed() <-chan bool {
	p := n.NewBoolPacket(65)
	return p.Out
}

func Dead() <-chan bool {
	p := n.NewBoolPacket(66)
	return p.Out
}

func WarMode() <-chan bool {
	p := n.NewBoolPacket(171, <-Self())
	return p.Out
}

func WarTargetID() <-chan uint32 {
	p := n.NewUint32Packet(c.SCGetWarTarget)
	return p.Out
}

func SetWarMode(value bool) {
	n.NewVoidPacket(c.SCSetWarMode, value)
}

func Attack(attackedID uint32) {
	n.NewVoidPacket(c.SCAttack, attackedID)
}

func UseSelfPaperdollScroll() {
	n.NewVoidPacket(c.SCUseSelfPaperdollScroll)
}

func UseOtherPaperdollScroll(oid uint32) {
	n.NewVoidPacket(c.SCUseOtherPaperdollScroll, oid)
}

func TargetID() <-chan uint32 {
	p := n.NewUint32Packet(72)
	return p.Out
}

func CancelTarget() {
	n.NewVoidPacket(73)
}

func TargetToObject(ObjectID uint32) {
	n.NewVoidPacket(74, ObjectID)
}

func TargetToXYZ(x, y uint16, z byte) {
	n.NewVoidPacket(c.SCTargetToXYZ, x, y, z)
}

func TargetToTile(tileModel, x, y uint16, z byte) {
	n.NewVoidPacket(c.SCTargetToTile, x, y, z)
}

func WaitTargetObject(ObjID uint32) {
	n.NewVoidPacket(77, ObjID)
}

func WaitTargetTile(tile, x, y uint16, z byte) {
	n.NewVoidPacket(c.SCWaitTargetTile, tile, x, y, z)
}

func WaitTargetXYZ(x, y uint16, z byte) {
	n.NewVoidPacket(c.SCWaitTargetXYZ, x, y, z)
}

func WaitTargetSelf() {
	n.NewVoidPacket(80)
}

func WaitTargetType(ObjType uint16) {
	n.NewVoidPacket(81, ObjType)
}

func CancelWaitTarget() {
	n.NewVoidPacket(82)
}

func WaitTargetGround(ObjType uint16) {
	n.NewVoidPacket(83, ObjType)
}

func WaitTargetLast() {
	n.NewVoidPacket(84)
}

func UsePrimaryAbility() {
	n.NewVoidPacket(85)
}

func UseSecondaryAbility() {
	n.NewVoidPacket(86)
}

func GetActiveAbility() <-chan string {
	p := n.NewStringPacket(87)
	return p.Out
}

func ToggleFly() {
	n.NewVoidPacket(88)
}

func getSkillId(skillName string) <-chan uint32 {
	p := n.NewUint32Packet(c.SCGetSkillID, skillName)
	return p.Out
}

func UseSkill(skillName string) {
	n.NewVoidPacket(c.SCUseSkill, <-getSkillId(skillName))
}

func ChangeSkillLockState(skillName string, skillState byte) {
	n.NewVoidPacket(c.SCChangeSkillLockState, <-getSkillId(skillName), skillState)
}

func SetSkillLockState(skillName string, skillState byte) {
	ChangeSkillLockState(skillName, skillState)
}

func GetSkillCap(SkillName string) <-chan float64 {
	p := n.NewFloatPacket(92, <-getSkillId(SkillName))
	return p.Out
}
func GetSkillValue(SkillName string) <-chan float64 {
	p := n.NewFloatPacket(93, <-getSkillId(SkillName))
	return p.Out
}
func GetSkillCurrentValue(SkillName string) <-chan float64 {
	p := n.NewFloatPacket(351, <-getSkillId(SkillName))
	return p.Out
}

func ReqVirtuesGump() {
	n.NewVoidPacket(94)
}

func UseVirtue(VirtueName string) {
	if v, ok := c.VIRTUES[VirtueName]; ok {
		n.NewVoidPacket(95, v)
	} else {
		log.Fatalf("Unknown virtue %v", VirtueName)
	}
}

func Cast(spellName string) {
	n.NewVoidPacket(c.SCCastSpell, c.SPELLS[strings.ToLower(spellName)])
}

func CastToObj(spellName string, oid uint32) {
	WaitTargetObject(oid)
	Cast(spellName)
}

func IsActiveSpellAbility(spellName string) <-chan bool {
	p := n.NewBoolPacket(c.SCIsActiveSpellAbility, c.SPELLS[strings.ToLower(spellName)])
	return p.Out
}

func UnsetCatchBag() {
	n.NewVoidPacket(100)
}

func SetCatchBag(ObjectID uint32) {
	n.NewVoidPacket(99, ObjectID)
}

func UseObject(ObjectID uint32) {
	n.NewVoidPacket(101, ObjectID)
}

func UseType(objType uint16, color uint16) <-chan uint32 {
	p := n.NewUint32Packet(c.SCUseType, objType, color)
	return p.Out
}

func UseFromGround(objType, color uint16) <-chan uint32 {
	p := n.NewUint32Packet(c.SCUseFromGround, objType, color)
	return p.Out
}

func ClickOnObject(ObjectID uint32) {
	n.NewVoidPacket(104, ObjectID)
}

func FoundedParamID() <-chan uint32 {
	p := n.NewUint32Packet(105)
	return p.Out
}

func LineID() <-chan uint32 {
	p := n.NewUint32Packet(106)
	return p.Out
}

func LineType() <-chan uint16 {
	p := n.NewUint16Packet(c.SCGetLineType)
	return p.Out
}

func LineName() <-chan string {
	p := n.NewStringPacket(114)
	return p.Out
}

func LineTime() <-chan time.Time {
	p := n.NewTimePacket(c.SCGetLineTime)
	return p.Out
}

func LineMsgType() <-chan byte {
	p := n.NewBytePacket(c.SCGetLineMsgType)
	return p.Out
}

func LineTextColor() <-chan uint16 {
	p := n.NewUint16Packet(c.SCGetLineTextColor)
	return p.Out
}

func LineTextFont() <-chan uint16 {
	p := n.NewUint16Packet(c.SCGetLineTextFont)
	return p.Out
}

func LineIndex() <-chan uint32 {
	p := n.NewUint32Packet(112)
	return p.Out
}

func LineCount() <-chan uint32 {
	p := n.NewUint32Packet(113)
	return p.Out
}

func AddJournalIgnore(str string) {
	n.NewVoidPacket(115, str)
}

func ClearJournalIgnore() {
	n.NewVoidPacket(116)
}

func AddChatUserIgnore(user string) {
	n.NewVoidPacket(117, user)
}

func AddToJournal(msg string) {
	n.NewVoidPacket(304, msg)
}

func ClearChatUserIgnore() {
	n.NewVoidPacket(118)
}

func ClearJournal() {
	n.NewVoidPacket(119)
}

func ClearSystemJournal() {
	n.NewVoidPacket(346)
}

func LastJournalMessage() <-chan string {
	p := n.NewStringPacket(120)
	return p.Out
}

func InJournal(s string) <-chan int32 {
	p := n.NewIntPacket(c.SCInJournal, s)
	return p.Out
}

// InJournalBetweenTimes
//
// RU: Поиск последней строки в журнале по слову (или по словам) во временном интервале.
// Если строка не найдена возвратит -1, если найдена,
// возвратит индекс строки в журнале начиная с 0.
//
// EN: Search for last entry in journal by word(words) in time intervalreturns if string is found - index of string starting from 0
// if string is not found -1
func InJournalBetweenTimes(str string, timeBegin time.Time, timeEnd time.Time) <-chan int32 {
	p := n.NewIntPacket(c.SCInJournalBetweenTimes, str, timeBegin, timeEnd)
	return p.Out
}

func Journal(stringIndex uint32) <-chan string {
	p := n.NewStringPacket(c.SCJournal, stringIndex)
	return p.Out
}

func SetJournalLine(stringIndex uint32, text string) {
	n.NewVoidPacket(c.SCSetJournalLine, stringIndex, text)
}

func LowJournal() <-chan uint32 {
	p := n.NewUint32Packet(125)
	return p.Out
}

func HighJournal() <-chan uint32 {
	p := n.NewUint32Packet(126)
	return p.Out
}

func waitJournalLineType(startTime time.Time, str string, maxWaitTime time.Duration, t string) <-chan bool {
	if maxWaitTime == 0 {
		maxWaitTime = time.Duration(time.Hour * 24)
	}

	stop := time.Now().Add(maxWaitTime)
	r := make(chan bool)
	defer close(r)
	go func() {
		for {
			if time.Now().After(stop) {
				r <- false
				break
			}

			if <-InJournalBetweenTimes(str, startTime, time.Now()) >= 0 {
				if t != "" && <-LineName() == "System" {
					r <- true
					break
				} else if t == "normal" {
					r <- true
				}
			}
			time.Sleep(time.Millisecond * 50)
		}
	}()
	return r
}

func WaitJournalLine(startTime time.Time, str string, maxWaitTime time.Duration) <-chan bool {
	return waitJournalLineType(startTime, str, maxWaitTime, "normal")
}

func WaitJournalLineSystem(startTime time.Time, str string, maxWaitTime time.Duration) <-chan bool {
	return waitJournalLineType(startTime, str, maxWaitTime, "System")
}

func SetFindDistance(Value uint32) {
	n.NewVoidPacket(127, Value)
}

func GetFindDistance() <-chan uint32 {
	p := n.NewUint32Packet(128)
	return p.Out
}

func SetFindVertical(Value uint32) {
	n.NewVoidPacket(129, Value)
}

func GetFindVertical() <-chan uint32 {
	p := n.NewUint32Packet(130)
	return p.Out
}

func SetFindInNulPoint(v bool) {
	n.NewVoidPacket(c.SCSetFindInNulPoint, v)
}

func GetFindInNulPoint() <-chan bool {
	p := n.NewBoolPacket(337)
	return p.Out
}

func FindTypeEx(objType, objColor uint16, container uint32, inSub bool) <-chan uint32 {
	p := n.NewUint32Packet(c.SCFindTypeEx, objType, objColor, container, inSub)
	return p.Out
}

func FindType(objType uint16, container uint32) <-chan uint32 {
	return FindTypeEx(objType, 0xffff, container, false)
}

// Ищет по заданному массиву типов ObjTypes и массиву цветов Colors в массиве контейнеров Containers.
// Если InSub включено, то поиск будет осуществляться в сабконтейнерах (рекурсивно).
// Внутри метод перебором проходит через все типы\цвета\контейнеры и выполняет обычный Findtypeex.
//
// Example 1 - ищет на земле заданные типы (животные, чаты и т.д.):
//
// FindDistance := 20;
// FindVertical := 10;
// FindTypesArrayEx([$29A, $29B, $190, $191, $25d, $25e, $192, $193, $25f, $260, $2ea, $2ec, $2ed, $84, $f6, $19, $db, $51, $7a, $2ee, $2e8, $2e9, $2eb, $117, $116, $115],[$FFFF],[Ground],false);
// AddToSystemJournal('FindCount = ' + IntToStr(FindCount));
func FindTypesArrayEx(objTypes, colors []uint16, containers []uint32, inSub bool) <-chan uint32 {
	p := n.NewUint32Packet(c.SCFindTypesArrayEx, objTypes, colors, containers, inSub)
	return p.Out
}

func FindNotoriety(objType uint16, notoriety byte) <-chan uint32 {
	p := n.NewUint32Packet(c.SCFindNotoriety, objType, notoriety)
	return p.Out
}

func FindAtCoord(x, y uint16) <-chan uint32 {
	p := n.NewUint32Packet(c.SCFindAtCoord, x, y)
	return p.Out

}
func Ignore(ObjID uint32) {
	n.NewVoidPacket(134, ObjID)
}

func IgnoreOff(ObjID uint32) {
	n.NewVoidPacket(135, ObjID)
}

func IgnoreReset() {
	n.NewVoidPacket(136)
}

func GetIgnoreList() <-chan []uint32 {
	p := n.NewUint32ArrayPacket(c.SCGetIgnoreList)
	return p.Out
}

func GetFoundList() <-chan []uint32 {
	p := n.NewUint32ArrayPacket(c.SCGetFindedList)
	return p.Out
}

func FindItem() <-chan uint32 {
	p := n.NewUint32Packet(139)
	return p.Out
}

func FindCount() <-chan uint32 {
	p := n.NewUint32Packet(140)
	return p.Out
}

func FindQuantity() <-chan uint32 {
	p := n.NewUint32Packet(141)
	return p.Out
}

func FindFullQuantity() <-chan uint32 {
	p := n.NewUint32Packet(142)
	return p.Out
}

func PredictedX() <-chan uint16 {
	p := n.NewUint16Packet(143)
	return p.Out
}

func PredictedY() <-chan uint16 {
	p := n.NewUint16Packet(144)
	return p.Out
}

func PredictedZ() <-chan byte {
	p := n.NewBytePacket(c.SCPredictedZ)
	return p.Out
}

func PredictedDirection() <-chan byte {
	p := n.NewBytePacket(c.SCPredictedDirection)
	return p.Out
}

func GetX(oid uint32) <-chan uint16 {
	p := n.NewUint16Packet(c.SCGetX, oid)
	return p.Out
}

func GetY(oid uint32) <-chan uint16 {
	p := n.NewUint16Packet(c.SCGetY, oid)
	return p.Out
}

func GetZ(oid uint32) <-chan int8 {
	p := n.NewInt8Packet(c.SCGetZ, oid)
	return p.Out
}

func GetName(oid uint32) <-chan string {
	p := n.NewStringPacket(c.SCGetName, oid)
	return p.Out
}

func GetAltName(oid uint32) <-chan string {
	p := n.NewStringPacket(c.SCGetAltName, oid)
	return p.Out

}

func GetTitle(oid uint32) <-chan string {
	p := n.NewStringPacket(c.SCGetTitle)
	return p.Out
}

func GetTooltip(oid uint32) <-chan string {
	p := n.NewStringPacket(c.SCGetCliloc, oid)
	return p.Out
}

func GetCliloc(oid uint32) <-chan string {
	p := n.NewStringPacket(c.SCGetCliloc, oid)
	return p.Out
}

func GetType(oid uint32) <-chan uint16 {
	p := n.NewUint16Packet(c.SCGetType, oid)
	return p.Out
}

/*
_get_tooltip_obj = _ScriptMethod(152)  # GetToolTipRec
_get_tooltip_obj.restype = _buffer  # Array of TClilocRec
_get_tooltip_obj.argtypes = [_uint]  # ObjID
func GetTooltipRec(ObjID){
    result = []
}
    data = _get_tooltip_obj(ObjID)
    count = _uint.from_buffer(data)
    offset = 4
    for i in range(count):
        cliloc, length = _struct.unpack_from('<iI', data, offset)
        offset += 8
        strings = []
        for j in range(length):
            string = _str.from_buffer(data, offset)
            offset += _struct.calcsize(string.fmt)
            strings.append(string.value)
        result.append({'Cliloc_ID': cliloc, 'Params': strings})return result
*/
func GetClilocByID(oid uint32) <-chan string {
	p := n.NewStringPacket(c.SCGetClilocByID, oid)
	return p.Out
}

func GetQuantity(oid uint32) <-chan int32 {
	p := n.NewIntPacket(c.SCGetQuantity, oid)
	return p.Out
}

func IsObjectExists(oid uint32) <-chan bool {
	p := n.NewBoolPacket(c.SCIsObjectExists, oid)
	return p.Out
}

func IsNPC(oid uint32) <-chan bool {
	p := n.NewBoolPacket(c.SCIsNPC, oid)
	return p.Out
}

func GetPrice(oid uint32) <-chan uint32 {
	p := n.NewUint32Packet(c.SCGetPrice, oid)
	return p.Out
}

func GetDirection(oid uint32) <-chan byte {
	p := n.NewBytePacket(c.SCGetDirection, oid)
	return p.Out
}

func GetDistance(oid uint32) <-chan int32 {
	p := n.NewIntPacket(c.SCGetDistance, oid)
	return p.Out
}

func GetColor(oid uint32) <-chan uint16 {
	p := n.NewUint16Packet(c.SCGetColor, oid)
	return p.Out
}

func GetStr(oid uint32) <-chan int32 {
	p := n.NewIntPacket(c.SCGetStr, oid)
	return p.Out
}

func GetInt(oid uint32) <-chan int32 {
	p := n.NewIntPacket(c.SCGetInt, oid)
	return p.Out
}

func GetDex(oid uint32) <-chan int32 {
	p := n.NewIntPacket(c.SCGetDex, oid)
	return p.Out
}

func GetHP(oid uint32) <-chan int32 {
	p := n.NewIntPacket(c.SCGetHP, oid)
	return p.Out
}

func GetMaxHP(oid uint32) <-chan int32 {
	p := n.NewIntPacket(c.SCGetMaxHP, oid)
	return p.Out
}

func GetMana(oid uint32) <-chan int32 {
	p := n.NewIntPacket(c.SCGetMana, oid)
	return p.Out
}

func GetMaxMana(ObjID uint32) <-chan int32 {
	p := n.NewIntPacket(166, ObjID)
	return p.Out
}

func GetStam(oid uint32) <-chan int32 {
	p := n.NewIntPacket(c.SCGetStam, oid)
	return p.Out
}
func GetMaxStam(ObjID uint32) <-chan int32 {
	p := n.NewIntPacket(168, ObjID)
	return p.Out
}

func GetNotoriety(oid uint32) <-chan byte {
	p := n.NewBytePacket(c.SCGetNotoriety)
	return p.Out
}

func GetParent(oid uint32) <-chan uint32 {
	p := n.NewUint32Packet(c.SCGetParent, oid)
	return p.Out
}

func IsWarMode(oid uint32) <-chan bool {
	p := n.NewBoolPacket(c.SCIsWarMode, oid)
	return p.Out
}

func IsDead(ObjID uint32) <-chan bool {
	p := n.NewBoolPacket(173, ObjID)
	return p.Out
}

func IsRunning(ObjID uint32) <-chan bool {
	p := n.NewBoolPacket(174, ObjID)
	return p.Out
}

func IsContainer(ObjID uint32) <-chan bool {
	p := n.NewBoolPacket(175, ObjID)
	return p.Out
}

func IsHidden(ObjID uint32) <-chan bool {
	p := n.NewBoolPacket(176, ObjID)
	return p.Out
}

func IsMovable(ObjID uint32) <-chan bool {
	p := n.NewBoolPacket(177, ObjID)
	return p.Out
}

func IsYellowHits(ObjID uint32) <-chan bool {
	p := n.NewBoolPacket(178, ObjID)
	return p.Out
}

func IsPoisoned(ObjID uint32) <-chan bool {
	p := n.NewBoolPacket(179, ObjID)
	return p.Out
}

func IsParalyzed(ObjID uint32) <-chan bool {
	p := n.NewBoolPacket(180, ObjID)
	return p.Out
}

func IsFemale(ObjID uint32) <-chan bool {
	p := n.NewBoolPacket(181, ObjID)
	return p.Out
}

func OpenDoor() {
	n.NewVoidPacket(182)
}

func Bow() {
	n.NewVoidPacket(183)
}

func Salute() {
	n.NewVoidPacket(184)
}

func GetPickupedItem() <-chan uint32 {
	p := n.NewUint32Packet(185)
	return p.Out
}

func SetPickupedItem(ID uint32) {
	n.NewVoidPacket(186, ID)
}

func GetDropCheckCoord() <-chan bool {
	p := n.NewBoolPacket(187)
	return p.Out
}

func SetDropCheckCoord(value bool) {
	n.NewVoidPacket(c.SCSetDropCheckCoord)
}

func GetDropDelay() <-chan uint32 {
	p := n.NewUint32Packet(189)
	return p.Out
}

func SetDropDelay(Value uint32) {
	n.NewVoidPacket(190, Value)
}

func DragItem(oid uint32, count int32) <-chan bool {
	p := n.NewBoolPacket(c.SCDragItem, oid, count)
	return p.Out
}

func DropItem(container uint32, x, y, z int32) <-chan bool {
	p := n.NewBoolPacket(c.SCDropItem, container, x, y, z)
	return p.Out
}

func MoveItem(itemID uint32, count int32, container uint32, x, y, z int32) <-chan bool {
	r := make(chan bool)
	go func() {
		if !<-DragItem(itemID, count) {
			r <- false
		}

		time.Sleep(time.Millisecond * 100)

		r <- <-DropItem(container, x, y, z)
	}()
	return r
}

func Grab(oid uint32, count int32) <-chan bool {
	return MoveItem(oid, count, <-Backpack(), 0, 0, 0)
}

func Drop(oid uint32, count, x, y, z int32) <-chan bool {
	return MoveItem(oid, count, Ground(), x, y, z)
}

func Ground() uint32 {
	return 0
}

func DropHere(oid uint32) <-chan bool {
	return MoveItem(oid, 0, Ground(), 0, 0, 0)
}

func MoveItems(container uint32,
	itemsType, itemsColor uint16,
	moveIntoID uint32,
	x, y, z int32,
	delayMS uint32,
	// Count of stacks not the total quantity to move
	maxCount int) <-chan bool {
	r := make(chan bool)
	go func() {
		<-FindTypeEx(itemsType, itemsColor, container, false)
		items := <-GetFoundList()

		if len(items) > 0 {
			if maxCount == 0 {
				maxCount = len(items)
			}

			for i := 0; i < maxCount; i++ {
				MoveItem(items[i], 0, moveIntoID, x, y, z)
				time.Sleep(time.Millisecond * time.Duration(delayMS))
			}
		}

		r <- true
	}()
	return r
}

func EmptyContainer(container, destContainer, delayMs uint32) <-chan bool {
	return MoveItems(container, 0xffff, 0xffff, destContainer, 0, 0, 0, delayMs, 0)
}

func RequestContextMenu(ID uint32) {
	n.NewVoidPacket(193, ID)
}

func SetContextMenuHook(menuID uint32, entryNumber byte) {
	n.NewVoidPacket(c.SCSetContextMenuHook, menuID, entryNumber)
}

/*
_get_context_menu = _ScriptMethod(195)  # GetContextMenu
_get_context_menu.restype = _buffer
func GetContextMenu(){
    result = []
}
    data = _get_context_menu()
    count = _uint.from_buffer(data)
    offset = count.size
    while 42:
        if offset >= len(data) - 1:
            break
        string = _str.from_buffer(data, offset)
        offset += string.size
        result.append(string.value)return result
_get_context_menu_record = _ScriptMethod(345)  # GetContextMenuRec
_get_context_menu_record.restype = _buffer  # TODO: What is this do?
func GetContextMenuRec(){
    """
}
    fmt = 'HH'
    data = _get_context_menu_record()
    keys = 'Tag', 'Flags'
    serial, count, tmp = _struct.unpack('>IBI', data[:9])
    l = []
    for i in range(count):
        l.append(_struct.unpack('HHIHH', data[9+i*12:9+i*12+12]))
    """return None
*/
func ClearContextMenu() {
	n.NewVoidPacket(196)
}

func IsTrade() <-chan bool {
	p := n.NewBoolPacket(197)
	return p.Out
}

func GetTradeContainer(tradeNum, num byte) <-chan uint32 {
	p := n.NewUint32Packet(c.SCGetTradeContainer, tradeNum, num)
	return p.Out

}

func GetTradeOpponent(tradeNum byte) <-chan uint32 {
	p := n.NewUint32Packet(c.SCGetTradeOpponent, tradeNum)
	return p.Out
}

func TradeCount() <-chan byte {
	p := n.NewBytePacket(c.SCGetTradeCount)
	return p.Out
}

func GetTradeOpponentName(tradeNum byte) <-chan string {
	p := n.NewStringPacket(c.SCGetTradeOpponentName, tradeNum)
	return p.Out
}

func TradeCheck(tradeNum, num byte) <-chan bool {
	p := n.NewBoolPacket(c.SCTradeCheck, tradeNum, num)
	return p.Out
}

func ConfirmTrade(tradeNum byte) {
	n.NewVoidPacket(c.SCConfirmTrade, tradeNum)
}

func CancelTrade(tradeNum byte) <-chan bool {
	p := n.NewBoolPacket(c.SCCancelTrade, tradeNum)
	return p.Out
}

func WaitMenu(menuCaption, elementCaption string) {
	n.NewVoidPacket(c.SCWaitMenu, menuCaption, elementCaption)
}

func AutoMenu(menuCaption, elementCaption string) {
	n.NewVoidPacket(c.SCAutoMenu, menuCaption, elementCaption)
}

func MenuHookPresent() <-chan bool {
	p := n.NewBoolPacket(207)
	return p.Out
}

func MenuPresent() <-chan bool {
	p := n.NewBoolPacket(208)
	return p.Out
}

func CancelMenu() {
	n.NewVoidPacket(209)
}

func CancelAllMenuHooks() {
	CancelMenu()
}

func CloseMenu() {
	n.NewVoidPacket(210)
}

/*
_get_menu = _ScriptMethod(338)  # GetMenuItems
_get_menu.restype = _buffer
_get_menu.argtypes = [_str]  # MenuCaption
func GetMenu(MenuCaption){
    result = []
}
    data = _get_menu(MenuCaption)
    count = _uint.from_buffer(data)
    offset = count.size
    while 42:
        if offset >= len(data) - 1:
            break
        string = _str.from_buffer(data, offset)
        result.append(string.value)
        offset += string.sizereturn result
func GetMenuItems(MenuCaption){
    p :=
return '\n'.join(GetMenu(MenuCaption))
}
_get_last_menu = _ScriptMethod(339)  # GetLastMenuItems
_get_last_menu.restype = _buffer
func GetLastMenu(){
    result = []
}
    data = _get_last_menu()
    count = _uint.from_buffer(data)
    offset = count.size
    while 42:
        if offset >= len(data) - 1:
            break
        string = _str.from_buffer(data, offset)
        result.append(string.value)
        offset += string.sizereturn result
func GetLastMenuItems(){
    p :=
return '\n'.join(GetLastMenu())
}
*/

func WaitGump(value int32) {
	n.NewVoidPacket(c.SCWaitGumpInt, value)
}

func WaitTextEntry(value string) {
	n.NewVoidPacket(212, value)
}

func GumpAutoTextEntry(textEntryID int32, value string) {
	n.NewVoidPacket(c.SCGumpAutoTextEntry, textEntryID, value)
}

func GumpAutoRadiobutton(radiobuttonID, value int32) {
	n.NewVoidPacket(c.SCGumpAutoRadiobutton, radiobuttonID, value)
}

func GumpAutoCheckBox(checkBoxID, value int32) {
	n.NewVoidPacket(c.SCGumpAutoCheckBox, checkBoxID, value)
}

func NumGumpButton(gumpIndex uint16, value int32) <-chan bool {
	p := n.NewBoolPacket(c.SCNumGumpButton, gumpIndex, value)
	return p.Out
}

func NumGumpTextEntry(gumpIndex uint16, textEntryID int32, value string) <-chan bool {
	p := n.NewBoolPacket(c.SCNumGumpTextEntry, gumpIndex, textEntryID, value)
	return p.Out
}

func NumGumpRadiobutton(gumpIndex uint16, radiobuttonID, value int32) <-chan bool {
	p := n.NewBoolPacket(c.SCNumGumpRadiobutton, gumpIndex, radiobuttonID, value)
	return p.Out
}
func NumGumpCheckBox(gumpIndex uint16, checkBoxID, value int32) <-chan bool {
	p := n.NewBoolPacket(c.SCNumGumpCheckBox, gumpIndex, checkBoxID, value)
	return p.Out
}

func GetGumpsCount() <-chan uint16 {
	p := n.NewUint16Packet(220)
	return p.Out
}

func CloseSimpleGump(GumpIndex uint16) {
	n.NewVoidPacket(221, GumpIndex)
}

func IsGump() bool {
	return <-GetGumpsCount() > 0
}

func GetGumpSerial(gumpIndex uint16) <-chan uint32 {
	p := n.NewUint32Packet(c.SCGetGumpSerial, gumpIndex)
	return p.Out
}

func GetGumpID(gumpIndex uint16) <-chan uint32 {
	p := n.NewUint32Packet(c.SCGetGumpID, gumpIndex)
	return p.Out
}

func IsGumpCanBeClosed(gumpIndex uint16) <-chan bool {
	p := n.NewBoolPacket(c.SCGetGumpNoClose, gumpIndex)
	return p.Out
}

/*
_get_gump_text = _ScriptMethod(225)  # GetGumpTextLines
_get_gump_text.restype = _buffer
_get_gump_text.argtypes = [_ushort]  # GumpIndex
func GetGumpTextLines(GumpIndex){
    result = []
}
    data = _get_gump_text(GumpIndex)
    count = _uint.from_buffer(data)
    offset = count.size
    while 42:
        if offset >= len(data) - 1:
            break
        string = _str.from_buffer(data, offset)
        offset += string.size
        result.append(string.value)return result
_get_gump_full_lines = _ScriptMethod(226)  # GetGumpFullLines
_get_gump_full_lines.restype = _buffer
_get_gump_full_lines.argtypes = [_ushort]  # GumpIndex
func GetGumpFullLines(GumpIndex){
    result = []
}
    data = _get_gump_full_lines(GumpIndex)
    count = _uint.from_buffer(data)
    offset = count.size
    while 42:
        if offset >= len(data) - 1:
            break
        string = _str.from_buffer(data, offset)
        offset += string.size
        result.append(string.value)return result
_get_gump_short_lines = _ScriptMethod(227)  # GetGumpShortLines
_get_gump_short_lines.restype = _buffer
_get_gump_short_lines.argtypes = [_ushort]  # GumpIndex
func GetGumpShortLines(GumpIndex){
    result = []
}
    data = _get_gump_short_lines(GumpIndex)
    count = _uint.from_buffer(data)
    offset = count.size
    while 42:
        if offset >= len(data) - 1:
            break
        string = _str.from_buffer(data, offset)
        offset += string.size
        result.append(string.value)return result
_get_gump_buttons = _ScriptMethod(228)  # GetGumpButtonsDescription
_get_gump_buttons.restype = _buffer
_get_gump_buttons.argtypes = [_ushort]  # GumpIndex
func GetGumpButtonsDescription(GumpIndex){
    result = []
}
    data = _get_gump_buttons(GumpIndex)
    count = _uint.from_buffer(data)
    offset = count.size
    while 42:
        if offset >= len(data) - 1:
            break
        string = _str.from_buffer(data, offset)
        offset += string.size
        result.append(string.value)return result
_get_gump_info = _ScriptMethod(229)  # GetGumpInfo
_get_gump_info.restype = _buffer  # TGumpInfo
_get_gump_info.argtypes = [_ushort]  # GumpIndex
class _Group:
    args = [_int] * 3
    container = 'groups'
    keys = 'GroupNumber', 'Page', 'ElemNum'
class _EndGroup(_Group):
    container = 'EndGroups'
class _GumpButton:
    args = [_int] * 9
    container = 'GumpButtons'
    keys = ('X', 'Y', 'ReleasedID', 'PressedID', 'Quit', 'PageID',returnValue', 'Page', 'ElemNum')
class _ButtonTileArt:
    args = [_int] * 12
    container = 'ButtonTileArts'
    keys = ('X', 'Y', 'ReleasedID', 'PressedID', 'Quit', 'PageID',returnValue', 'ArtID', 'Hue', 'ArtX', 'ArtY', 'ElemNum')
class _CheckBox:
    args = [_int] * 8
    container = 'CheckBoxes'returnValue',
            'Page', 'ElemNum')
class _ChekerTrans:
    args = [_int] * 6
    container = 'ChekerTrans'
    keys = 'X', 'Y', 'Width', 'Height', 'Page', 'ElemNum'
class _CroppedText:
    args = [_int] * 8
    container = 'CroppedText'
    keys = 'X', 'Y', 'Width', 'Height', 'Color', 'TextID', 'Page', 'ElemNum'
class _GumpPic:
    args = [_int] * 6
    container = 'GumpPics'
    keys = 'X', 'Y', 'ID', 'Hue', 'Page', 'ElemNum'
class _GumpPicTiled:
    fmt = '=7i'
    args = [_int] * 7
    container = 'GumpPicTiled'
    keys = 'X', 'Y', 'Width', 'Height', 'GumpID', 'Page', 'ElemNum'
class _Radiobutton:
    args = [_int] * 8
    container = 'RadioButtons'returnValue',
            'Page', 'ElemNum')
class _ResizePic:
    args = [_int] * 7
    container = 'ResizePics'
    keys = 'X', 'Y', 'GumpID', 'Width', 'Height', 'Page', 'ElemNum'
class _GumpText:
    args = [_int] * 6
    container = 'GumpText'
    keys = 'X', 'Y', 'Color', 'TextID', 'Page', 'ElemNum'
class _TextEntry:
    args = [_int] * 7 + [_str, _int, _int]
    container = 'TextEntries'returnValue',
            'DefaultTextID', 'RealValue', 'Page', 'ElemNum')
class _Text:
    args = [_str]
    container = 'Text'
    keys = None
class _TextEntryLimited:
    args = [_int] * 10
    container = 'TextEntriesLimited'returnValue',
            'DefaultTextID', 'Limit', 'Page', 'ElemNum')
class _TilePic:
    args = [_int] * 5
    container = 'TilePics'
    keys = 'X', 'Y', 'ID', 'Page', 'ElemNum'
class _TilePicHue:
    args = [_int] * 6
    container = 'TilePicHue'
    keys = 'X', 'Y', 'ID', 'Color', 'Page', 'ElemNum'
class _Tooltip:
    args = [_uint, _str, _int, _int]
    container = 'Tooltips'
    keys = 'ClilocID', 'Arguments', 'Page', 'ElemNum'
class _HtmlGump:
    args = [_int] * 9
    container = 'HtmlGump'
    keys = ('X', 'Y', 'Width', 'Height', 'TextID', 'Background', 'Scrollbar',
            'Page', 'ElemNum')
class _XmfHtmlGump:
    args = [_int] * 4 + [_uint] + [_int] * 4
    container = 'XmfHtmlGump'
    keys = ('X', 'Y', 'Width', 'Height', 'ClilocID', 'Background', 'Scrollbar',
            'Page', 'ElemNum')
class _XmfHTMLGumpColor:
    args = [_int] * 4 + [_uint] + [_int] * 5
    container = 'XmfHTMLGumpColor'
    keys = ('X', 'Y', 'Width', 'Height', 'ClilocID', 'Background', 'Scrollbar',
            'Hue', 'Page', 'ElemNum')
class _XmfHTMLTok:
    args = [_int] * 7 + [_uint, _str, _int, _int]
    container = 'XmfHTMLTok'
    keys = ('X', 'Y', 'Width', 'Height', 'Background', 'Scrollbar', 'Color',
            'ClilocID', 'Arguments', 'Page', 'ElemNum')
class _ItemProperty:
    args = [_uint, _int]
    container = 'ItemProperties'
    keys = 'Prop', 'ElemNum'
class _Gump:
    fmt = '<2I2hi4?'
    args = [_uint, _uint, _short, _short, _int] + [_bool] * 4
    keys = ('Serial', 'GumpID', 'X', 'Y', 'Pages', 'NoMove', 'NoResize',
            'NoDispose', 'NoClose')
*/
func GetGumpInfo(gi uint16) chan m.Gump {
	p := n.NewGetGumpInfoPacket(gi)
	return p.Out
}

/*
func GetGumpInfo(GumpIndex){
    data = _get_gump_info(GumpIndex)
}
    values = _struct.unpack_from(_Gump.fmt, data, 0)
    result = dict(zip(_Gump.keys, values))
    offset = _struct.calcsize(_Gump.fmt)
    # parse elements
    elements = (_Group, _EndGroup, _GumpButton, _ButtonTileArt, _CheckBox,
                _ChekerTrans, _CroppedText, _GumpPic, _GumpPicTiled,
                _Radiobutton, _ResizePic, _GumpText, _TextEntry, _Text,
                _TextEntryLimited, _TilePic, _TilePicHue, _Tooltip,
                _HtmlGump, _XmfHtmlGump, _XmfHTMLGumpColor, _XmfHTMLTok,
                _ItemProperty)
    for cls in elements:
        result[cls.container] = []
        count = _uint.from_buffer(data, offset)
        offset += count.size
        for i in range(count):
            values = []
            for arg in cls.args:
                element = arg.from_buffer(data, offset)
                offset += element.size
                values.append(element.value)
            if cls is _Text:
                result[cls.container].append(
                    *[values])  # there is only one element
            else:
                element = dict(zip(cls.keys, values))
                if 'ClilocID' in cls.keys and 'Arguments' in cls.keys:  # need to represent clilocs
                    text = GetClilocByID(element['ClilocID'])
                    args = element.get('Arguments', '')
                    args = args.split('@')[1:] or []
                    for arg in args:
                        if '~' in text:
                            if arg.startswith('#'):  # another cliloc
                                arg = GetClilocByID(int(arg.strip('#')))
                            s = text.index('~')
                            e = text.index('~', s + 1)
                            text = text.replace(text[s:e + 1], arg,
                                                1) or arg  # TODO: wtf?
                    element['Arguments'] = text
                result[cls.container].append(element)return result
*/
func AddGumpIgnoreByID(ID uint32) {
	n.NewVoidPacket(230, ID)
}

func AddGumpIgnoreBySerial(Serial uint32) {
	n.NewVoidPacket(231, Serial)
}

func ClearGumpsIgnore() {
	n.NewVoidPacket(232)
}

func RhandLayer() byte {
	return 0x01
}

func LhandLayer() byte {
	return 0x02
}

func ShoesLayer() byte {
	return 0x03
}

func PantsLayer() byte {
	return 0x04
}
func ShirtLayer() byte {
	return 0x05
}

func HatLayer() byte {
	return 0x06
}

func GlovesLayer() byte {
	return 0x07
}

func RingLayer() byte {
	return 0x08
}

func TalismanLayer() byte {
	return 0x09
}

func NeckLayer() byte {
	return 0x0A
}
func HairLayer() byte {
	return 0x0B
}

func WaistLayer() byte {
	return 0x0C
}

func TorsoLayer() byte {
	return 0x0D
}

func BraceLayer() byte {
	return 0x0E
}

func BeardLayer() byte {
	return 0x10
}

func TorsoHLayer() byte {
	return 0x11
}

func EarLayer() byte {
	return 0x12
}
func ArmsLayer() byte {
	return 0x13
}
func CloakLayer() byte {
	return 0x14
}

func BpackLayer() byte {
	return 0x15
}

func RobeLayer() byte {
	return 0x16
}

func EggsLayer() byte {
	return 0x17
}
func LegsLayer() byte {
	return 0x18
}

func HorseLayer() byte {
	return 0x19
}

func RstkLayer() byte {
	return 0x1A
}
func NRstkLayer() byte {
	return 0x1B
}

func SellLayer() byte {
	return 0x1C
}

func BankLayer() byte {
	return 0x1D
}

func ObjAtLayerEx(layerType byte, playerID uint32) <-chan uint32 {
	p := n.NewUint32Packet(c.SCObjAtLayerEx, layerType, playerID)
	return p.Out
}

func ObjAtLayer(LayerType byte) <-chan uint32 {
	return ObjAtLayerEx(LayerType, <-Self())
}

func GetLayer(Obj uint32) <-chan byte {
	p := n.NewBytePacket(c.SCGetLayer)
	return p.Out
}

func WearItem(layer byte, oid uint32) {
	n.NewVoidPacket(c.SCWearItem, layer, oid)
}

/*
func Disarm(){
    backpack = Backpack()
}
    tmp = []
    for layer in LhandLayer(), RhandLayer():
        item = ObjAtLayer(layer)
        if item:
            tmp.append(MoveItem(item, 1, backpack, 0, 0, 0))
    p :=
return all(tmp)
func disarm(){return Disarm()
}
func Equip(Layer, Obj){
    if Layer and DragItem(Obj, 1):
}
        p :=
return WearItem(Layer, Obj)return False
func equip(Layer, Obj){
    p :=
return Equip(Layer, Obj)
}

func Equipt(Layer, ObjType){
    item = FindType(ObjType, Backpack())
}
    if item:
        p :=


return Equip(Layer, item)return False
func equipt(Layer, ObjType){
    p :=
return Equipt(Layer, ObjType)
}
func UnEquip(Layer){
    item = ObjAtLayer(Layer)
}
    if item:
        p :=
return MoveItem(item, 1, Backpack(), 0, 0, 0)return False
*/
func GetDressSpeed() <-chan uint16 {
	p := n.NewUint16Packet(236)
	return p.Out
}

func SetDressSpeed(Value uint16) {
	n.NewVoidPacket(237, Value)
}

func GetClientVersionInt() <-chan uint32 {
	p := n.NewUint32Packet(355)
	return p.Out
}

/*
_wearable_layers = (RhandLayer(), LhandLayer(), ShoesLayer(), PantsLayer(),
                    ShirtLayer(), HatLayer(), GlovesLayer(), RingLayer(),
                    NeckLayer(), WaistLayer(), TorsoLayer(), BraceLayer(),
                    TorsoHLayer(), EarLayer(), ArmsLayer(), CloakLayer(),
*/
func UnequipItemsSetMacro() {
	n.NewVoidPacket(356)
}

/*
func Undress(){
    tmp = []
}
    client_version_int = GetClientVersionInt()
    if client_version_int < 7007400:
        delay = GetDressSpeed()
        char = Self()
        backpack = Backpack()
        for layer in _wearable_layers:
            item = ObjAtLayerEx(layer, char)
            if item:
                tmp.append(MoveItem(item, 1, backpack, 0, 0, 0))
                Wait(delay)
    else:
        UnequipItemsSetMacro()
        tmp.append(True)
    # no need to wait - all this done inside
    p :=
return all(tmp)
*/
func SetDress() {
	n.NewVoidPacket(238)
}

func EquipItemsSetMacro() {
	n.NewVoidPacket(357)
}

/*
_get_dress_set = _ScriptMethod(239)  # GetDressSet
_get_dress_set.restype = _buffer  # TLayersObjectsList
func EquipDressSet(){
    res = []
}
    client_version_int = GetClientVersionInt()
    if client_version_int < 7007400:
        delay = GetDressSpeed()
        data = _get_dress_set()
        count = _uint.from_buffer(data)
        data = data[count.size:]
        offset = 0
        for i in range(count):
            layer, item = _struct.unpack_from('<BI', data, offset)
            offset += 5
            if item:
                res.append(Equip(layer, item))
                Wait(delay)
    else:
        EquipItemsSetMacro()
        res.append(True)
    # no need to wait - all this done inside
    p :=
return all(res)
func DressSavedSet(){
    EquipDressSet()
}
func Count(ObjType){
    FindType(ObjType, Backpack())
}return FindFullQuantity()
func CountGround(ObjType){
    FindType(ObjType, Ground())
}return FindFullQuantity()
func CountEx(ObjType, Color, Container){
    FindTypeEx(ObjType, Color, Container, False)
}return FindFullQuantity()

*/ // return 0x0F7A }return 0x0F7B }return 0x0F84 }return 0x0F85 }return 0x0F86 }return 0x0F88 }return 0x0F8C }return 0x0F8D }

func BP() uint16 { return 0x0F7A }
func BM() uint16 { return 0x0F7B }
func GA() uint16 { return 0x0F84 }
func GS() uint16 { return 0x0F85 }
func MR() uint16 { return 0x0F86 }
func NS() uint16 { return 0x0F88 }
func SA() uint16 { return 0x0F8C }
func SS() uint16 { return 0x0F8D }

func BPCount() <-chan uint32 {
	<-FindTypeEx(BP(), 0, <-Backpack(), true)
	return FindFullQuantity()
}

func BMCount() <-chan uint32 {
	<-FindTypeEx(BM(), 0, <-Backpack(), true)
	return FindFullQuantity()
}
func GACount() <-chan uint32 {
	<-FindTypeEx(GA(), 0, <-Backpack(), true)
	return FindFullQuantity()
}
func GSCount() <-chan uint32 {
	<-FindTypeEx(GS(), 0, <-Backpack(), true)
	return FindFullQuantity()
}
func MRCount() <-chan uint32 {
	<-FindTypeEx(MR(), 0, <-Backpack(), true)
	return FindFullQuantity()
}
func NSCount() <-chan uint32 {
	<-FindTypeEx(NS(), 0, <-Backpack(), true)
	return FindFullQuantity()
}
func SACount() <-chan uint32 {
	<-FindTypeEx(SA(), 0, <-Backpack(), true)
	return FindFullQuantity()
}
func SSCount() <-chan uint32 {
	<-FindTypeEx(SS(), 0, <-Backpack(), true)
	return FindFullQuantity()
}

func AutoBuy(itemType, itemColor, quantity uint16) {
	n.NewVoidPacket(c.SCAutoBuy, itemType, itemColor, quantity)
}

func GetShopList() <-chan []string {
	p := n.NewGetShopListPacket()
	return p.Out
}

func ClearShopList() {
	n.NewVoidPacket(242)
}

func AutoBuyEx(itemType, itemColor, quantity uint16, price uint32, itemName string) {
	n.NewVoidPacket(c.SCAutoBuyEx, itemType, itemColor, itemName, quantity, price, itemName)
}

func GetAutoBuyDelay() <-chan uint16 {
	p := n.NewUint16Packet(244)
	return p.Out
}

func SetAutoBuyDelay(Value uint16) {
	n.NewVoidPacket(245, Value)
}

func GetAutoSellDelay() <-chan uint16 {
	p := n.NewUint16Packet(246)
	return p.Out
}

func SetAutoSellDelay(Value uint16) {
	n.NewVoidPacket(247, Value)
}

func AutoSell(itemType, itemColor, quantity uint16) {
	n.NewVoidPacket(c.SCAutoSell, itemType, itemColor, quantity)
}
func RequestStats(ObjID uint32) {
	n.NewVoidPacket(249, ObjID)
}

func HelpRequest() {
	n.NewVoidPacket(250)
}

func QuestRequest() {
	n.NewVoidPacket(251)
}

func RenameMobile(mobId uint32, newName string) {
	n.NewVoidPacket(c.SCRenameMobile, mobId, newName)
}

func MobileCanBeRenamed(Mob_ID uint32) <-chan bool {
	p := n.NewBoolPacket(253, Mob_ID)
	return p.Out
}

func SetStatState(statNum, statState byte) {
	n.NewVoidPacket(254, statNum, statState)
}

func GetStaticArtBitmap(id uint32, hue uint16) <-chan []byte {
	p := n.NewByteArrayPacket(c.SCGetStaticArtBitmap, id, hue)
	return p.Out
}

func PrintScriptMethodsList(fileName string, sortedList bool) {
	n.NewVoidPacket(256, fileName, sortedList)
}

func Alarm() {
	n.NewVoidPacket(257)
}

func UOSay(text string) {
	n.NewVoidPacket(308, text)
}

func UOSayColor(text string, color uint16) {
	n.NewVoidPacket(309, text, color)
}

func SetGlobal(varRegion byte, varName, varValue string) {
	n.NewVoidPacket(c.SCSetGlobal, varRegion, varName, varValue)
}

func GetGlobal(varRegion byte, varName string) <-chan string {
	p := n.NewStringPacket(c.SCGetGlobal, varRegion, varName)
	return p.Out
}

/*
_reg_stealth = 0, '0', 'reg_stealth', 'stealth'
_reg_char = 1, '1', 'reg_char', 'char'
_set_global = _ScriptMethod(310)  # SetGlobal
_set_global.argtypes = [_ubyte,  # GlobalRegion
                        _str,  # VarName
                        _str]  # VarValue
func SetGlobal(GlobalRegion, VarName, VarValue){
    if isinstance(GlobalRegion, str):
}
        GlobalRegion = GlobalRegion.lower()
    for region in _reg_stealth, _reg_char:
        if GlobalRegion in region:
            _set_global(region[0], VarName, VarValue)
            break
    else:
        raise ValueError('GlobalRegion must be "stealth" or "char".')
_get_global = _ScriptMethod(311)
_get_global.restype = _str
_get_global.argtypes = [_ubyte,  # GlobalRegion
                        _str]  # VarName
func GetGlobal(GlobalRegion, VarName){
    if isinstance(GlobalRegion, str):
}
        GlobalRegion = GlobalRegion.lower()
    for region in _reg_stealth, _reg_char:
        if GlobalRegion in region:
            p :=
return _get_global(region[0], VarName)
    else:
        raise ValueError('GlobalRegion must be "stealth" or "char".')
*/
func ConsoleEntryReply(text string) {
	n.NewVoidPacket(312, text)
}
func ConsoleEntryUnicodeReply(text string) {
	n.NewVoidPacket(313, text)
}

func GameServerIPString() <-chan string {
	p := n.NewStringPacket(341)
	return p.Out
}

/*
_easyuo_sub_key = 'Software\\EasyUO'
func SetEasyUO(num, Regvalue){
    if b'' == '':  # py2
}
        import _winreg as winreg
    else:
        import winreg
    key = winreg.HKEY_CURRENT_USER
    access = winreg.KEY_WRITE
    with winreg.OpenKey(key, _easyuo_sub_key, 0, access) as easyuo_key:
        winreg.SetValueEx(easyuo_key, '*' + str(num), 0, winreg.REG_SZ,
                          Regvalue)
func GetEasyUO(num){
    if b'' == '':  # py2
}
        import _winreg as winreg
    else:
        import winreg
    key = winreg.HKEY_CURRENT_USER
    access = winreg.KEY_READ
    with winreg.OpenKey(key, _easyuo_sub_key, 0, access) as easyuo_key:
        type_, data = winreg.QueryValueEx(easyuo_key, '*' + str(num))return data
func EUO2StealthType(EUO){
    # TODO: 2 and 3 compatible code: int(codecs.encode(b'A', 'hex'), 16)
}
    res = 0
    multi = 1
    for char in EUO:
        if b'' == '':  # py2
            tmp = int(char.encode('hex'), 16)
        else:
            tmp = int.from_bytes(char.encode(), 'little')
        res += multi * (tmp - 65)
        multi *= 26
    res = (res - 7) ^ 0x0045return 0 if res > 0xFFFF else res
func EUO2StealthID(EUO){
    # TODO: 2 and 3 compatible code: int(codecs.encode(b'A', 'hex'), 16)
}
    res = 0
    multi = 1
    for char in EUO:
        if b'' == '':  # py2
            tmp = int(char.encode('hex'), 16)
        else:
            tmp = int.from_bytes(char.encode(), 'little')
        res += multi * (tmp - 65)
        multi *= 26return (res - 7) ^ 0x0045
*/

func InviteToParty(ID uint32) {
	n.NewVoidPacket(262, ID)
}

func RemoveFromParty(ID uint32) {
	n.NewVoidPacket(263, ID)
}

func PartyMessageTo(oid uint32, msg string) {
	n.NewVoidPacket(c.SCPartyMessageTo, msg)
}

func PartySay(msg string) {
	n.NewVoidPacket(265, msg)
}

func PartyCanLootMe(value bool) {
	n.NewVoidPacket(c.SCPartyCanLootMe, value)
}

func PartyAcceptInvite() {
	n.NewVoidPacket(267)
}

func PartyDeclineInvite() {
	n.NewVoidPacket(268)
}

func PartyLeave() {
	n.NewVoidPacket(269)
}

func InParty() <-chan bool {
	p := n.NewBoolPacket(271)
	return p.Out
}

func PartyMembersList() <-chan []uint32 {
	p := n.NewUint32ArrayPacket(c.SCPartyMembersList)
	return p.Out
}

func ICQConnected() <-chan bool {
	p := n.NewBoolPacket(272)
	return p.Out
}

func ICQConnect(UIN uint32, password string) {
	n.NewVoidPacket(273, UIN, password)
}

func ICQDisconnect() {
	n.NewVoidPacket(274)
}

/*
_icq_set_status = _ScriptMethod(275)  # ICQ_SetStatus
_icq_set_status.argtypes = [_ubyte]  # Num
func ICQSetStatus(Num){
    _icq_set_status(Num)
}
_icq_set_x_status = _ScriptMethod(276)  # ICQ_SetXStatus
_icq_set_x_status.argtypes = [_ubyte]  # Num
func ICQSetXStatus(Num){
    _icq_set_x_status(Num)
}
_icq_send_message = _ScriptMethod(277)  # ICQ_SendText
_icq_send_message.argtypes = [_uint,  # DestinationUIN
                              _str]  # Text
func ICQSendText(DestinationUIN, Text){
    _icq_send_message(DestinationUIN, Text)
}
_messengers = {0: 1,  # default - telegram
               1: 1, 'Telegram': 1, 'telegram': 1,
               2: 2, 'Viber': 2, 'viber': 2,
               3: 3, 'Discord': 3, 'discord': 3}_messenger_get_connected = _ScriptMethod(501)  # Messenger_GetConnected
_messenger_get_connected.restype = _bool
_messenger_get_connected.argtypes = [_ubyte]  # MesID
func MessengerGetConnected(MesID){
    if MesID not in _messengers.keys():
}
        error = 'MessengerGetConnected: MesID must be "Telegram", "Viber" or "Discord"'
        raise ValueError(error)
    p :=
return _messenger_get_connected(_messengers[MesID])
_messenger_set_connected = _ScriptMethod(502)  # Messenger_SetConnected
_messenger_set_connected.argtypes = [_ubyte,  # MesID
                                     _bool]  # Value
func MessengerSetConnected(MesID, Value){
    if MesID not in _messengers.keys():
}
        error = 'MessengerGetConnected: MesID must be "Telegram", "Viber" or "Discord"'
        raise ValueError(error)
    _messenger_set_connected(_messengers[MesID], Value)
_messenger_get_token = _ScriptMethod(503)  # Messenger_GetToken
_messenger_get_token.restype = _str
_messenger_get_token.argtypes = [_ubyte]  # MesID
func MessengerGetToken(MesID){
    if MesID not in _messengers.keys():
}
        error = 'MessengerGetConnected: MesID must be "Telegram", "Viber" or "Discord"'
        raise ValueError(error)
    p :=
return _messenger_get_token(_messengers[MesID])
_messenger_set_token = _ScriptMethod(504)  # Messenger_SetToken
_messenger_set_token.argtypes = [_ubyte,  # MesID
                                 _str]  # Value
func MessengerSetToken(MesID, Value){
    if MesID not in _messengers.keys():
}
        error = 'MessengerGetConnected: MesID must be "Telegram", "Viber" or "Discord"'
        raise ValueError(error)
    _messenger_set_token(_messengers[MesID], Value)
_messenger_get_name = _ScriptMethod(505)  # Messenger_GetName
_messenger_get_name.restype = _str
_messenger_get_name.argtypes = [_ubyte]  # MesID
func MessengerGetName(MesID){
    if MesID not in _messengers.keys():
}
        error = 'MessengerGetConnected: MesID must be "Telegram", "Viber" or "Discord"'
        raise ValueError(error)
    p :=
return _messenger_get_name(_messengers[MesID])
_messenger_send_message = _ScriptMethod(506)  # Messenger_SendMessage
_messenger_send_message.argtypes = [_ubyte,  # MesID
                                    _str,  # Msg
                                    _str]  # UserID
func MessengerSendMessage(MesID, Msg, UserID){
    if MesID not in _messengers.keys():
}
        error = 'MessengerGetConnected: MesID must be "Telegram", "Viber" or "Discord"'
        raise ValueError(error)
    _messenger_send_message(_messengers[MesID], Msg, UserID)
_tile_groups = {0: 0, 'tfLand': 0, 'tfland': 0, 'Land': 0, 'land': 0,
                1: 1, 'tfStatic': 1, 'tfstatic': 1, 'Static': 1, 'static': 1}

_get_tile_flags = _ScriptMethod(278)  # GetTileFlags
_get_tile_flags.restype = _uint
_get_tile_flags.argtypes = [_ubyte,  # TileGroup
                            _ushort]  # Tile
*/
func GetTileFlags(tileGroup byte, tile uint32) <-chan uint32 {
	p := n.NewUint32Packet(c.SCGetTileFlags, tileGroup, tile)
	return p.Out
}

/*
_uint_to_flags = _ScriptMethod(350)  # ConvertIntegerToFlags
_uint_to_flags.restype = _buffer
_uint_to_flags.argtypes = [_ubyte,  # Group
                           _uint]  # Flags

func ConvertIntegerToFlags(Group, Flags){
    if Group not in _tile_groups.keys():
}
        raise ValueError('GetTileFlags: Group must be "Land" or "Static"')
    result = []
    data = _uint_to_flags(_tile_groups[Group], Flags)
    count = _uint.from_buffer(data)
    offset = count.size
    while 42:
        if offset >= len(data) - 1:
            break
        string = _str.from_buffer(data, offset)
        offset += string.size
        result.append(string.value)return result
_get_land_tile_data = _ScriptMethod(280)  # GetLandTileData
_get_land_tile_data.restype = _buffer  # TLandTileData
_get_land_tile_data.argtypes = [_ushort]  # Tile
func GetLandTileData(Tile){
    result = {}
}
    data = _get_land_tile_data(Tile)
    if data:
        result['Flags'] = ConvertIntegerToFlags(0, _uint.from_buffer(data))
        result['Flags2'] = ConvertIntegerToFlags(0, _uint.from_buffer(data[4:]))
        result['TextureID'] = _ushort.from_buffer(data[8:])
        length = _uint.from_buffer(data[10:])
        result['Name'] = data[14:14 + length].rstrip(b'\x00')
        if b'' != '':  # py3
            result['Name'] = result['Name'].decode()return result
_get_static_tile_data = _ScriptMethod(281)  # GetStaticTileData
_get_static_tile_data.restype = _buffer  # TStaticTileDataNew
_get_static_tile_data.argtypes = [_ushort]  # Tile
func GetStaticTileData(Tile){
    result = {}
}
    data = _get_static_tile_data(Tile)
    if data:
        result['Flags'] = ConvertIntegerToFlags(1, _ulong.from_buffer(data))
        result['Weight'] = _ushort.from_buffer(data, 8)
        result['AnimID'] = _ushort.from_buffer(data, 10)
        result['Height'] = _int.from_buffer(data, 12)
        result['RadarColorRGBA'] = _struct.unpack_from('<4B', data, 16)
        length = _uint.from_buffer(data, 20)
        result['Name'] = data[24: 24 + length].rstrip(b'\x00')
        if b'' != '':  # py3
            result['Name'] = result['Name'].decode()return result
_get_cell = _ScriptMethod(13)  # GetCell
_get_cell.restype = _buffer  # TMapCell
_get_cell.argtypes = [_ushort,  # X
                      _ushort,  # Y
                      _ubyte]  # WorldNum
*/

func GetCell(x, y uint16, worldNum byte) <-chan m.MapCell {
	p := n.NewGetMapCellPacket(x, y, worldNum)
	return p.Out
}

func GetLayerCount(x, y uint16, worldNum byte) <-chan byte {
	p := n.NewBytePacket(c.SCGetLayerCount, x, y, worldNum)
	return p.Out
}

func ReadStaticsXY(x, y uint16, wnum byte) <-chan []m.StaticsXY {
	p := n.NewReadStaticsXYPacket(x, y, wnum)
	return p.Out
}

func GetSurfaceZ(x, y uint16, worldNum byte) <-chan byte {
	p := n.NewBytePacket(c.SCGetSurfaceZ, x, y, worldNum)
	return p.Out
}

func IsWorldCellPassable(currX, currY uint16, currZ int8, destX, destY uint16, worldNum byte) <-chan m.WorldCellPassable {
	p := n.NewIsWorldCellPassablePacket(currX, currY, currZ, destX, destY, worldNum)
	return p.Out
}

/*
_get_statics_array = _ScriptMethod(286)  # GetStaticTilesArray
_get_statics_array.restype = _buffer  # Array of TFoundTile
_get_statics_array.argtypes = [_ushort,  # Xmin
                               _ushort,  # Ymin
                               _ushort,  # Xmax
                               _ushort,  # Ymax
                               _ubyte,  # WorldNum
                               _uint,  # Len
                               _buffer]  # TileTypes: Array of Word
*/
func GetStaticTilesArray(xmin, ymin, xmax, ymax uint16, worldNum byte, tileTypes []uint16) <-chan []m.FoundTile {
	p := n.NewGetStaticTilesArrayPacket(xmin, ymin, xmax, ymax, worldNum, tileTypes)
	return p.Out
}

/*
_get_lands_array = _ScriptMethod(287)  # GetLandTilesArray
_get_lands_array.restype = _buffer  # Array of TFoundTile
_get_lands_array.argtypes = [_ushort,  # Xmin
                             _ushort,  # Ymin
                             _ushort,  # Xmax
                             _ushort,  # Ymax
                             _ubyte,  # WorldNum
                             _uint,  # Len
                             _buffer]  # TileTypes: Array of Word
func GetLandTilesArray(Xmin, Ymin, Xmax, Ymax, WorldNum, TileTypes){
    if not _iterable(TileTypes):
}
        TileTypes = [TileTypes]
    result = []
    data = _get_lands_array(
        Xmin, Ymin, Xmax, Ymax, WorldNum, len(TileTypes),
        _struct.pack('<' + 'H' * len(TileTypes), *TileTypes)
    )
    count = _uint.from_buffer(data)
    fmt = '<3Hb'
    size = _struct.calcsize(fmt)
    for i in range(count):
        result.append(_struct.unpack_from(fmt, data, count.size + i * size))return result
*/
func ClientPrint(text string) {
	n.NewVoidPacket(289, text)
}

func ClientPrintEx(senderID uint32, color, font uint16, text string) {
	n.NewVoidPacket(c.SCClientPrintEx, senderID, color, font, text)
}

/*
_wnd = {0: 0, '0': 0, 'wtpaperdoll': 0, 'paperdoll': 0,
        1: 1, '1': 1, 'wtstatus': 1, 'status': 1,
        2: 2, '2': 2, 'wtcharprofile': 2, 'charprofile': 2, 'profile': 2,
        3: 3, '3': 3, 'wtcontainer': 3, 'container': 3}_close_client_ui_window = _ScriptMethod(291)  # CloseClientUIWindow
_close_client_ui_window.argtypes = [_ubyte,  # UIWindowType
                                    _uint]  # ID
func CloseClientUIWindow(UIWindowType, ID){
    if isinstance(UIWindowType, str):
}
        UIWindowType = UIWindowType.lower()
    if UIWindowType not in _wnd.keys():
        raise ValueError('CloseClientUIWindow: UIWindowType must be '
                         '"Paperdoll", "Status", "CharProfile" or "Container"')
    _close_client_ui_window(_wnd[UIWindowType], ID)
*/
func ClientRequestObjectTarget() {
	n.NewVoidPacket(c.SCClientRequestObjectTarget)
}

func ClientRequestTileTarget() {
	n.NewVoidPacket(c.SCClientRequestTileTarget)
}

func ClientTargetResponsePresent() <-chan bool {
	p := n.NewBoolPacket(294)
	return p.Out
}

func ClientTargetResponse() <-chan m.TargetInfo {
	p := n.NewClientTargetInfoPacket()
	return p.Out
}

func WaitForClientTargetResponse(maxWaitTime time.Duration) <-chan bool {
	r := make(chan bool)
	t := time.Now().Add(maxWaitTime)
	go func() {
		for {
			time.Sleep(time.Millisecond * 1000)
			if <-ClientTargetResponsePresent() {
				r <- true
				break
			}

			if time.Now().After(t) {
				r <- false
				break
			}
		}
	}()
	return r
}

func CheckLag(timeoutMS uint32) <-chan bool {
	p := n.NewBoolPacket(299, timeoutMS)
	return p.Out
}

func GetQuestArrow() <-chan m.Point2D {
	p := n.NewPoint2DPacket(c.SCGetQuestArrow)
	return p.Out
}

func GetSilentMode() <-chan bool {
	p := n.NewBoolPacket(302)
	return p.Out
}

func ClearInfoWindow() {
	n.NewVoidPacket(348)
}

func SetSilentMode(value bool) {
	n.NewVoidPacket(c.SCSetSilentMode, value)
}

func FillNewWindow(s string) {
	n.NewVoidPacket(303, s)
}

func StealthPath() <-chan string {
	p := n.NewStringPacket(305)
	return p.Out
}

func CurrentScriptPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

func GetStealthProfilePath() <-chan string {
	p := n.NewStringPacket(306)
	return p.Out
}

func GetShardPath() <-chan string {
	p := n.NewStringPacket(307)
	return p.Out
}

func Step(direction byte, running bool) <-chan byte {
	p := n.NewBytePacket(c.SCStep, direction, running)
	return p.Out
}

func StepQ(direction byte, running bool) <-chan int32 {
	p := n.NewIntPacket(c.SCStepQ, direction, running)
	return p.Out
}

func MoveXYZ(xdst, ydst uint16, zdst int8, optimized bool, accuracyXY, accurancyZ int32, running bool) <-chan bool {
	p := n.NewBoolPacket(c.SCMoveXYZ, xdst, ydst, zdst, accuracyXY, accurancyZ, running)
	return p.Out
}

func MoveXY(xdst, ydst uint16, optimized bool, accuracy int32, running bool) <-chan bool {
	p := n.NewBoolPacket(c.SCMoveXY, xdst, ydst, optimized, accuracy, running)
	return p.Out
}

func SetBadLocation(x, y uint16) {
	n.NewVoidPacket(c.SCSetBadLocation, x, y)
}

func SetGoodLocation(x, y uint16) {
	n.NewVoidPacket(c.SCSetGoodLocation, x, y)
}

func ClearBadLocationList() {
	n.NewVoidPacket(330)
}

func SetBadObject(otype, color uint16, radius byte) {
	n.NewVoidPacket(c.SCSetBadObject, otype, color, radius)
}

func ClearBadObjectList() {
	n.NewVoidPacket(332)
}

/*
_los_check_type = {1: 1, '1': 1, 'lossphere': 1, 'sphere': 1,
                   2: 2, '2': 2, 'lossphereadv': 2, 'sphereadv': 2,
                   3: 3, '3': 3, 'lospol': 3, 'pol': 3,
                   4: 4, '4': 4, 'losrunuo': 4, 'runuo': 4, 'servuo': 4}_los_check_options = {0: 0, '0': 0, None: 0,
                      0x100: 0x100,
                      'losspherecheckcorners': 0x100,
                      'spherecheckcorners': 0x100,
                      0x200: 0x200,
                      'lospolusenoshoot': 0x200,
                      'polusenoshoot': 0x200,
                      0x400: 0x400,
                      'lospollosthroughwindow': 0x400,
                      'pollosthroughwindow': 0x400}_check_los = _ScriptMethod(333)  # CheckLOS
_check_los.restype = _bool
_check_los.argtypes = [_ushort,  # xf
                       _ushort,  # yf
                       _byte,  # zf
                       _ushort,  # xt
                       _ushort,  # yt
                       _byte,  # zt
                       _ubyte,  # WorldNum
                       _ubyte,  # LOSCheckType
                       _uint]  # LOSOptions
func CheckLOS(xf, yf, zf, xt, yt, zt, WorldNum, LOSCheckType, LOSOptions=None){
    if not _iterable(LOSOptions) or isinstance(LOSOptions, str):
}
        LOSOptions = [LOSOptions]
    if isinstance(LOSCheckType, str):
        LOSCheckType = LOSCheckType.lower()
    if LOSCheckType not in _los_check_type.keys():
        raise ValueError('CheckLOS: LOSCheckType must be "Sphere", "SphereAdv"'
                         ', "Pol" or "RunUO".')
    options = 0
    for option in LOSOptions:
        if isinstance(option, str):
            option = option.lower()
        if option not in _los_check_options.keys():
            raise ValueError('CheckLOS: LOSOptions must be set of '
                             '"SphereCheckCorners", "PolUseNoShoot", '
                             '"PolLosThroughWindow" or None.')
        options |= _los_check_options[option]
    p :=
return _check_los(xf, yf, zf, xt, yt, zt, WorldNum, LOSCheckType, options)
_get_path_array = _ScriptMethod(334)  # GetPathArray
_get_path_array.restype = _buffer  # Array of TMyPoint
_get_path_array.argtypes = [_ushort,  # DestX
                            _ushort,  # DestY
                            _bool,  # Optimized
                            _int]  # Accuracy
func GetPathArray(DestX, DestY, Optimized, Accuracy){
    result = []
}
    data = _get_path_array(DestX, DestY, Optimized, Accuracy)
    count = _uint.from_buffer(data)
    fmt = '<2Hb'
    size = _struct.calcsize(fmt)
    for i in range(count):
        result.append(_struct.unpack_from(fmt, data, count.size + i * size))return result
_get_path_array_3d = _ScriptMethod(335)  # GetPathArray3D
_get_path_array_3d.restype = _buffer  # Array of TMyPoint
_get_path_array_3d.argtypes = [_ushort,  # StartX
                               _ushort,  # StartY
                               _byte,  # StartZ
                               _ushort,  # FinishX
                               _ushort,  # FinishY
                               _byte,  # FinishZ
                               _ubyte,  # WorldNum
                               _int,  # AccuracyXY
                               _int,  # AccuracyZ
                               _bool]  # Run
def GetPathArray3D(StartX, StartY, StartZ, FinishX, FinishY, FinishZ, WorldNum,
                   AccuracyXY, AccuracyZ, Run):
    result = []
    data = _get_path_array_3d(StartX, StartY, StartZ, FinishX, FinishY,
                              FinishZ, WorldNum, AccuracyXY, AccuracyZ, Run)
    count = _uint.from_buffer(data)
    fmt = '<2Hb'
    size = _struct.calcsize(fmt)
    for i in range(count):
        result.append(_struct.unpack_from(fmt, data, count.size + i * size))return result
func Dist(x1, y1, x2, y2){
    dx = abs(x2 - x1)
}
    dy = abs(y2 - y1)return dx if dx > dy else dy
func CalcCoord(x, y, Dir){
    if Dir > 7:
}return x, y
    dirs = {0: (0, -1),
            1: (1, -1),
            2: (1, 0),
            3: (1, 1),
            4: (0, 1),
            5: (-1, 1),
            6: (-1, 0),
            7: (-1, -1)}
    dx, dy = dirs[Dir]return x + dx, y + dy
func CalcDir(Xfrom, Yfrom, Xto, Yto){
    dx = abs(Xto - Xfrom)
}
    dy = abs(Yto - Yfrom)
    if dx == dy == 0:return 100
    elif (dx / (dy + 0.1)) >= 2:return 6 if Xfrom > Xto else 2
    elif (dy / (dx + 0.1)) >= 2:return 0 if Yfrom > Yto else 4
    elif Xfrom > Xto:return 7 if Yfrom > Yto else 5
    elif Xfrom < Xto:return 1 if Yfrom > Yto else 3
*/
func SetRunUnmountTimer(Value uint16) {
	n.NewVoidPacket(316, Value)
}

func SetWalkMountTimer(Value uint16) {
	n.NewVoidPacket(317, Value)
}

func SetRunMountTimer(Value uint16) {
	n.NewVoidPacket(318, Value)
}

func SetWalkUnmountTimer(Value uint16) {
	n.NewVoidPacket(319, Value)
}

func GetRunMountTimer() <-chan uint16 {
	p := n.NewUint16Packet(320)
	return p.Out
}

func GetWalkMountTimer() <-chan uint16 {
	p := n.NewUint16Packet(321)
	return p.Out
}

func GetRunUnmountTimer() <-chan uint16 {
	p := n.NewUint16Packet(322)
	return p.Out
}

func GetWalkUnmountTimer() <-chan uint16 {
	p := n.NewUint16Packet(323)
	return p.Out
}

func GetLastStepQUsedDoor() <-chan uint32 {
	p := n.NewUint32Packet(344)
	return p.Out
}

func StopMover() {
	n.NewVoidPacket(353)
}

func MoverStop() {
	StopMover()
}

func SetARExtParams(shardName, charName string, useAtEveryConnect bool) {
	n.NewVoidPacket(c.SCSetARExtParams, shardName, charName, useAtEveryConnect)
}

func UseItemOnMobile(itemSerial, targetSerial uint32) {
	n.NewVoidPacket(c.SCUseItemOnMobile, itemSerial, targetSerial)
}

func BandageSelf() {
	n.NewVoidPacket(360)
}

func GlobalChatJoinChannel(chName string) {
	n.NewVoidPacket(361, chName)
}

func GlobalChatLeaveChannel() {
	n.NewVoidPacket(362)
}

func GlobalChatSendMsg(msgText string) {
	n.NewVoidPacket(363, msgText)
}

func GlobalChatActiveChannel() <-chan string {
	p := n.NewStringPacket(364)
	return p.Out
}

/*
global_chat_channel_list = _ScriptMethod(365)  # SCGlobalChatChannelsList
global_chat_channel_list.restype = _buffer
func GlobalChatChannelsList(){
    result = []
}
    data = global_chat_channel_list()
    count = _uint.from_buffer(data)
    offset = count.size
    for i in range(count):
        string = _str.from_buffer(data, offset)
        offset += string.size
        result.append(string.value)return result
_set_open_doors = _ScriptMethod(400)  # SetMoveOpenDoor
_set_open_doors.argtypes = [_bool]  # Value
*/
func SetMoveOpenDoor(v bool) {
	n.NewVoidPacket(c.SCSetMoveOpenDoor, v)
}

func GetMoveOpenDoor() <-chan bool {
	p := n.NewBoolPacket(401)
	return p.Out
}

func SetMoveThroughNPC(Value uint16) {
	n.NewVoidPacket(402, Value)
}

func GetMoveThroughNPC() <-chan uint16 {
	p := n.NewUint16Packet(403)
	return p.Out
}

func SetMoveThroughCorner(value bool) {
	n.NewVoidPacket(c.SCGetMoveThroughCorner, value)
}

func GetMoveThroughCorner() <-chan bool {
	p := n.NewBoolPacket(405)
	return p.Out
}

func SetMoveHeuristicMult(value int32) {
	n.NewVoidPacket(c.SCSetMoveHeuristicMult, value)
}

func GetMoveHeuristicMult() <-chan uint32 {
	p := n.NewUint32Packet(407)
	return p.Out
}

func SetMoveCheckStamina(Value uint16) {
	n.NewVoidPacket(408, Value)
}

func GetMoveCheckStamina() <-chan uint16 {
	p := n.NewUint16Packet(409)
	return p.Out
}

func SetMoveTurnCost(value uint32) {
	n.NewVoidPacket(c.SCSetMoveTurnCost, value)
}

func GetMoveTurnCost() <-chan uint32 {
	p := n.NewUint32Packet(411)
	return p.Out
}

func SetMoveBetweenTwoCorners(value bool) {
	n.NewVoidPacket(c.SCSetMoveBetweenTwoCorners, value)
}

func GetMoveBetweenTwoCorners() <-chan bool {
	p := n.NewBoolPacket(413)
	return p.Out
}

func GetMultis() <-chan []m.Multi {
	p := n.NewGetMultisPacket()
	return p.Out
}

/*
func GetMenuItemsEx(MenuCaption){
    """
}
    GetMenuItemsEx(MenuCaption: str) => Array of MenuItems    MenuItems:
        model: int (item type i guess)
        color: int
        text: str    Example:
        menu_items = GetMenuItemsEx('Inscription items')
        print(menu_items[0].text)
        >> 1 Blank scroll
    """    class MenuItems:
        model = None
        color = None
        text = None        func __str__(self){
            template = 'Model: {0}, Color: {1}, Text: {2}'
}
            p :=
return '{ ' + template.format(hex(self.model), hex(self.color),
                                          self.text) + ' }'        func __repr__(self){return self.__str__()
}    data = _get_menu_items_ex(MenuCaption)
    result = []
    count = _struct.unpack_from('<I', data, 0)
    offset = count.size
    while offset < len(data):
        model, color = _struct.unpack_from('<HH', data, offset)
        offset += 4
        text = _str.from_buffer(data, offset)
        offset += text.size        item = MenuItems()
        item.model = model
        item.color = color
        item.text = text.valuereturn result
*/
func CloseClientGump(iD uint32) {
	n.NewVoidPacket(342, iD)
}

func GetNextStepZ(currX, currY, destX, destY uint16, worldNum byte, currZ int8) <-chan int8 {
	p := n.NewInt8Packet(c.SCGetNextStepZ, currX, currY, destX, destY, worldNum, currZ)
	return p.Out
}

func ClientHide(ID uint32) <-chan bool {
	p := n.NewBoolPacket(368, ID)
	return p.Out
}

func GetSkillLockState(skillName string) <-chan int8 {
	p := n.NewInt8Packet(c.SCGetSkillLockState, skillName)
	return p.Out
}

func GetStatLockState(skillName string) <-chan byte {
	p := n.NewBytePacket(c.SCGetStatLockState)
	return p.Out
}

func EquipLastWeapon() {
	n.NewVoidPacket(c.SCEquipLastWeapon)
}

func BookGetPageText(page uint16) <-chan string {
	p := n.NewStringPacket(c.SCBookGetPageText, page)
	return p.Out
}

func BookSetText(text string) {
	n.NewVoidPacket(c.SCBookSetText, text)
}

func BookSetPageText(page uint16, text string) {
	n.NewVoidPacket(c.SCBookSetText, page, text)
}

func BookClearText() {
	n.NewVoidPacket(c.SCBookClearText)
}

func BookSetHeader(title, author string) {
	n.NewVoidPacket(c.SCBookSetHeader, title, author)
}

/*
# Character creation_create_char = _ScriptMethod(371)
_create_char.argtypes =[
    _str,  # ProfileName
    _str,  # ShardName
    _str,  # CharName
    _bool, # Gender
    _byte, # Race
    _byte, # Str
    _byte, # Dex
    _byte, # Int
    _str,  # Skill1
    _str,  # Skill2
    _str,  # Skill3
    _str,  # Skill4
    _int,  # SkillValue1
    _int,  # SkillValue2
    _int,  # SkillValue3
    _int,  # SkillValue4
    _byte, # Start City
    _uint, # Free Slot
]
def CreateChar(
        ProfileName,  # ProfileName
        ShardName,  # ShardName
        CharName,  # CharName
        Gender,  # Gender
        Race,  # Race
        Strn,  # Str
        Dex,  # Dex
        Int,  # Int
        Skill1,  # Skill1
        Skill2,  # Skill2
        Skill3,  # Skill3
        Skill4,  # Skill4
        SkillValue1,  # SkillValue1
        SkillValue2,  # SkillValue2
        SkillValue3,  # SkillValue3
        SkillValue4,  # SkillValue4
        City,  # Start City
        Slot,  # Free Slot
):
    _create_char(
        ProfileName,    # ProfileName
        ShardName,      # ShardName
        CharName,       # CharName
        Gender,         # Gender
        Race,           # Race
        Strn,           # Str
        Dex,            # Dex
        Int,            # Int
        Skill1,         # Skill1
        Skill2,         # Skill2
        Skill3,         # Skill3
        Skill4,         # Skill4
        SkillValue1,    # SkillValue1
        SkillValue2,    # SkillValue2
        SkillValue3,    # SkillValue3
        SkillValue4,    # SkillValue4
        City,           # Start City
        Slot,           # Free Slot
    )
*/
//# Script control functions
func GetScriptCount() <-chan uint16 {
	p := n.NewUint16Packet(c.SCGetScriptsCount)
	return p.Out
}

func GetScriptPath(scriptIndex uint16) <-chan string {
	p := n.NewStringPacket(c.SCGetScriptPath, scriptIndex)
	return p.Out
}

func GetScriptName(ScriptIndex uint16) <-chan string {
	p := n.NewStringPacket(c.SCGetScriptName, ScriptIndex)
	return p.Out
}

func GetScriptState(scriptIndex uint16) <-chan int8 {
	p := n.NewInt8Packet(c.SCGetScriptState, scriptIndex)
	return p.Out

}

func StartScript(scriptPath string) <-chan uint16 {
	p := n.NewUint16Packet(c.SCStartScript, scriptPath)
	return p.Out
}

func StopScript(scriptIndex uint16) {
	n.NewVoidPacket(456)
}

func PauseResumeScript(scriptIndex uint16) {
	n.NewVoidPacket(456)
}

func StopAllScripts() {
	n.NewVoidPacket(457)
}
