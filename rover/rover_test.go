package rover

import "testing"

func TestNew(t *testing.T) {
	rover := New()

	if rover == nil {
		t.Errorf("expected rover, got nil")
	}
}

func TestNewInitialPosition(t *testing.T) {
	rover := New()

	if rover.Pos.X != 0 || rover.Pos.Y != 0 {
		t.Errorf("expected position (0,0), got (%d, %d)", rover.Pos.X, rover.Pos.Y)
	}
}
