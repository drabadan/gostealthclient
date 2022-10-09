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
	MaxWeight          uint16 `bin:"len:4"`
	Race               byte   `bin:"len:1"`
	StatCap            uint16 `bin:"len:4"`
	PetsCurrent        byte   `bin:"len:1"`
	PetsMax            byte   `bin:"len:1"`
	FireResist         uint16 `bin:"len:4"`
	ColdResist         uint16 `bin:"len:4"`
	PoisonResist       uint16 `bin:"len:4"`
	EnergyResist       uint16 `bin:"len:4"`
	Luck               int16  `bin:"len:4"`
	DamageMin          uint16 `bin:"len:4"`
	DamageMax          uint16 `bin:"len:4"`
	TithingPoints      uint32 `bin:"len:8"`
	HitChanceIncr      uint16 `bin:"len:4"`
	SwingSpeedIncr     uint16 `bin:"len:4"`
	DamageChanceIncr   uint16 `bin:"len:4"`
	LowerReagentCost   uint16 `bin:"len:4"`
	HpRegen            uint16 `bin:"len:4"`
	StamRegen          uint16 `bin:"len:4"`
	ManaRegen          uint16 `bin:"len:4"`
	ReflectPhysDamage  uint16 `bin:"len:4"`
	EnhancePotions     uint16 `bin:"len:4"`
	DefenseChanceIncr  uint16 `bin:"len:4"`
	SpellDamageIncr    uint16 `bin:"len:4"`
	FasterCastRecovery uint16 `bin:"len:4"`
	FasterCasting      uint16 `bin:"len:4"`
	LowerManaCost      uint16 `bin:"len:4"`
	StrengthIncr       uint16 `bin:"len:4"`
	DextIncr           uint16 `bin:"len:4"`
	IntIncr            uint16 `bin:"len:4"`
	HpIncr             uint16 `bin:"len:4"`
	StamIncr           uint16 `bin:"len:4"`
	ManaIncr           uint16 `bin:"len:4"`
	MaxHpIncr          uint16 `bin:"len:4"`
	MaxStamIncr        uint16 `bin:"len:4"`
	MaxManaIncrease    uint16 `bin:"len:4"`
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
	ID   uint32     `bin:"len:8"`
	Tile uint16     `bin:"len:4"`
	X    Coordinate `bin:"len:4"`
	Y    Coordinate `bin:"len:4"`
	Z    int8       `bin:"len:2"`
}
