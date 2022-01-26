package gostealthclient

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Connected() <-chan bool {
	p := NewBoolPacket(SCGetConnectedStatus)
	p.send(senderFunc)
	return p.out
}
func AddToSystemJournal(text string) {
	p := NewVoidPacket(SCAddToSystemJournal, text)
	p.send(senderFunc)
}
func GetStealthInfo() <-chan stealthClientInfo {
	p := NewStealthClientInfoPacket()
	go p.transform()
	p.send(senderFunc)
	return p.out
}
func Connect() {
	p := NewVoidPacket(SCConnect)
	p.send(senderFunc)
}
func Disconnect() {
	p := NewVoidPacket(SCDisconnect)
	p.send(senderFunc)
}
func SetPauseScriptOnDisconnectStatus(value bool) {
	p := NewVoidPacket(SCSetPauseScriptOnDisconnectStatus, value)
	p.send(senderFunc)
}
func GetPauseScriptOnDisconnectStatus() <-chan bool {
	p := NewBoolPacket(SCGetPauseScriptOnDisconnectStatus)
	p.send(senderFunc)
	return p.out
}
func SetARStatus(value bool) {
	p := NewVoidPacket(SCSetARStatus, value)
	p.send(senderFunc)
}
func GetARStatus() <-chan bool {
	p := NewBoolPacket(SCGetARStatus)
	p.send(senderFunc)
	return p.out
}
func CharName() <-chan string {
	p := NewStringPacket(SCGetCharName)
	p.send(senderFunc)
	return p.out
}
func ChangeProfile(profileName string) <-chan int32 {
	p := NewIntPacket(SCChangeProfile, profileName)
	p.send(senderFunc)
	return p.out
}
func ChangeProfileEx(pName, shardName, charName string) <-chan uint16 {
	p := NewUint16Packet(SCChangeProfileEx, pName, shardName, charName)
	p.send(senderFunc)
	return p.out
}
func ProfileName() <-chan string {
	p := NewStringPacket(SCGetProfileName)
	p.send(senderFunc)
	return p.out
}
func Self() <-chan uint32 {
	p := NewUint32Packet(SCGetSelfID)
	p.send(senderFunc)
	return p.out
}
func Sex() <-chan byte {
	p := NewBytePacket(SCGetSelfSex)
	p.send(senderFunc)
	return p.out
}
func GetCharTitle() <-chan string {
	p := NewStringPacket(SCGetCharTitle)
	p.send(senderFunc)
	return p.out
}
func Gold() <-chan uint16 {
	p := NewUint16Packet(SCGetSelfGold)
	p.send(senderFunc)
	return p.out
}
func Armor() <-chan uint16 {
	p := NewUint16Packet(SCGetSelfArmor)
	p.send(senderFunc)
	return p.out
}
func Weight() <-chan uint16 {
	p := NewUint16Packet(SCGetSelfWeight)
	p.send(senderFunc)
	return p.out
}
func MaxWeight() <-chan uint16 {
	p := NewUint16Packet(SCGetSelfMaxWeight)
	p.send(senderFunc)
	return p.out
}
func WorldNum() <-chan byte {
	p := NewBytePacket(SCGetWorldNum)
	p.send(senderFunc)
	return p.out
}
func Race() <-chan byte {
	p := NewBytePacket(SCGetSelfRace)
	p.send(senderFunc)
	return p.out
}
func MaxPets() <-chan byte {
	p := NewBytePacket(SCGetSelfPetsMax)
	p.send(senderFunc)
	return p.out
}
func PetsCurrent() <-chan byte {
	p := NewBytePacket(SCGetSelfPetsCurrent)
	p.send(senderFunc)
	return p.out
}
func FireResist() <-chan uint16 {
	p := NewUint16Packet(SCGetSelfFireResist)
	p.send(senderFunc)
	return p.out
}
func ColdResist() <-chan uint16 {
	p := NewUint16Packet(SCGetSelfColdResist)
	p.send(senderFunc)
	return p.out
}
func PoisonResist() <-chan uint16 {
	p := NewUint16Packet(SCGetSelfPoisonResist)
	p.send(senderFunc)
	return p.out
}
func EnergyResist() <-chan uint16 {
	p := NewUint16Packet(SCGetSelfEnergyResist)
	p.send(senderFunc)
	return p.out
}

func ConnectedTime() <-chan time.Time {
	p := NewTimePacket(SCGetConnectedTime)
	p.send(senderFunc)
	return p.out
}

func DisconnectedTime() <-chan time.Time {
	p := NewTimePacket(SCGetDisconnectedTime)
	p.send(senderFunc)
	return p.out
}

func LastContainer() <-chan uint32 {
	p := NewUint32Packet(SCGetLastContainer)
	p.send(senderFunc)
	return p.out
}

func LastTarget() <-chan uint32 {
	p := NewUint32Packet(SCGetLastTarget)
	p.send(senderFunc)
	return p.out
}

func LastAttack() <-chan uint32 {
	p := NewUint32Packet(SCGetLastAttack)
	p.send(senderFunc)
	return p.out
}

func LastStatus() <-chan uint32 {
	p := NewUint32Packet(43)
	p.send(senderFunc)
	return p.out
}

func LastObject() <-chan uint32 {
	p := NewUint32Packet(44)
	p.send(senderFunc)
	return p.out
} /*
_get_buff_bar_info = _ScriptMethod(349)  # GetBuffBarInfo
_get_buff_bar_info.restype = _buffer  # TBuffBarInfo
func GetBuffBarInfo(){
    result = []
}
    fmt = '<HdHII'
    size = _struct.calcsize(fmt)
    keys = ('Attribute_ID', 'TimeStart', 'Seconds', 'ClilocID1', 'ClilocID2')
    data = _get_buff_bar_info()
    if b'' == '':  # py2
        data = bytes(data)
    count = _uint.from_buffer(data)
    data = data[4:]
    for i in range(count):
        values = _struct.unpack(fmt, data[i * size:i * size + size])
        buff = dict(zip(keys, values))
        buff['TimeStart'] = _ddt2pdt(buff['TimeStart'])
        result.append(buff)
    return result*/
func ShardName() <-chan string {
	p := NewStringPacket(47)
	p.send(senderFunc)
	return p.out
}
func ProfileShardName() <-chan string {
	p := NewStringPacket(343)
	p.send(senderFunc)
	return p.out
}
func ProxyIP() <-chan string {
	p := NewStringPacket(60)
	p.send(senderFunc)
	return p.out
}
func ProxyPort() <-chan uint16 {
	p := NewUint16Packet(SCGetProxyPort)
	p.send(senderFunc)
	return p.out
}
func UseProxy() <-chan bool {
	p := NewBoolPacket(62)
	p.send(senderFunc)
	return p.out
}
func Backpack() <-chan uint32 {
	p := NewUint32Packet(48)
	p.send(senderFunc)
	return p.out
}
func Ground() byte {
	return 0
}
func Str() <-chan uint32 {
	p := NewUint32Packet(49)
	p.send(senderFunc)
	return p.out
}
func Int() <-chan uint32 {
	p := NewUint32Packet(50)
	p.send(senderFunc)
	return p.out
}
func Dex() <-chan uint32 {
	p := NewUint32Packet(51)
	p.send(senderFunc)
	return p.out
}
func Life() <-chan uint32 {
	p := NewUint32Packet(52)
	p.send(senderFunc)
	return p.out
}
func HP() <-chan uint32 {
	return Life()
}
func Mana() <-chan uint32 {
	p := NewUint32Packet(53)
	p.send(senderFunc)
	return p.out
}
func Stam() <-chan uint32 {
	p := NewUint32Packet(54)
	p.send(senderFunc)
	return p.out
}
func MaxLife() <-chan uint32 {
	p := NewUint32Packet(55)
	p.send(senderFunc)
	return p.out
}
func MaxHP() <-chan uint32 {
	return MaxLife()
}
func MaxMana() <-chan uint32 {
	p := NewUint32Packet(56)
	p.send(senderFunc)
	return p.out
}
func MaxStam() <-chan uint32 {
	p := NewUint32Packet(57)
	p.send(senderFunc)
	return p.out
}
func Luck() <-chan uint16 {
	p := NewUint16Packet(SCGetSelfLuck)
	p.send(senderFunc)
	return p.out
} /*
_get_extended_info = _ScriptMethod(59)  # GetExtInfo
_get_extended_info.restype = _buffer  # TExtendedInfo
func GetExtInfo(){
    keys = ('MaxWeight', 'Race', 'StatCap', 'PetsCurrent', 'PetsMax',
}
            'FireResist', 'ColdResist', 'PoisonResist', 'EnergyResist',
            'Luck', 'DamageMin', 'DamageMax', 'Tithing_points',
            'ArmorMax', 'fireresistMax', 'coldresistMax',
            'poisonresistMax', 'energyresistMax', 'DefenseChance',
            'DefensceChanceMax', 'Hit_Chance_Incr', 'Damage_Incr',
            'Swing_Speed_Incr', 'Lower_Reagent_Cost', 'Spell_Damage_Incr',
            'Faster_Cast_Recovery', 'Faster_Casting', 'Lower_Mana_Cost',
            'HP_Regen', 'Stam_Regen', 'Mana_Regen', 'Reflect_Phys_Damage',
            'Enhance_Potions', 'Strength_Incr', 'Dex_Incr', 'Int_Incr',
            'HP_Incr', 'Mana_Incr')
    fmt = '<HBH2B4Hh2Hi26H'
    data = _get_extended_info()
    if b'' == '':  # py2
        data = bytes(data)
    values = _struct.unpack(fmt, data)
    p :=
p.send(senderFunc)
// return dict(zip(keys, values))
*/
func Hidden() <-chan bool {
	p := NewBoolPacket(63)
	p.send(senderFunc)
	return p.out
}

func Poisoned() <-chan bool {
	p := NewBoolPacket(64)
	p.send(senderFunc)
	return p.out
}

func Paralyzed() <-chan bool {
	p := NewBoolPacket(65)
	p.send(senderFunc)
	return p.out
}

func Dead() <-chan bool {
	p := NewBoolPacket(66)
	p.send(senderFunc)
	return p.out
}

func WarMode() <-chan bool {
	p := NewBoolPacket(171, <-Self())
	p.send(senderFunc)
	return p.out
}

func WarTargetID() <-chan uint32 {
	p := NewUint32Packet(SCGetWarTarget)
	p.send(senderFunc)
	return p.out
}

func SetWarMode(value bool) {
	p := NewVoidPacket(SCSetWarMode, value)
	p.send(senderFunc)
}

func Attack(attackedID uint32) {
	p := NewVoidPacket(SCAttack, attackedID)
	p.send(senderFunc)
}

func UseSelfPaperdollScroll() {
	p := NewVoidPacket(SCUseSelfPaperdollScroll)
	p.send(senderFunc)
}

func UseOtherPaperdollScroll(oid uint32) {
	p := NewVoidPacket(SCUseOtherPaperdollScroll, oid)
	p.send(senderFunc)
}

func TargetID() <-chan uint32 {
	p := NewUint32Packet(72)
	p.send(senderFunc)
	return p.out
}

func CancelTarget() {
	p := NewVoidPacket(73)
	p.send(senderFunc)
}

func TargetToObject(ObjectID uint32) {
	p := NewVoidPacket(74, ObjectID)
	p.send(senderFunc)
}

func TargetToXYZ(x, y uint16, z byte) {
	p := NewVoidPacket(SCTargetToXYZ, x, y, z)
	p.send(senderFunc)
}

func TargetToTile(tileModel, x, y uint16, z byte) {
	p := NewVoidPacket(SCTargetToTile, x, y, z)
	p.send(senderFunc)
}

func WaitTargetObject(ObjID uint32) {
	p := NewVoidPacket(77, ObjID)
	p.send(senderFunc)
}

func WaitTargetTile(tile, x, y uint16, z byte) {
	p := NewVoidPacket(SCWaitTargetTile, tile, x, y, z)
	p.send(senderFunc)
}

func WaitTargetXYZ(x, y uint16, z byte) {
	p := NewVoidPacket(SCWaitTargetXYZ, x, y, z)
	p.send(senderFunc)
}

func WaitTargetSelf() {
	p := NewVoidPacket(80)
	p.send(senderFunc)
}

func WaitTargetType(ObjType uint16) {
	p := NewVoidPacket(81, ObjType)
	p.send(senderFunc)
}

func CancelWaitTarget() {
	p := NewVoidPacket(82)
	p.send(senderFunc)
}

func WaitTargetGround(ObjType uint16) {
	p := NewVoidPacket(83, ObjType)
	p.send(senderFunc)
}

func WaitTargetLast() {
	p := NewVoidPacket(84)
	p.send(senderFunc)
}

func UsePrimaryAbility() {
	p := NewVoidPacket(85)
	p.send(senderFunc)
}

func UseSecondaryAbility() {
	p := NewVoidPacket(86)
	p.send(senderFunc)
}

func GetActiveAbility() <-chan string {
	p := NewStringPacket(87)
	p.send(senderFunc)
	return p.out
}

func ToggleFly() {
	p := NewVoidPacket(88)
	p.send(senderFunc)
}

func getSkillId(skillName string) <-chan uint32 {
	p := NewUint32Packet(SCGetSkillID, skillName)
	p.send(senderFunc)
	return p.out
}

func UseSkill(skillName string) {
	p := NewVoidPacket(SCUseSkill, <-getSkillId(skillName))
	p.send(senderFunc)
}

func ChangeSkillLockState(skillName string, skillState byte) {
	p := NewVoidPacket(SCChangeSkillLockState, <-getSkillId(skillName), skillState)
	p.send(senderFunc)
}

func SetSkillLockState(skillName string, skillState byte) {
	ChangeSkillLockState(skillName, skillState)
}

func GetSkillCap(SkillName string) <-chan float64 {
	p := NewFloatPacket(92, <-getSkillId(SkillName))
	p.send(senderFunc)
	return p.out
}
func GetSkillValue(SkillName string) <-chan float64 {
	p := NewFloatPacket(93, <-getSkillId(SkillName))
	p.send(senderFunc)
	return p.out
}
func GetSkillCurrentValue(SkillName string) <-chan float64 {
	p := NewFloatPacket(351, <-getSkillId(SkillName))
	p.send(senderFunc)
	return p.out
}

func ReqVirtuesGump() {
	p := NewVoidPacket(94)
	p.send(senderFunc)
}

func UseVirtue(VirtueName string) {
	if v, ok := _VIRTUES[VirtueName]; ok {
		p := NewVoidPacket(95, v)
		p.send(senderFunc)
	} else {
		log.Fatalf("Unknown virtue %v", VirtueName)
		os.Exit(500)
	}
}

func Cast(spellName string) {
	p := NewVoidPacket(SCCastSpell, _SPELLS[strings.ToLower(spellName)])
	p.send(senderFunc)
}

func CastToObj(spellName string, oid uint32) {
	WaitTargetObject(oid)
	Cast(spellName)
}

func IsActiveSpellAbility(spellName string) <-chan bool {
	p := NewBoolPacket(SCIsActiveSpellAbility, _SPELLS[strings.ToLower(spellName)])
	p.send(senderFunc)
	return p.out
}

func UnsetCatchBag() {
	p := NewVoidPacket(100)
	p.send(senderFunc)
}

func SetCatchBag(ObjectID uint32) {
	p := NewVoidPacket(99, ObjectID)
	p.send(senderFunc)
}

func UseObject(ObjectID uint32) {
	p := NewVoidPacket(101, ObjectID)
	p.send(senderFunc)
}

func UseType(objType uint16, color uint16) <-chan uint32 {
	p := NewUint32Packet(SCUseType, objType, color)
	p.send(senderFunc)
	return p.out
}

func UseFromGround(objType, color uint16) <-chan uint32 {
	p := NewUint32Packet(SCUseFromGround, objType, color)
	p.send(senderFunc)
	return p.out
}

func ClickOnObject(ObjectID uint32) {
	p := NewVoidPacket(104, ObjectID)
	p.send(senderFunc)
}

func FoundedParamID() <-chan uint32 {
	p := NewUint32Packet(105)
	p.send(senderFunc)
	return p.out
}

func LineID() <-chan uint32 {
	p := NewUint32Packet(106)
	p.send(senderFunc)
	return p.out
}

func LineType() <-chan uint16 {
	p := NewUint16Packet(SCGetLineType)
	p.send(senderFunc)
	return p.out
}

func LineName() <-chan string {
	p := NewStringPacket(114)
	p.send(senderFunc)
	return p.out
}

func LineTime() <-chan time.Time {
	p := NewTimePacket(SCGetLineTime)
	p.send(senderFunc)
	return p.out
}

func LineMsgType() <-chan byte {
	p := NewBytePacket(SCGetLineMsgType)
	p.send(senderFunc)
	return p.out
}

func LineTextColor() <-chan uint16 {
	p := NewUint16Packet(SCGetLineTextColor)
	p.send(senderFunc)
	return p.out
}

func LineTextFont() <-chan uint16 {
	p := NewUint16Packet(SCGetLineTextFont)
	p.send(senderFunc)
	return p.out
}

func LineIndex() <-chan uint32 {
	p := NewUint32Packet(112)
	p.send(senderFunc)
	return p.out
}

func LineCount() <-chan uint32 {
	p := NewUint32Packet(113)
	p.send(senderFunc)
	return p.out
}

func AddJournalIgnore(str string) {
	p := NewVoidPacket(115, str)
	p.send(senderFunc)
}

func ClearJournalIgnore() {
	p := NewVoidPacket(116)
	p.send(senderFunc)
}

func AddChatUserIgnore(user string) {
	p := NewVoidPacket(117, user)
	p.send(senderFunc)
}

func AddToJournal(msg string) {
	p := NewVoidPacket(304, msg)
	p.send(senderFunc)
}

func ClearChatUserIgnore() {
	p := NewVoidPacket(118)
	p.send(senderFunc)
}

func ClearJournal() {
	p := NewVoidPacket(119)
	p.send(senderFunc)
}

func ClearSystemJournal() {
	p := NewVoidPacket(346)
	p.send(senderFunc)
}

func LastJournalMessage() <-chan string {
	p := NewStringPacket(120)
	p.send(senderFunc)
	return p.out
}

func InJournal(s string) <-chan int32 {
	p := NewIntPacket(SCInJournal, s)
	p.send(senderFunc)
	return p.out
}

// InJournalBetweenTimes
//
// RU: Поиск последней строки в журнале по слову (или по словам) во временном интервале.
// Если строка не найдена возвратит -1, если найдена,
// возвратит индекс строки в журнале начиная с 0.
//
// EN: Search for last entry in journal by word(words) in time interval
// Returns if string is found - index of string starting from 0
// if string is not found -1
func InJournalBetweenTimes(str string, timeBegin time.Time, timeEnd time.Time) <-chan int32 {
	p := NewIntPacket(SCInJournalBetweenTimes, str, timeBegin, timeEnd)
	p.send(senderFunc)
	return p.out
}

func Journal(stringIndex uint32) <-chan string {
	p := NewStringPacket(SCJournal, stringIndex)
	p.send(senderFunc)
	return p.out
}

func SetJournalLine(stringIndex uint32, text string) {
	p := NewVoidPacket(SCSetJournalLine, stringIndex, text)
	p.send(senderFunc)
}

func LowJournal() <-chan uint32 {
	p := NewUint32Packet(125)
	p.send(senderFunc)
	return p.out
}

func HighJournal() <-chan uint32 {
	p := NewUint32Packet(126)
	p.send(senderFunc)
	return p.out
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
	p := NewVoidPacket(127, Value)
	p.send(senderFunc)
}

func GetFindDistance() <-chan uint32 {
	p := NewUint32Packet(128)
	p.send(senderFunc)
	return p.out
}

func SetFindVertical(Value uint32) {
	p := NewVoidPacket(129, Value)
	p.send(senderFunc)
}

func GetFindVertical() <-chan uint32 {
	p := NewUint32Packet(130)
	p.send(senderFunc)
	return p.out
}

func SetFindInNulPoint(v bool) {
	p := NewVoidPacket(SCSetFindInNulPoint, v)
	p.send(senderFunc)
}

func GetFindInNulPoint() <-chan bool {
	p := NewBoolPacket(337)
	p.send(senderFunc)
	return p.out
}

func FindTypeEx(objType, objColor uint16, container uint32, inSub bool) <-chan uint32 {
	p := NewUint32Packet(SCFindTypeEx, objType, objColor, container, inSub)
	p.send(senderFunc)
	return p.out
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
	p := NewUint32Packet(SCFindTypesArrayEx, objTypes, colors, containers, inSub)
	p.send(senderFunc)
	return p.out
}

func FindNotoriety(objType uint16, notoriety byte) <-chan uint32 {
	p := NewUint32Packet(SCFindNotoriety, objType, notoriety)
	p.send(senderFunc)
	return p.out
}

func FindAtCoord(x, y uint16) <-chan uint32 {
	p := NewUint32Packet(SCFindAtCoord, x, y)
	p.send(senderFunc)
	return p.out

}
func Ignore(ObjID uint32) {
	p := NewVoidPacket(134, ObjID)
	p.send(senderFunc)
}

func IgnoreOff(ObjID uint32) {
	p := NewVoidPacket(135, ObjID)
	p.send(senderFunc)
}

func IgnoreReset() {
	p := NewVoidPacket(136)
	p.send(senderFunc)
}

func GetIgnoreList() <-chan []uint32 {
	p := NewUint32ArrayPacket(SCGetIgnoreList)
	p.send(senderFunc)
	return p.out
}

func GetFoundList() <-chan []uint32 {
	p := NewUint32ArrayPacket(SCGetFindedList)
	p.send(senderFunc)
	return p.out
}

func FindItem() <-chan uint32 {
	p := NewUint32Packet(139)
	p.send(senderFunc)
	return p.out
}

func FindCount() <-chan uint32 {
	p := NewUint32Packet(140)
	p.send(senderFunc)
	return p.out
}

func FindQuantity() <-chan uint32 {
	p := NewUint32Packet(141)
	p.send(senderFunc)
	return p.out
}

func FindFullQuantity() <-chan uint32 {
	p := NewUint32Packet(142)
	p.send(senderFunc)
	return p.out
}

func PredictedX() <-chan uint16 {
	p := NewUint16Packet(143)
	p.send(senderFunc)
	return p.out
}

func PredictedY() <-chan uint16 {
	p := NewUint16Packet(144)
	p.send(senderFunc)
	return p.out
}

func PredictedZ() <-chan byte {
	p := NewBytePacket(SCPredictedZ)
	p.send(senderFunc)
	return p.out
}

func PredictedDirection() <-chan byte {
	p := NewBytePacket(SCPredictedDirection)
	p.send(senderFunc)
	return p.out
}

func GetX(oid uint32) <-chan uint16 {
	p := NewUint16Packet(SCGetX, oid)
	p.send(senderFunc)
	return p.out
}

func GetY(oid uint32) <-chan uint16 {
	p := NewUint16Packet(SCGetY, oid)
	p.send(senderFunc)
	return p.out
}

func GetZ(oid uint32) <-chan int8 {
	p := NewInt8Packet(SCGetZ, oid)
	p.send(senderFunc)
	return p.out
}

func GetName(oid uint32) <-chan string {
	p := NewStringPacket(SCGetName, oid)
	p.send(senderFunc)
	return p.out
}

func GetAltName(oid uint32) <-chan string {
	p := NewStringPacket(SCGetAltName, oid)
	p.send(senderFunc)
	return p.out

}

func GetTitle(oid uint32) <-chan string {
	p := NewStringPacket(SCGetTitle)
	p.send(senderFunc)
	return p.out
}

func GetTooltip(oid uint32) <-chan string {
	p := NewStringPacket(SCGetCliloc, oid)
	p.send(senderFunc)
	return p.out
}

func GetCliloc(oid uint32) <-chan string {
	p := NewStringPacket(SCGetCliloc, oid)
	p.send(senderFunc)
	return p.out
}

func GetType(oid uint32) <-chan uint16 {
	p := NewUint16Packet(SCGetType, oid)
	p.send(senderFunc)
	return p.out
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
        result.append({'Cliloc_ID': cliloc, 'Params': strings})
    return result
*/
func GetClilocByID(oid uint32) <-chan string {
	p := NewStringPacket(SCGetClilocByID, oid)
	p.send(senderFunc)
	return p.out
}

func GetQuantity(oid uint32) <-chan int32 {
	p := NewIntPacket(SCGetQuantity, oid)
	p.send(senderFunc)
	return p.out
}

func IsObjectExists(oid uint32) <-chan bool {
	p := NewBoolPacket(SCIsObjectExists, oid)
	p.send(senderFunc)
	return p.out
}

func IsNPC(oid uint32) <-chan bool {
	p := NewBoolPacket(SCIsNPC, oid)
	p.send(senderFunc)
	return p.out
}

func GetPrice(oid uint32) <-chan uint32 {
	p := NewUint32Packet(SCGetPrice, oid)
	p.send(senderFunc)
	return p.out
}

func GetDirection(oid uint32) <-chan byte {
	p := NewBytePacket(SCGetDirection, oid)
	p.send(senderFunc)
	return p.out
}

func GetDistance(oid uint32) <-chan int32 {
	p := NewIntPacket(SCGetDistance, oid)
	p.send(senderFunc)
	return p.out
}

func GetColor(oid uint32) <-chan uint16 {
	p := NewUint16Packet(SCGetColor, oid)
	p.send(senderFunc)
	return p.out
}

func GetStr(oid uint32) <-chan int32 {
	p := NewIntPacket(SCGetStr, oid)
	p.send(senderFunc)
	return p.out
}

func GetInt(oid uint32) <-chan int32 {
	p := NewIntPacket(SCGetInt, oid)
	p.send(senderFunc)
	return p.out
}

func GetDex(oid uint32) <-chan int32 {
	p := NewIntPacket(SCGetDex, oid)
	p.send(senderFunc)
	return p.out
}

func GetHP(oid uint32) <-chan int32 {
	p := NewIntPacket(SCGetHP, oid)
	p.send(senderFunc)
	return p.out
}

func GetMaxHP(oid uint32) <-chan int32 {
	p := NewIntPacket(SCGetMaxHP, oid)
	p.send(senderFunc)
	return p.out
}

func GetMana(oid uint32) <-chan int32 {
	p := NewIntPacket(SCGetMana, oid)
	p.send(senderFunc)
	return p.out
}

func GetMaxMana(ObjID uint32) <-chan int32 {
	p := NewIntPacket(166, ObjID)
	p.send(senderFunc)
	return p.out
}

func GetStam(oid uint32) <-chan int32 {
	p := NewIntPacket(SCGetStam, oid)
	p.send(senderFunc)
	return p.out
}
func GetMaxStam(ObjID uint32) <-chan int32 {
	p := NewIntPacket(168, ObjID)
	p.send(senderFunc)
	return p.out
}

func GetNotoriety(oid uint32) <-chan byte {
	p := NewBytePacket(SCGetNotoriety)
	p.send(senderFunc)
	return p.out
}

func GetParent(oid uint32) <-chan uint32 {
	p := NewUint32Packet(SCGetParent, oid)
	p.send(senderFunc)
	return p.out
}

func IsWarMode(oid uint32) <-chan bool {
	p := NewBoolPacket(SCIsWarMode, oid)
	p.send(senderFunc)
	return p.out
}

func IsDead(ObjID uint32) <-chan bool {
	p := NewBoolPacket(173, ObjID)
	p.send(senderFunc)
	return p.out
}

func IsRunning(ObjID uint32) <-chan bool {
	p := NewBoolPacket(174, ObjID)
	p.send(senderFunc)
	return p.out
}

func IsContainer(ObjID uint32) <-chan bool {
	p := NewBoolPacket(175, ObjID)
	p.send(senderFunc)
	return p.out
}

func IsHidden(ObjID uint32) <-chan bool {
	p := NewBoolPacket(176, ObjID)
	p.send(senderFunc)
	return p.out
}

func IsMovable(ObjID uint32) <-chan bool {
	p := NewBoolPacket(177, ObjID)
	p.send(senderFunc)
	return p.out
}

func IsYellowHits(ObjID uint32) <-chan bool {
	p := NewBoolPacket(178, ObjID)
	p.send(senderFunc)
	return p.out
}

func IsPoisoned(ObjID uint32) <-chan bool {
	p := NewBoolPacket(179, ObjID)
	p.send(senderFunc)
	return p.out
}

func IsParalyzed(ObjID uint32) <-chan bool {
	p := NewBoolPacket(180, ObjID)
	p.send(senderFunc)
	return p.out
}

func IsFemale(ObjID uint32) <-chan bool {
	p := NewBoolPacket(181, ObjID)
	p.send(senderFunc)
	return p.out
}

func OpenDoor() {
	p := NewVoidPacket(182)
	p.send(senderFunc)
}

func Bow() {
	p := NewVoidPacket(183)
	p.send(senderFunc)
}

func Salute() {
	p := NewVoidPacket(184)
	p.send(senderFunc)
}

func GetPickupedItem() <-chan uint32 {
	p := NewUint32Packet(185)
	p.send(senderFunc)
	return p.out
}

func SetPickupedItem(ID uint32) {
	p := NewVoidPacket(186, ID)
	p.send(senderFunc)
}

func GetDropCheckCoord() <-chan bool {
	p := NewBoolPacket(187)
	p.send(senderFunc)
	return p.out
}

func SetDropCheckCoord(value bool) {
	p := NewVoidPacket(SCSetDropCheckCoord)
	p.send(senderFunc)
}

func GetDropDelay() <-chan uint32 {
	p := NewUint32Packet(189)
	p.send(senderFunc)
	return p.out
}

func SetDropDelay(Value uint32) {
	p := NewVoidPacket(190, Value)
	p.send(senderFunc)
}

/*
_drag_item = _ScriptMethod(191)  # DragItem
_drag_item.restype = _bool
_drag_item.argtypes = [_uint,  # ItemID
                       _int]  # Count
func DragItem(ItemID, Count){
    p :=
p.send(senderFunc)
// return _drag_item(ItemID, Count)
}
_drop_item = _ScriptMethod(192)  # DropItem
_drop_item.restype = _bool
_drop_item.argtypes = [_uint,  # MoveIntoID
                       _int,  # X
                       _int,  # Y
                       _int]  # Z
func DropItem(MoveIntoID, X, Y, Z){
    p :=
p.send(senderFunc)
// return _drop_item(MoveIntoID, X, Y, Z)
}
func MoveItem(ItemID, Count, MoveIntoID, X, Y, Z){
    if not DragItem(ItemID, Count):
}
        return False
    Wait(100)
    p :=
p.send(senderFunc)
// return DropItem(MoveIntoID, X, Y, Z)
func Grab(ItemID, Count){
    p :=
p.send(senderFunc)
// return MoveItem(ItemID, Count, Backpack(), 0, 0, 0)
}
func Drop(ItemID, Count, X, Y, Z){
    p :=
p.send(senderFunc)
// return MoveItem(ItemID, Count, Ground(), X, Y, Z)
}
func DropHere(ItemID){
    p :=
p.send(senderFunc)
// return MoveItem(ItemID, 0, Ground(), 0, 0, 0)
}
def MoveItems(Container, ItemsType, ItemsColor, MoveIntoID, X, Y, Z,
              DelayMS, MaxCount=0):
    FindTypeEx(ItemsType, ItemsColor, Container, False)
    items = GetFoundList()
    if not items:  # nothing found
        return False
    drop_delay = GetDropDelay()
    if not 50 <= drop_delay <= 10000:
        drop_delay = 50 if drop_delay < 50 else 10000
    if drop_delay > DelayMS:
        DelayMS = 0
    SetDropDelay(drop_delay)
    if not 0 < MaxCount < len(items):
        MaxCount = len(items)
    for i in range(MaxCount):
        MoveItem(items[i], 0, MoveIntoID, X, Y, Z)
        Wait(DelayMS)
    return True
func EmptyContainer(Container, DestContainer, delay_ms){
    return MoveItems(Container, -1, -1, DestContainer,
}
                     0xFFFF, 0xFFFF, 0, delay_ms)
*/
func RequestContextMenu(ID uint32) {
	p := NewVoidPacket(193, ID)
	p.send(senderFunc)
}

/*
_wait_context_menu = _ScriptMethod(194)  # SetContextMenuHook
_wait_context_menu.argtypes = [_uint,  # MenuID
                               _ubyte]  # EntryNumber
func SetContextMenuHook(MenuID, EntryNumber){
    _wait_context_menu(MenuID, EntryNumber)
}
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
        result.append(string.value)
    return result
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
    """
    return None
*/
func ClearContextMenu() {
	p := NewVoidPacket(196)
	p.send(senderFunc)
}

func IsTrade() <-chan bool {
	p := NewBoolPacket(197)
	p.send(senderFunc)
	return p.out
}

/*
_get_trade_container_serial = _ScriptMethod(198)  # GetTradeContainer
_get_trade_container_serial.restype = _uint
_get_trade_container_serial.argtypes = [_ubyte,  # TradeNum
                                        _ubyte]  # Num
func GetTradeContainer(TradeNum, Num){
    p :=
p.send(senderFunc)
// return _get_trade_container_serial(TradeNum, Num)
}
_get_trade_opponent_serial = _ScriptMethod(199)  # GetTradeOpponent
_get_trade_opponent_serial.restype = _uint
_get_trade_opponent_serial.argtypes = [_ubyte]  # TradeNum
func GetTradeOpponent(TradeNum){
    p :=
p.send(senderFunc)
// return _get_trade_opponent_serial(TradeNum)
}
_get_trades_count = _ScriptMethod(200)  # GetTradeCount
_get_trades_count.restype = _ubyte
func TradeCount(){
    return _get_trades_count()
}
_get_trade_opponent_name = _ScriptMethod(201)  # GetTradeOpponentName
_get_trade_opponent_name.restype = _str
_get_trade_opponent_name.argtypes = [_ubyte]  # TradeNum
func GetTradeOpponentName(TradeNum){
    p :=
p.send(senderFunc)
// return _get_trade_opponent_name(TradeNum)
}
_get_trade_state = _ScriptMethod(202)  # TradeCheck
_get_trade_state.restype = _bool
_get_trade_state.argtypes = [_ubyte,  # TradeNum
                             _ubyte]  # Num
func TradeCheck(TradeNum, Num){
    p :=
p.send(senderFunc)
// return _get_trade_state(TradeNum, Num)
}
_confirm_trade = _ScriptMethod(203)  # ConfirmTrade
_confirm_trade.argtypes = [_ubyte]  # TradeNum
func ConfirmTrade(TradeNum){
    _confirm_trade(TradeNum)
}
_cancel_trade = _ScriptMethod(204)  # CancelTrade
_cancel_trade.restype = _bool
_cancel_trade.argtypes = [_ubyte]  # TradeNum
func CancelTrade(TradeNum){
    p :=
p.send(senderFunc)
// return _cancel_trade(TradeNum)
}
_wait_menu = _ScriptMethod(205)  # WaitMenu
_wait_menu.argtypes = [_str,  # MenuCaption
                       _str]  # ElementCaption
func WaitMenu(MenuCaption, ElementCaption){
    _wait_menu(MenuCaption, ElementCaption)
}
_auto_menu = _ScriptMethod(206)  # AutoMenu
_auto_menu.argtypes = [_str,  # MenuCaption
                       _str]  # ElementCaption
func AutoMenu(MenuCaption, ElementCaption){
    _auto_menu(MenuCaption, ElementCaption)
}
*/
func MenuHookPresent() <-chan bool {
	p := NewBoolPacket(207)
	p.send(senderFunc)
	return p.out
}

func MenuPresent() <-chan bool {
	p := NewBoolPacket(208)
	p.send(senderFunc)
	return p.out
}

func CancelMenu() {
	p := NewVoidPacket(209)
	p.send(senderFunc)
}

func CancelAllMenuHooks() {
	CancelMenu()
}

func CloseMenu() {
	p := NewVoidPacket(210)
	p.send(senderFunc)
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
        offset += string.size
    return result
func GetMenuItems(MenuCaption){
    p :=
p.send(senderFunc)
// return '\n'.join(GetMenu(MenuCaption))
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
        offset += string.size
    return result
func GetLastMenuItems(){
    p :=
p.send(senderFunc)
// return '\n'.join(GetLastMenu())
}
_wait_gump = _ScriptMethod(211)  # WaitGumpInt
_wait_gump.argtypes = [_int]  # Value
func WaitGump(Value){
    _wait_gump(int(Value))
}
*/
func WaitTextEntry(value string) {
	p := NewVoidPacket(212, value)
	p.send(senderFunc)
}

/*
_auto_text_entry = _ScriptMethod(213)  # GumpAutoTextEntry
_auto_text_entry.argtypes = [_int,  # TextEntryID
                             _str]  # Value
func GumpAutoTextEntry(TextEntryID, Value){
    _auto_text_entry(TextEntryID, Value)
}
_auto_radiobutton = _ScriptMethod(214)  # GumpAutoRadiobutton
_auto_radiobutton.argtypes = [_int,  # RadiobuttonID
                              _int]  # Value
func GumpAutoRadiobutton(RadiobuttonID, Value){
    _auto_radiobutton(RadiobuttonID, Value)
}
_auto_checkbox = _ScriptMethod(215)  # GumpAutoCheckBox
_auto_checkbox.argtypes = [_int,  # CBID
                           _int]  # Value
func GumpAutoCheckBox(CBID, Value){
    _auto_checkbox(CBID, Value)
}
_send_gump_button = _ScriptMethod(216)  # NumGumpButton
_send_gump_button.restype = _bool
_send_gump_button.argtypes = [_ushort,  # GumpIndex
                              _int]  # Value
func NumGumpButton(GumpIndex, Value){
    p :=
p.send(senderFunc)
// return _send_gump_button(GumpIndex, Value)
}
_send_gump_text_entry = _ScriptMethod(217)  # NumGumpTextEntry
_send_gump_text_entry.restype = _bool
_send_gump_text_entry.argtypes = [_ushort,  # GumpIndex
                                  _int,  # TextEntryID
                                  _str]  # Value
func NumGumpTextEntry(GumpIndex, TextEntryID, Value){
    p :=
p.send(senderFunc)
// return _send_gump_text_entry(GumpIndex, TextEntryID, Value)
}
_send_gump_radiobutton = _ScriptMethod(218)  # NumGumpRadiobutton
_send_gump_radiobutton.restype = _bool
_send_gump_radiobutton.argtypes = [_ushort,  # GumpIndex
                                   _int,  # RadiobuttonID
                                   _int]  # Value
func NumGumpRadiobutton(GumpIndex, RadiobuttonID, Value){
    p :=
p.send(senderFunc)
// return _send_gump_radiobutton(GumpIndex, RadiobuttonID, Value)
}
_send_gump_checkbox = _ScriptMethod(219)  # NumGumpCheckBox
_send_gump_checkbox.restype = _bool
_send_gump_checkbox.argtypes = [_ushort,  # GumpIndex
                                _int,  # CBID
                                _int]  # Value
func NumGumpCheckBox(GumpIndex, CBID, Value){
    p :=
p.send(senderFunc)
// return _send_gump_checkbox(GumpIndex, CBID, Value)
}
*/
func GetGumpsCount() <-chan uint16 {
	p := NewUint16Packet(220)
	p.send(senderFunc)
	return p.out
}

/*
 */
func CloseSimpleGump(GumpIndex uint16) {
	p := NewVoidPacket(221, GumpIndex)
	p.send(senderFunc)
}

/*
func IsGump(){
    return GetGumpsCount() > 0
}
_get_gump_serial = _ScriptMethod(222)  # GetGumpSerial
_get_gump_serial.restype = _uint
_get_gump_serial.argtypes = [_ushort]  # GumpIndex
func GetGumpSerial(GumpIndex){
    p :=
p.send(senderFunc)
// return _get_gump_serial(GumpIndex)
}
_get_gump_type = _ScriptMethod(223)  # GetGumpID
_get_gump_type.restype = _uint
_get_gump_type.argtypes = [_ushort]  # GumpIndex
func GetGumpID(GumpIndex){
    p :=
p.send(senderFunc)
// return _get_gump_type(GumpIndex)
}
_get_gump_no_close = _ScriptMethod(224)  # GetGumpNoClose
_get_gump_no_close.restype = _bool
_get_gump_no_close.argtypes = [_ushort]  # GumpIndex
func IsGumpCanBeClosed(GumpIndex){
    p :=
p.send(senderFunc)
// return _get_gump_no_close(GumpIndex)
}
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
        result.append(string.value)
    return result
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
        result.append(string.value)
    return result
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
        result.append(string.value)
    return result
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
        result.append(string.value)
    return result
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
    keys = ('X', 'Y', 'ReleasedID', 'PressedID', 'Quit', 'PageID',
            'ReturnValue', 'Page', 'ElemNum')
class _ButtonTileArt:
    args = [_int] * 12
    container = 'ButtonTileArts'
    keys = ('X', 'Y', 'ReleasedID', 'PressedID', 'Quit', 'PageID',
            'ReturnValue', 'ArtID', 'Hue', 'ArtX', 'ArtY', 'ElemNum')
class _CheckBox:
    args = [_int] * 8
    container = 'CheckBoxes'
    keys = ('X', 'Y', 'ReleasedID', 'PressedID', 'Status', 'ReturnValue',
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
    container = 'RadioButtons'
    keys = ('X', 'Y', 'ReleasedID', 'PressedID', 'Status', 'ReturnValue',
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
    container = 'TextEntries'
    keys = ('X', 'Y', 'Width', 'Height', 'Color', 'ReturnValue',
            'DefaultTextID', 'RealValue', 'Page', 'ElemNum')
class _Text:
    args = [_str]
    container = 'Text'
    keys = None
class _TextEntryLimited:
    args = [_int] * 10
    container = 'TextEntriesLimited'
    keys = ('X', 'Y', 'Width', 'Height', 'Color', 'ReturnValue',
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
                result[cls.container].append(element)
    return result
*/
func AddGumpIgnoreByID(ID uint32) {
	p := NewVoidPacket(230, ID)
	p.send(senderFunc)
}

func AddGumpIgnoreBySerial(Serial uint32) {
	p := NewVoidPacket(231, Serial)
	p.send(senderFunc)
}

func ClearGumpsIgnore() {
	p := NewVoidPacket(232)
	p.send(senderFunc)
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
	p := NewUint32Packet(SCObjAtLayerEx, layerType, playerID)
	p.send(senderFunc)
	return p.out
}

func ObjAtLayer(LayerType byte) <-chan uint32 {
	return ObjAtLayerEx(LayerType, <-Self())
}

func GetLayer(Obj uint32) <-chan byte {
	p := NewBytePacket(SCGetLayer)
	p.send(senderFunc)
	return p.out
}

func WearItem(layer byte, oid uint32) {
	p := NewVoidPacket(SCWearItem, layer, oid)
	p.send(senderFunc)
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
p.send(senderFunc)
// return all(tmp)
func disarm(){
    return Disarm()
}
func Equip(Layer, Obj){
    if Layer and DragItem(Obj, 1):
}
        p :=
p.send(senderFunc)
// return WearItem(Layer, Obj)
    return False
func equip(Layer, Obj){
    p :=
p.send(senderFunc)
// return Equip(Layer, Obj)
}
func Equipt(Layer, ObjType){
    item = FindType(ObjType, Backpack())
}
    if item:
        p :=
p.send(senderFunc)
// return Equip(Layer, item)
    return False
func equipt(Layer, ObjType){
    p :=
p.send(senderFunc)
// return Equipt(Layer, ObjType)
}
func UnEquip(Layer){
    item = ObjAtLayer(Layer)
}
    if item:
        p :=
p.send(senderFunc)
// return MoveItem(item, 1, Backpack(), 0, 0, 0)
    return False
*/
func GetDressSpeed() <-chan uint16 {
	p := NewUint16Packet(236)
	p.send(senderFunc)
	return p.out
}

func SetDressSpeed(Value uint16) {
	p := NewVoidPacket(237, Value)
	p.send(senderFunc)
}

func GetClientVersionInt() <-chan uint32 {
	p := NewUint32Packet(355)
	p.send(senderFunc)
	return p.out
}

/*
_wearable_layers = (RhandLayer(), LhandLayer(), ShoesLayer(), PantsLayer(),
                    ShirtLayer(), HatLayer(), GlovesLayer(), RingLayer(),
                    NeckLayer(), WaistLayer(), TorsoLayer(), BraceLayer(),
                    TorsoHLayer(), EarLayer(), ArmsLayer(), CloakLayer(),
*/
func UnequipItemsSetMacro() {
	p := NewVoidPacket(356)
	p.send(senderFunc)
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
p.send(senderFunc)
// return all(tmp)
*/
func SetDress() {
	p := NewVoidPacket(238)
	p.send(senderFunc)
}

func EquipItemsSetMacro() {
	p := NewVoidPacket(357)
	p.send(senderFunc)
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
p.send(senderFunc)
// return all(res)
func DressSavedSet(){
    EquipDressSet()
}
func Count(ObjType){
    FindType(ObjType, Backpack())
}
    return FindFullQuantity()
func CountGround(ObjType){
    FindType(ObjType, Ground())
}
    return FindFullQuantity()
func CountEx(ObjType, Color, Container){
    FindTypeEx(ObjType, Color, Container, False)
}
    return FindFullQuantity()

*/
func BP() uint16 { return 0x0F7A }
func BM() uint16 { return 0x0F7B }
func GA() uint16 { return 0x0F84 }
func GS() uint16 { return 0x0F85 }
func MR() uint16 { return 0x0F86 }
func NS() uint16 { return 0x0F88 }
func SA() uint16 { return 0x0F8C }
func SS() uint16 { return 0x0F8D }

/*
func BPCount(){
    FindTypeEx(BP(), 0, Backpack(), True)
}
    return FindFullQuantity()
func BMCount(){
    FindTypeEx(BM(), 0, Backpack(), True)
}
    return FindFullQuantity()
func GACount(){
    FindTypeEx(GA(), 0, Backpack(), True)
}
    return FindFullQuantity()
func GSCount(){
    FindTypeEx(GS(), 0, Backpack(), True)
}
    return FindFullQuantity()
func MRCount(){
    FindTypeEx(MR(), 0, Backpack(), True)
}
    return FindFullQuantity()
func NSCount(){
    FindTypeEx(NS(), 0, Backpack(), True)
}
    return FindFullQuantity()
func SACount(){
    FindTypeEx(SA(), 0, Backpack(), True)
}
    return FindFullQuantity()
func SSCount(){
    FindTypeEx(SS(), 0, Backpack(), True)
}
    return FindFullQuantity()
_auto_buy = _ScriptMethod(240)  # AutoBuy
_auto_buy.argtypes = [_ushort,  # ItemType
                      _ushort,  # ItemColor
                      _ushort]  # Quantity
func AutoBuy(ItemType, ItemColor, Quantity){
    _auto_buy(ItemType, ItemColor, Quantity)
}
_get_shop_list = _ScriptMethod(241)  # GetShopList
_get_shop_list.restype = _buffer
func GetShopList(){
    result = []
}
    data = _get_shop_list()
    count = _uint.from_buffer(data)
    offset = count.size
    while 42:
        if offset >= len(data) - 1:
            break
        string = _str.from_buffer(data, offset)
        offset += string.size
        result.append(string.value)
    return result
*/
func ClearShopList() {
	p := NewVoidPacket(242)
	p.send(senderFunc)
}

/*
_auto_buy_extended = _ScriptMethod(243)  # AutoBuyEx
_auto_buy_extended.argtypes = [_ushort,  # ItemType
                               _ushort,  # ItemColor
                               _ushort,  # Quantity
                               _uint,  # Price
                               _str]  # ItemName
func AutoBuyEx(ItemType, ItemColor, Quantity, Price, ItemName){
    _auto_buy_extended(ItemType, ItemColor, Quantity, Price, ItemName)
}
*/
func GetAutoBuyDelay() <-chan uint16 {
	p := NewUint16Packet(244)
	p.send(senderFunc)
	return p.out
}

func SetAutoBuyDelay(Value uint16) {
	p := NewVoidPacket(245, Value)
	p.send(senderFunc)
}

func GetAutoSellDelay() <-chan uint16 {
	p := NewUint16Packet(246)
	p.send(senderFunc)
	return p.out
}

func SetAutoSellDelay(Value uint16) {
	p := NewVoidPacket(247, Value)
	p.send(senderFunc)
}

func AutoSell(itemType, itemColor, quantity uint16) {
	p := NewVoidPacket(SCAutoSell, itemType, itemColor, quantity)
	p.send(senderFunc)
}
func RequestStats(ObjID uint32) {
	p := NewVoidPacket(249, ObjID)
	p.send(senderFunc)
}

func HelpRequest() {
	p := NewVoidPacket(250)
	p.send(senderFunc)
}

func QuestRequest() {
	p := NewVoidPacket(251)
	p.send(senderFunc)
}

func RenameMobile(mobId uint32, newName string) {
	p := NewVoidPacket(SCRenameMobile, mobId, newName)
	p.send(senderFunc)
}

func MobileCanBeRenamed(Mob_ID uint32) <-chan bool {
	p := NewBoolPacket(253, Mob_ID)
	p.send(senderFunc)
	return p.out
}

/*
_lock_stat = _ScriptMethod(254)  # ChangeStatLockState
_lock_stat.argtypes = [_ubyte,  # statNum
                       _ubyte]  # statState
func SetStatState(statNum, statState){
    _lock_stat(statNum, statState)
}
_get_static_art_bitmap = _ScriptMethod(255)  # GetStaticArtBitmap
_get_static_art_bitmap.restype = _buffer  # Bitmap file in bytes
_get_static_art_bitmap.argtypes = [_uint,  # Id
                                   _ushort]  # Hue
func GetStaticArtBitmap(Id, Hue){
    p :=
p.send(senderFunc)
// return _get_static_art_bitmap(Id, Hue)
}
_print_script_methods = _ScriptMethod(256)  # PrintScriptMethodsList
_print_script_methods.argtypes = [_str,  # FileName
                                  _bool]  # SortedList
func PrintScriptMethodsList(FileName, SortedList){
    _print_script_methods(FileName, SortedList)
}
*/
func Alarm() {
	p := NewVoidPacket(257)
	p.send(senderFunc)
}

func UOSay(text string) {
	p := NewVoidPacket(308, text)
	p.send(senderFunc)
}

/*
_uo_say_color = _ScriptMethod(309)  # SendTextToUOColor
_uo_say_color.argtypes = [_str,  # Text
                          _ushort]  # Color
func UOSayColor(Text, Color){
    _uo_say_color(Text, Color)
}
_reg_stealth = 0, '0', 'reg_stealth', 'stealth'
_reg_char = 1, '1', 'reg_char', 'char'_set_global = _ScriptMethod(310)  # SetGlobal
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
p.send(senderFunc)
// return _get_global(region[0], VarName)
    else:
        raise ValueError('GlobalRegion must be "stealth" or "char".')
_console_entry_reply = _ScriptMethod(312)
_console_entry_reply.argtypes = [_str]  # Text
func ConsoleEntryReply(Text){
    _console_entry_reply(Text)
}
*/
func ConsoleEntryUnicodeReply(text string) {
	p := NewVoidPacket(313, text)
	p.send(senderFunc)
}

func GameServerIPString() <-chan string {
	p := NewStringPacket(341)
	p.send(senderFunc)
	return p.out
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
        type_, data = winreg.QueryValueEx(easyuo_key, '*' + str(num))
    return data
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
    res = (res - 7) ^ 0x0045
    return 0 if res > 0xFFFF else res
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
        multi *= 26
    return (res - 7) ^ 0x0045
*/

func InviteToParty(ID uint32) {
	p := NewVoidPacket(262, ID)
	p.send(senderFunc)
}

func RemoveFromParty(ID uint32) {
	p := NewVoidPacket(263, ID)
	p.send(senderFunc)
}

func PartyMessageTo(oid uint32, msg string) {
	p := NewVoidPacket(SCPartyMessageTo, msg)
	p.send(senderFunc)
}

func PartySay(msg string) {
	p := NewVoidPacket(265, msg)
	p.send(senderFunc)
}

func PartyCanLootMe(value bool) {
	p := NewVoidPacket(SCPartyCanLootMe, value)
	p.send(senderFunc)
}

func PartyAcceptInvite() {
	p := NewVoidPacket(267)
	p.send(senderFunc)
}

func PartyDeclineInvite() {
	p := NewVoidPacket(268)
	p.send(senderFunc)
}

func PartyLeave() {
	p := NewVoidPacket(269)
	p.send(senderFunc)
}

func InParty() <-chan bool {
	p := NewBoolPacket(271)
	p.send(senderFunc)
	return p.out
}

/*
_get_party_members = _ScriptMethod(270)  # PartyMembersList
_get_party_members.restype = _buffer  # Array of Cardinal
func PartyMembersList(){
    result = []
}
    data = _get_party_members()
    count = _uint.from_buffer(data)
    if count:
        fmt = '<' + count * 'I'
        result.extend(_struct.unpack(fmt, data))
    return result
*/
func ICQConnected() <-chan bool {
	p := NewBoolPacket(272)
	p.send(senderFunc)
	return p.out
}

/*
_icq_connect = _ScriptMethod(273)  # ICQ_Connect
_icq_connect.argtypes = [_uint,  # UIN
                         _str]  # Password
func ICQConnect(UIN, Password){
    _icq_connect(UIN, Password)
}
*/
func ICQDisconnect() {
	p := NewVoidPacket(274)
	p.send(senderFunc)
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
p.send(senderFunc)
// return _messenger_get_connected(_messengers[MesID])
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
p.send(senderFunc)
// return _messenger_get_token(_messengers[MesID])
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
p.send(senderFunc)
// return _messenger_get_name(_messengers[MesID])
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
                1: 1, 'tfStatic': 1, 'tfstatic': 1, 'Static': 1, 'static': 1}_get_tile_flags = _ScriptMethod(278)  # GetTileFlags
_get_tile_flags.restype = _uint
_get_tile_flags.argtypes = [_ubyte,  # TileGroup
                            _ushort]  # Tile
func GetTileFlags(TileGroup, Tile){
    if TileGroup not in _tile_groups.keys():
}
        raise ValueError('GetTileFlags: TileGroup must be "Land" or "Static"')
    group = _tile_groups[TileGroup]
    p :=
p.send(senderFunc)
// return _get_tile_flags(group, Tile)
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
        result.append(string.value)
    return result
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
            result['Name'] = result['Name'].decode()
    return result
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
            result['Name'] = result['Name'].decode()
    return result
_get_cell = _ScriptMethod(13)  # GetCell
_get_cell.restype = _buffer  # TMapCell
_get_cell.argtypes = [_ushort,  # X
                      _ushort,  # Y
                      _ubyte]  # WorldNum
func GetCell(X, Y, WorldNum){
    result = {}
}
    data = _get_cell(X, Y, WorldNum)
    if data:
        fmt = '<Hb'
        keys = 'Tile', 'Z'
        values = _struct.unpack(fmt, data)
        result.update(zip(keys, values))
    return result
_get_layer_count = _ScriptMethod(282)  # GetLayerCount
_get_layer_count.restype = _ubyte
_get_layer_count.argtypes = [_ushort,  # X
                             _ushort,  # Y
                             _ubyte]  # WorldNum
func GetLayerCount(X, Y, WorldNum){
    p :=
p.send(senderFunc)
// return _get_layer_count(X, Y, WorldNum)
}
*/

func ReadStaticsXY(x, y uint16, wnum byte) <-chan []StaticsXY {
	p := NewReadStaticsXYPacket(x, y, wnum)
	p.send(senderFunc)
	return p.out
}

func GetSurfaceZ(x, y uint16, worldNum byte) <-chan byte {
	p := NewBytePacket(SCGetSurfaceZ, x, y, worldNum)
	p.send(senderFunc)
	return p.out
}

/*
_is_cell_passable = _ScriptMethod(285)  # IsWorldCellPassable
_is_cell_passable.restype = _buffer  # Boolean, ShortInt  4 bytes
_is_cell_passable.argtypes = [_ushort,  # CurrX
                              _ushort,  # CurrY
                              _byte,  # CurrZ
                              _ushort,  # DestX
                              _ushort,  # DestY
                              _ubyte]  # WorldNum
*/
func IsWorldCellPassable(currX, currY uint16, currZ int8, destX, destY uint16, worldNum byte) <-chan WorldCellPassable {
	p := NewIsWorldCellPassablePacket(currX, currY, currZ, destX, destY, worldNum)
	p.send(senderFunc)
	return p.out
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
func GetStaticTilesArray(Xmin, Ymin, Xmax, Ymax, WorldNum, TileTypes){
    if not _iterable(TileTypes):
}
        TileTypes = [TileTypes]
    result = []
    data = _get_statics_array(
        Xmin, Ymin, Xmax, Ymax, WorldNum, len(TileTypes),
        _struct.pack('<' + 'H' * len(TileTypes), *TileTypes)
    )
    count = _uint.from_buffer(data)
    fmt = '<3Hb'
    size = _struct.calcsize(fmt)
    for i in range(count):
        result.append(_struct.unpack_from(fmt, data, count.size + i * size))
    return result
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
        result.append(_struct.unpack_from(fmt, data, count.size + i * size))
    return result
*/
func ClientPrint(text string) {
	p := NewVoidPacket(289, text)
	p.send(senderFunc)
}

func ClientPrintEx(senderID uint32, color, font uint16, text string) {
	p := NewVoidPacket(SCClientPrintEx, senderID, color, font, text)
	p.send(senderFunc)
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
	p := NewVoidPacket(292)
	p.send(senderFunc)
}

func ClientRequestTileTarget() {
	p := NewVoidPacket(293)
	p.send(senderFunc)
}

func ClientTargetResponsePresent() <-chan bool {
	p := NewBoolPacket(294)
	p.send(senderFunc)
	return p.out
}

/*
_client_target_response = _ScriptMethod(295)  # ClientTargetResponse
_client_target_response.restype = _buffer  # TTargetInfo
func ClientTargetResponse(){
    result = {}
}
    data = _client_target_response()
    if data:
        fmt = '<I3Hb'
        keys = 'ID', 'Tile', 'X', 'Y', 'Z'
        values = _struct.unpack(fmt, data)
        result.update(zip(keys, values))
    return result
func WaitForClientTargetResponse(MaxWaitTimeMS){
    end = _time.time() + MaxWaitTimeMS / 1000
}
    while _time.time() < end:
        if ClientTargetResponsePresent():
            return True
        Wait(10)
    return False
*/

func CheckLag(timeoutMS uint32) <-chan bool {
	p := NewBoolPacket(299, timeoutMS)
	p.send(senderFunc)
	return p.out
}

func GetQuestArrow() <-chan Point2D {
	p := NewPoint2DPacket(SCGetQuestArrow)
	p.send(senderFunc)
	return p.out
}

func GetSilentMode() <-chan bool {
	p := NewBoolPacket(302)
	p.send(senderFunc)
	return p.out
}

func ClearInfoWindow() {
	p := NewVoidPacket(348)
	p.send(senderFunc)
}

func SetSilentMode(value bool) {
	p := NewVoidPacket(SCSetSilentMode, value)
	p.send(senderFunc)
}

func FillNewWindow(s string) {
	p := NewVoidPacket(303, s)
	p.send(senderFunc)
}

func StealthPath() <-chan string {
	p := NewStringPacket(305)
	p.send(senderFunc)
	return p.out
}

func CurrentScriptPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

func GetStealthProfilePath() <-chan string {
	p := NewStringPacket(306)
	p.send(senderFunc)
	return p.out
}

func GetShardPath() <-chan string {
	p := NewStringPacket(307)
	p.send(senderFunc)
	return p.out
}

func Step(direction byte, running bool) <-chan byte {
	p := NewBytePacket(SCStep, direction, running)
	p.send(senderFunc)
	return p.out
}

func StepQ(direction byte, running bool) <-chan int32 {
	p := NewIntPacket(SCStepQ, direction, running)
	p.send(senderFunc)
	return p.out
}

func MoveXYZ(xdst, ydst uint16, zdst int8, optimized bool, accuracyXY, accurancyZ int32, running bool) <-chan bool {
	p := NewBoolPacket(SCMoveXYZ, xdst, ydst, zdst, accuracyXY, accurancyZ, running)
	p.send(senderFunc)
	return p.out
}

func MoveXY(xdst, ydst uint16, optimized bool, accuracy int32, running bool) <-chan bool {
	p := NewBoolPacket(SCMoveXY, xdst, ydst, optimized, accuracy, running)
	p.send(senderFunc)
	return p.out
}

func SetBadLocation(x, y uint16) {
	p := NewVoidPacket(SCSetBadLocation, x, y)
	p.send(senderFunc)
}

func SetGoodLocation(x, y uint16) {
	p := NewVoidPacket(SCSetGoodLocation, x, y)
	p.send(senderFunc)
}

func ClearBadLocationList() {
	p := NewVoidPacket(330)
	p.send(senderFunc)
}

func SetBadObject(otype, color uint16, radius byte) {
	p := NewVoidPacket(SCSetBadObject, otype, color, radius)
	p.send(senderFunc)
}

func ClearBadObjectList() {
	p := NewVoidPacket(332)
	p.send(senderFunc)
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
p.send(senderFunc)
// return _check_los(xf, yf, zf, xt, yt, zt, WorldNum, LOSCheckType, options)
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
        result.append(_struct.unpack_from(fmt, data, count.size + i * size))
    return result
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
        result.append(_struct.unpack_from(fmt, data, count.size + i * size))
    return result
func Dist(x1, y1, x2, y2){
    dx = abs(x2 - x1)
}
    dy = abs(y2 - y1)
    return dx if dx > dy else dy
func CalcCoord(x, y, Dir){
    if Dir > 7:
}
        return x, y
    dirs = {0: (0, -1),
            1: (1, -1),
            2: (1, 0),
            3: (1, 1),
            4: (0, 1),
            5: (-1, 1),
            6: (-1, 0),
            7: (-1, -1)}
    dx, dy = dirs[Dir]
    return x + dx, y + dy
func CalcDir(Xfrom, Yfrom, Xto, Yto){
    dx = abs(Xto - Xfrom)
}
    dy = abs(Yto - Yfrom)
    if dx == dy == 0:
        return 100
    elif (dx / (dy + 0.1)) >= 2:
        return 6 if Xfrom > Xto else 2
    elif (dy / (dx + 0.1)) >= 2:
        return 0 if Yfrom > Yto else 4
    elif Xfrom > Xto:
        return 7 if Yfrom > Yto else 5
    elif Xfrom < Xto:
        return 1 if Yfrom > Yto else 3
*/
func SetRunUnmountTimer(Value uint16) {
	p := NewVoidPacket(316, Value)
	p.send(senderFunc)
}

func SetWalkMountTimer(Value uint16) {
	p := NewVoidPacket(317, Value)
	p.send(senderFunc)
}

func SetRunMountTimer(Value uint16) {
	p := NewVoidPacket(318, Value)
	p.send(senderFunc)
}

func SetWalkUnmountTimer(Value uint16) {
	p := NewVoidPacket(319, Value)
	p.send(senderFunc)
}

func GetRunMountTimer() <-chan uint16 {
	p := NewUint16Packet(320)
	p.send(senderFunc)
	return p.out
}

func GetWalkMountTimer() <-chan uint16 {
	p := NewUint16Packet(321)
	p.send(senderFunc)
	return p.out
}

func GetRunUnmountTimer() <-chan uint16 {
	p := NewUint16Packet(322)
	p.send(senderFunc)
	return p.out
}

func GetWalkUnmountTimer() <-chan uint16 {
	p := NewUint16Packet(323)
	p.send(senderFunc)
	return p.out
}

func GetLastStepQUsedDoor() <-chan uint32 {
	p := NewUint32Packet(344)
	p.send(senderFunc)
	return p.out
}

func StopMover() {
	p := NewVoidPacket(353)
	p.send(senderFunc)
}

func MoverStop() {
	StopMover()
}

/*
_set_reconnector_ext = _ScriptMethod(354)  # SetARExtParams
_set_reconnector_ext.argtypes = [_str,  # ShardName
                                 _str,  # CharName
                                 _bool]  # UseAtEveryConnect
func SetARExtParams(ShardName, CharName, UseAtEveryConnect){
    _set_reconnector_ext(ShardName, CharName, UseAtEveryConnect)
}
_use_item_on_mobile = _ScriptMethod(359)  # SCUseItemOnMobile
_use_item_on_mobile.argtypes = [_uint,  # ItemSerial
                                _uint]  # TargetSerial
func UseItemOnMobile(ItemSerial, TargetSerial){
    _use_item_on_mobile(ItemSerial, TargetSerial)
}
*/
func BandageSelf() {
	p := NewVoidPacket(360)
	p.send(senderFunc)
}

func GlobalChatJoinChannel(chName string) {
	p := NewVoidPacket(361, chName)
	p.send(senderFunc)
}

func GlobalChatLeaveChannel() {
	p := NewVoidPacket(362)
	p.send(senderFunc)
}

func GlobalChatSendMsg(msgText string) {
	p := NewVoidPacket(363, msgText)
	p.send(senderFunc)
}

func GlobalChatActiveChannel() <-chan string {
	p := NewStringPacket(364)
	p.send(senderFunc)
	return p.out
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
        result.append(string.value)
    return result
_set_open_doors = _ScriptMethod(400)  # SetMoveOpenDoor
_set_open_doors.argtypes = [_bool]  # Value
func SetMoveOpenDoor(Value){
    _set_open_doors(Value)
}
*/
func GetMoveOpenDoor() <-chan bool {
	p := NewBoolPacket(401)
	p.send(senderFunc)
	return p.out
}

func SetMoveThroughNPC(Value uint16) {
	p := NewVoidPacket(402, Value)
	p.send(senderFunc)
}

func GetMoveThroughNPC() <-chan uint16 {
	p := NewUint16Packet(403)
	p.send(senderFunc)
	return p.out
}

func SetMoveThroughCorner(value bool) {
	p := NewVoidPacket(SCGetMoveThroughCorner, value)
	p.send(senderFunc)
}

func GetMoveThroughCorner() <-chan bool {
	p := NewBoolPacket(405)
	p.send(senderFunc)
	return p.out
}

func SetMoveHeuristicMult(value int32) {
	p := NewVoidPacket(SCSetMoveHeuristicMult, value)
	p.send(senderFunc)
}

func GetMoveHeuristicMult() <-chan uint32 {
	p := NewUint32Packet(407)
	p.send(senderFunc)
	return p.out
}

func SetMoveCheckStamina(Value uint16) {
	p := NewVoidPacket(408, Value)
	p.send(senderFunc)
}

func GetMoveCheckStamina() <-chan uint16 {
	p := NewUint16Packet(409)
	p.send(senderFunc)
	return p.out
}

func SetMoveTurnCost(value uint32) {
	p := NewVoidPacket(SCSetMoveTurnCost, value)
	p.send(senderFunc)
}

func GetMoveTurnCost() <-chan uint32 {
	p := NewUint32Packet(411)
	p.send(senderFunc)
	return p.out
}

func SetMoveBetweenTwoCorners(value bool) {
	p := NewVoidPacket(SCSetMoveBetweenTwoCorners, value)
	p.send(senderFunc)
}

func GetMoveBetweenTwoCorners() <-chan bool {
	p := NewBoolPacket(413)
	p.send(senderFunc)
	return p.out
}

func GetMultis() <-chan []Multi {
	p := NewGetMultisPacket()
	p.send(senderFunc)
	return p.out
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
p.send(senderFunc)
// return '{ ' + template.format(hex(self.model), hex(self.color),
                                          self.text) + ' }'        func __repr__(self){
            return self.__str__()
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
        item.text = text.value
        result.append(item)    return result
*/
func CloseClientGump(ID uint32) {
	p := NewVoidPacket(342, ID)
	p.send(senderFunc)
}

/*
_get_next_step_z = _ScriptMethod(366)  # GetNextStepZ
_get_next_step_z.restype = _byte
_get_next_step_z.argtypes = [_ushort,  # CurrX
                             _ushort,  # CurrY
                             _ushort,  # DestX
                             _ushort,  # DestY
                             _ubyte,  # WorldNum
                             _byte]  # CurrZ
func GetNextStepZ(CurrX, CurrY, DestX, DestY, WorldNum, CurrZ){
    p :=
p.send(senderFunc)
// return _get_next_step_z(CurrX, CurrY, DestX, DestY, WorldNum, CurrZ)
}
*/
func ClientHide(ID uint32) <-chan bool {
	p := NewBoolPacket(368, ID)
	p.send(senderFunc)
	return p.out
}

/*
_get_skill_lock_state = _ScriptMethod(369)  # GetSkillLockState
_get_skill_lock_state.restype = _byte
_get_skill_lock_state.argtypes = [_str]  # SkillName
func GetSkillLockState(SkillName){
    p :=
p.send(senderFunc)
// return _get_skill_lock_state(SkillName)
}
_get_stat_lock_state = _ScriptMethod(372)  # GetStatLockState
_get_stat_lock_state.restype = _byte
_get_stat_lock_state.argtypes = [_str]  # SkillName
func GetStatLockState(SkillName){
    _get_stat_lock_state(SkillName)
}
_equip_last_weapon = _ScriptMethod(370)
func EquipLastWeapon(){
    _equip_last_weapon()
}
# Book functions_book_get_page_text = _ScriptMethod(373)
_book_get_page_text.restype = _str
_book_get_page_text.argtypes = [_ushort]
func BookGetPageText(Page){
    p :=
p.send(senderFunc)
// return _book_get_page_text(Page)
}
_book_set_text = _ScriptMethod(374)
_book_set_text.argtypes = [_str]
func BookSetText(Text){
    _book_set_text(Text)
}
_book_set_page_text = _ScriptMethod(375)
_book_set_page_text.argtypes = [_ushort, _str]
func BookSetPageText(Page, Text){
    _book_set_page_text(Page, Text)
}
_book_clear_text = _ScriptMethod(376)
func BookClearText(){
    _book_clear_text()
}
_book_set_header = _ScriptMethod(377)
_book_set_header.argtypes = [_str, _str]
func BookSetHeader(title, author){
    _book_set_header(title, author)
}
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
	p := NewUint16Packet(SCGetScriptsCount)
	p.send(senderFunc)
	return p.out
}

func GetScriptPath(scriptIndex uint16) <-chan string {
	p := NewStringPacket(SCGetScriptPath, scriptIndex)
	p.send(senderFunc)
	return p.out
}

func GetScriptName(ScriptIndex uint16) <-chan string {
	p := NewStringPacket(SCGetScriptName, ScriptIndex)
	p.send(senderFunc)
	return p.out
}

func GetScriptState(scriptIndex uint16) <-chan int8 {
	p := NewInt8Packet(SCGetScriptState, scriptIndex)
	p.send(senderFunc)
	return p.out

}

func StartScript(scriptPath string) <-chan uint16 {
	p := NewUint16Packet(SCStartScript, scriptPath)
	p.send(senderFunc)
	return p.out
}

func StopScript(scriptIndex uint16) {
	p := NewVoidPacket(456)
	p.send(senderFunc)
}

func PauseResumeScript(scriptIndex uint16) {
	p := NewVoidPacket(456)
	p.send(senderFunc)
}

func StopAllScripts() {
	p := NewVoidPacket(457)
	p.send(senderFunc)

}
