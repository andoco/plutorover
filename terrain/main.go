package terrain

// Pos is a struct containing the X and Y components of a 2D cartesian coordinate.
type Pos struct {
	X int
	Y int
}

// Cardinal is a type for a cardinal compass point.
type Cardinal int

const (
	CardinalNorth Cardinal = iota
	CardinalEast
	CardinalSouth
	CardinalWest
)
