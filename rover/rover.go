package rover

import "github.com/andoco/plutorover/terrain"

// R is a struct for a commandable rover vehicle.
type R struct {
	Pos     terrain.Pos
	Heading terrain.Cardinal
}

// Process will execute a command string on the rover.
func (r *R) Process(cmd string) error {
	switch cmd {
	case "F":
		r.Pos = terrain.Pos{r.Pos.X, r.Pos.Y + 1}
	case "B":
		r.Pos = terrain.Pos{r.Pos.X, r.Pos.Y - 1}
	case "L":
		r.turnLeft()
	}
	return nil
}

func (r *R) turnLeft() {
	switch r.Heading {
	case terrain.CardinalNorth:
		r.Heading = terrain.CardinalWest
	}
}

// New creates and returns a new rover vehicle.
func New() *R {
	return &R{}
}
