package terrain

import "testing"

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
