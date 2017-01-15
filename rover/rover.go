package rover

import "github.com/andoco/plutorover/terrain"

// R is a struct for a commandable rover vehicle.
type R struct {
	Pos     terrain.Pos
	Heading terrain.Cardinal
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
	case terrain.CardinalNorth:
		y = 1
	case terrain.CardinalEast:
		x = 1
	case terrain.CardinalSouth:
		y = -1
	case terrain.CardinalWest:
		x = -1
	}

	r.Pos = terrain.Pos{r.Pos.X + x, r.Pos.Y + y}
}

func (r *R) backward() {
	x := 0
	y := 0

	switch r.Heading {
	case terrain.CardinalNorth:
		y = -1
	case terrain.CardinalEast:
		x = -1
	case terrain.CardinalSouth:
		y = 1
	case terrain.CardinalWest:
		x = 1
	}

	r.Pos = terrain.Pos{r.Pos.X + x, r.Pos.Y + y}
}

func (r *R) turnLeft() {
	r.Heading = terrain.RotateLeft(r.Heading)
}

func (r *R) turnRight() {
	r.Heading = terrain.RotateRight(r.Heading)
}

// New creates and returns a new rover vehicle.
func New() *R {
	return &R{}
}
