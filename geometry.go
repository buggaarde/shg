package shg

// A Point is an x/y coordinate
type Point struct {
	X, Y float64
}

// A Segment is a beginning and an end point
type Segment struct {
	Start, End Point
}

// The Index hash
type Index struct {
	X, Y int
}

// A Cell is a grid cell in the hash grid and contains entities.
type Cell struct {
	// emulating a hash set by using empty structs as values
	entities map[Entity]struct{}
}
