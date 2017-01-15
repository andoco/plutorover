package rover

import "github.com/andoco/plutorover/grid"

// R is a struct for a commandable rover vehicle.
type R struct {
	Grid    *grid.G
	Pos     grid.Pos
	Heading grid.Cardinal
}

// Process will execute a command string on the rover.
func (r *R) Process(commands string) error {
	for _, cmd := range commands {
		r.interpret(string(cmd))
	}
	return nil
}

func (r *R) interpret(cmd string) {
	switch cmd {
	case "F":
		r.forward()
	case "B":
		r.backward()
	case "L":
		r.turnLeft()
	case "R":
		r.turnRight()
	}
}

func (r *R) forward() {
	x := 0
	y := 0

	switch r.Heading {
	case grid.CardinalNorth:
		y = 1
	case grid.CardinalEast:
		x = 1
	case grid.CardinalSouth:
		y = -1
	case grid.CardinalWest:
		x = -1
	}

	r.Pos = grid.Pos{r.Pos.X + x, r.Pos.Y + y}
}

func (r *R) backward() {
	x := 0
	y := 0

	switch r.Heading {
	case grid.CardinalNorth:
		y = -1
	case grid.CardinalEast:
		x = -1
	case grid.CardinalSouth:
		y = 1
	case grid.CardinalWest:
		x = 1
	}

	r.Pos = grid.Pos{r.Pos.X + x, r.Pos.Y + y}
}

func (r *R) turnLeft() {
	r.Heading = grid.RotateLeft(r.Heading)
}

func (r *R) turnRight() {
	r.Heading = grid.RotateRight(r.Heading)
}

// New creates and returns a new rover vehicle.
func New(grid *grid.G) *R {
	return &R{Grid: grid}
}
