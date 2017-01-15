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

// Translate will move a point on the grid by amounts specified by x and y, while
// wrapping to remain within the grid.
func (g *G) Translate(p Pos, x, y int) Pos {
	p2 := Pos{p.X + x, p.Y + y}

	if p2.X >= g.Width {
		p2.X = p2.X - g.Width
	} else if p2.X < 0 {
		p2.X = g.Width + p2.X
	}

	if p2.Y >= g.Height {
		p2.Y = p2.Y - g.Height
	} else if p2.Y < 0 {
		p2.Y = g.Height + p2.Y
	}

	return p2
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
