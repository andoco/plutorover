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
		r.Pos = terrain.Pos{r.Pos.X, r.Pos.Y + 1}
	case "B":
		r.Pos = terrain.Pos{r.Pos.X, r.Pos.Y - 1}
	case "L":
		r.turnLeft()
	case "R":
		r.turnRight()
	}
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
