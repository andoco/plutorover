package grid

import "fmt"

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

// G is a struct for a 2D grid.
type G struct {
	Width  int
	Height int
}

func RotateLeft(dir Cardinal) Cardinal {
	switch dir {
	case CardinalNorth:
		return CardinalWest
	case CardinalWest:
		return CardinalSouth
	case CardinalSouth:
		return CardinalEast
	case CardinalEast:
		return CardinalNorth
	}

	panic(fmt.Sprintf("cannot turn using cardinal value %v", dir))
}

func RotateRight(dir Cardinal) Cardinal {
	switch dir {
	case CardinalNorth:
		return CardinalEast
	case CardinalEast:
		return CardinalSouth
	case CardinalSouth:
		return CardinalWest
	case CardinalWest:
		return CardinalNorth
	}

	panic(fmt.Sprintf("cannot turn using cardinal value %v", dir))
}
