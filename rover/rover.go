package rover

import "github.com/andoco/plutorover/terrain"

// R is a struct for a commandable rover vehicle.
type R struct {
	Pos terrain.Pos
}

// New creates and returns a new rover vehicle.
func New() *R {
	return &R{}
}
