package coords

import (
	"testing"
)

func TestGetMapBlockCoordsFromPlain(t *testing.T) {
	mb := GetMapBlockCoordsFromPlain(0, 0, 0)
	if mb.X != 0 || mb.Y != 0 || mb.Z != 0 {
		t.Fatal("unexpected mapblock", mb)
	}

	mb = GetMapBlockCoordsFromPlain(1, 15, 0)
	if mb.X != 0 || mb.Y != 0 || mb.Z != 0 {
		t.Fatal("unexpected mapblock", mb)
	}

	mb = GetMapBlockCoordsFromPlain(31, 16, 15)
	if mb.X != 1 || mb.Y != 1 || mb.Z != 0 {
		t.Fatal("unexpected mapblock", mb)
	}

	mb = GetMapBlockCoordsFromPlain(-1, -16, -2)
	if mb.X != -1 || mb.Y != -1 || mb.Z != -1 {
		t.Fatal("unexpected mapblock", mb)
	}
}
