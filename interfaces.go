package shg

type Entity interface {
	Position() (float64, float64)
	BoundingBox() (minx, miny, maxx, maxy float64)
	Segments() []Segment
	SetIndex(Index)
	LastIndex() Index
}

type SHG interface {
	Cell(Index) Cell
	Insert(Entity)
	Remove(Entity)
	Update(Entity)
}
