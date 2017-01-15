package rover

import (
	"testing"

	"github.com/andoco/plutorover/grid"
)

func TestNew(t *testing.T) {
	g := &grid.G{5, 5}
	rover := New(g)

	if rover == nil {
		t.Fatal("expected rover, got nil")
	}

	if rover.Grid != g {
		t.Errorf("expected %v, got %v", g, rover.Grid)
	}

	if rover.Pos.X != 0 || rover.Pos.Y != 0 {
		t.Errorf("expected position (0,0), got (%d, %d)", rover.Pos.X, rover.Pos.Y)
	}

	if rover.Heading != grid.CardinalNorth {
		t.Errorf("expected heading %d, got %d", grid.CardinalNorth, rover.Heading)
	}
}

func TestMovement(t *testing.T) {
	testCases := []struct {
		startPos     grid.Pos
		startHeading grid.Cardinal
		endPos       grid.Pos
		endHeading   grid.Cardinal
		cmd          string
		name         string
	}{
		{
			grid.Pos{0, 0},
			grid.CardinalNorth,
			grid.Pos{0, 1},
			grid.CardinalNorth,
			"F",
			"Forward from origin",
		}, {
			grid.Pos{0, 0},
			grid.CardinalNorth,
			grid.Pos{0, 2},
			grid.CardinalNorth,
			"FF",
			"Forward two times",
		}, {
			grid.Pos{0, 0},
			grid.CardinalNorth,
			grid.Pos{0, -1},
			grid.CardinalNorth,
			"B",
			"Backward from origin",
		}, {
			grid.Pos{0, 0},
			grid.CardinalNorth,
			grid.Pos{0, 0},
			grid.CardinalWest,
			"L",
			"Rotate left from north",
		}, {
			grid.Pos{0, 0},
			grid.CardinalNorth,
			grid.Pos{0, 0},
			grid.CardinalEast,
			"R",
			"Rotate right from north",
		}, {
			grid.Pos{0, 0},
			grid.CardinalNorth,
			grid.Pos{-1, 1},
			grid.CardinalNorth,
			"LFRF",
			"Rotating and moving forward",
		},
		{
			grid.Pos{0, 0},
			grid.CardinalNorth,
			grid.Pos{1, -1},
			grid.CardinalNorth,
			"LBRB",
			"Rotating and moving backward",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			grid := &grid.G{5, 5}
			rover := New(grid)
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
