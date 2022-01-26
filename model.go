package gostealthclient

type StaticsXY struct {
	Tile, X, Y, Color uint16
	Z                 byte
}

type Multi struct {
	Id                                          uint32
	X, Y, XMax, XMin, YMax, YMin, Width, Height uint16
	Z                                           byte
}

type WorldCellPassable struct {
	Passable bool
	Z        int8
}
