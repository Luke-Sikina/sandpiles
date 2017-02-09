package main

import (
	"testing"
)

var gridPositionTestData = []struct {
	xCoord, yCoord, xSize, ySize int
	expected                     bool
}{
	{1, 1, 10, 10, true},
	{0, 0, 10, 10, true},
	{9, 9, 10, 10, true},
	{-1, 0, 10, 10, false},
	{0, -1, 10, 10, false},
	{-1, -1, 10, 10, false},
	{10, 0, 10, 10, false},
	{10, 10, 10, 10, false},
}

func TestIsValidGridPosition(t *testing.T) {
	for _, scenario := range gridPositionTestData {
		coordinate := Coordinate{scenario.xCoord, scenario.yCoord}
		actual := coordinate.isValidGridPosition(scenario.xSize, scenario.ySize)
		if actual != scenario.expected {
			t.Errorf("Error for isValidGridPosition: for coordinate %d $d, in grid of size %dx%d, expected %b", scenario.xCoord, scenario.yCoord, scenario.xSize, scenario.ySize, scenario.expected)
		}

	}
}
