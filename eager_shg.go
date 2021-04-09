package shg

type EagerSpacialHashGrid struct {
	grid     [][]Cell
	cellSize float64
}

func NewEagerSpacialHashGrid(dimx, dimy int, cellSize float64) *EagerSpacialHashGrid {
	grid := make([][]Cell, dimx)
	for i := 0; i < dimx; i++ {
		grid[i] = make([]Cell, dimy)
	}

	return &EagerSpacialHashGrid{
		grid:     grid,
		cellSize: cellSize,
	}
}

func (s *EagerSpacialHashGrid) Insert(ent Entity) {
	x, y := ent.Position()
	index := Index{
		X: int(x / s.cellSize),
		Y: int(y / s.cellSize),
	}

	s.grid[index.X][index.Y].entities[ent] = struct{}{}
	ent.SetIndex(index)
}

func (s *EagerSpacialHashGrid) Remove(ent Entity) {
	index := ent.LastIndex()
	delete(s.grid[index.X][index.Y].entities, ent)
}

func (s *EagerSpacialHashGrid) Update(ent Entity) {
	s.Remove(ent)
	s.Insert(ent)
}

func (s *EagerSpacialHashGrid) Cell(index Index) Cell {
	x, y := index.X, index.Y
	return s.grid[x][y]
}
