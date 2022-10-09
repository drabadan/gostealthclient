package model

import "time"

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
