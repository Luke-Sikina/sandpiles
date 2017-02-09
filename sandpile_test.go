package main

import (
	"testing"
)

func TestIsValidGridPosition(t *testing.T) {
	coordinate := Coordinate{1, 1}
	actual := coordinate.isValidGridPosition(1, 1)
	if !actual {
		t.Error("Expected true for a coordinate of 1, 1 in a 10x10 grid")
	}
}
