package grid

import "testing"

func TestTranslate(t *testing.T) {
	testCases := []struct {
		inPos  Pos
		deltaX int
		deltaY int
		outPos Pos
		name   string
	}{
		{Pos{0, 0}, 1, 0, Pos{1, 0}, "positive x"},
		{Pos{0, 0}, -1, 0, Pos{-1, 0}, "negative x"},
		{Pos{0, 0}, 0, 1, Pos{0, 1}, "positive y"},
		{Pos{0, 0}, 0, -1, Pos{0, -1}, "negative y"},
		{Pos{0, 0}, 1, 1, Pos{1, 1}, "x and y"},
		{Pos{1, 1}, 0, 0, Pos{1, 1}, "zero delta"},
	}

	g := &G{5, 5}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := g.Translate(tc.inPos, tc.deltaX, tc.deltaY)

			if p != tc.outPos {
				t.Errorf("expected %v, got %v", tc.outPos, p)
			}
		})
	}
}

func TestRotateLeft(t *testing.T) {
	testCases := []struct {
		inDir  Cardinal
		outDir Cardinal
		name   string
	}{
		{CardinalNorth, CardinalWest, "from north"},
		{CardinalWest, CardinalSouth, "from west"},
		{CardinalSouth, CardinalEast, "from south"},
		{CardinalEast, CardinalNorth, "from east"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dir := RotateLeft(tc.inDir)

			if dir != tc.outDir {
				t.Errorf("expected %v, got %v", tc.outDir, dir)
			}
		})
	}
}

func TestRotateRight(t *testing.T) {
	testCases := []struct {
		inDir  Cardinal
		outDir Cardinal
		name   string
	}{
		{CardinalNorth, CardinalEast, "from north"},
		{CardinalEast, CardinalSouth, "from east"},
		{CardinalSouth, CardinalWest, "from south"},
		{CardinalWest, CardinalNorth, "from west"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dir := RotateRight(tc.inDir)

			if dir != tc.outDir {
				t.Errorf("expected %v, got %v", tc.outDir, dir)
			}
		})
	}
}
