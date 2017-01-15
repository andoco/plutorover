package rover

import (
	"testing"

	"github.com/andoco/plutorover/terrain"
)

func TestNew(t *testing.T) {
	rover := New()

	if rover == nil {
		t.Fatal("expected rover, got nil")
	}

	if rover.Pos.X != 0 || rover.Pos.Y != 0 {
		t.Errorf("expected position (0,0), got (%d, %d)", rover.Pos.X, rover.Pos.Y)
	}

	if rover.Heading != terrain.CardinalNorth {
		t.Errorf("expected heading %d, got %d", terrain.CardinalNorth, rover.Heading)
	}
}
