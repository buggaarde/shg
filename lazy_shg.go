package shg

// The LazySpacialHashGrid does not have any initialized cells at
// construction time
type LazySpacialHashGrid struct {
	grid     map[Index]Cell
	cellSize float64
}

func NewLazySpacialHashGrid(cellSize float64) *LazySpacialHashGrid {
	return &LazySpacialHashGrid{
		grid:     make(map[Index]Cell),
		cellSize: cellSize,
	}
}

func (s *LazySpacialHashGrid) Insert(ent Entity) {
	x, y := ent.Position()
	index := Index{
		X: int(x / s.cellSize),
		Y: int(y / s.cellSize),
	}

	s.grid[index].entities[ent] = struct{}{}
	ent.SetIndex(index)
}

func (s *LazySpacialHashGrid) Remove(ent Entity) {
	delete(s.grid[ent.LastIndex()].entities, ent)
}

func (s *LazySpacialHashGrid) Update(ent Entity) {
	s.Remove(ent)
	s.Insert(ent)
}

func (s *LazySpacialHashGrid) Cell(index Index) Cell {
	return s.grid[index]
}
