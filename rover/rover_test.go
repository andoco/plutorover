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

func TestMovement(t *testing.T) {
	testCases := []struct {
		startPos     terrain.Pos
		startHeading terrain.Cardinal
		endPos       terrain.Pos
		endHeading   terrain.Cardinal
		cmd          string
		name         string
	}{
		{
			terrain.Pos{0, 0},
			terrain.CardinalNorth,
			terrain.Pos{0, 1},
			terrain.CardinalNorth,
			"F",
			"Forward from origin",
		}, {
			terrain.Pos{0, 0},
			terrain.CardinalNorth,
			terrain.Pos{0, -1},
			terrain.CardinalNorth,
			"B",
			"Backward from origin",
		}, {
			terrain.Pos{0, 0},
			terrain.CardinalNorth,
			terrain.Pos{0, 0},
			terrain.CardinalWest,
			"L",
			"Rotate left from north",
		}, {
			terrain.Pos{0, 0},
			terrain.CardinalNorth,
			terrain.Pos{0, 0},
			terrain.CardinalEast,
			"R",
			"Rotate right from north",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rover := New()
			rover.Pos = tc.startPos

			rover.Process(tc.cmd)

			if rover.Pos != tc.endPos {
				t.Errorf("expected position %v, got %v", tc.endPos, rover.Pos)
			}

			if rover.Heading != tc.endHeading {
				t.Errorf("expected heading %v, got %v", tc.endHeading, rover.Heading)
			}
		})
	}
}
