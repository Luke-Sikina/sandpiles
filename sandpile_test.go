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
			t.Errorf("Error for isValidGridPosition: for coordinate %d %d, in grid of size %dx%d, expected %b", scenario.xCoord, scenario.yCoord, scenario.xSize, scenario.ySize, scenario.expected)
		}
	}
}

var getValidNeighborsTestData = []struct {
	xMax, yMax        int
	startingPoint     Coordinate
	expectedNeighbors []Coordinate
}{
	{10, 10, Coordinate{1, 1}, []Coordinate{Coordinate{1, 2}, Coordinate{1, 0}, Coordinate{0, 1}, Coordinate{2, 1}}},
	{10, 10, Coordinate{0, 1}, []Coordinate{Coordinate{0, 2}, Coordinate{0, 0}, Coordinate{1, 1}}},
	{10, 10, Coordinate{0, 0}, []Coordinate{Coordinate{0, 1}, Coordinate{1, 0}}},
	{10, 10, Coordinate{9, 9}, []Coordinate{Coordinate{9, 8}, Coordinate{8, 9}}},
}

func TestGetValidNeighbors(t *testing.T) {
	for scenario, data := range getValidNeighborsTestData {
		neighbors := data.startingPoint.getValidNeighbors(data.xMax, data.yMax)
		if len(neighbors) != len(data.expectedNeighbors) {
			t.Error("Too few/many neighbors")
		} else {
			for index, coord := range data.expectedNeighbors {
				if coord.x != neighbors[index].x {
					t.Errorf("Scenario %d: Wrong x value for coord number %d. Expected %d got %d", scenario, index, coord.x, neighbors[index].x)
				}
				if coord.y != neighbors[index].y {
					t.Errorf("Scenario %d: Wrong y value for coord number %d. Expected %d got %d", scenario, index, coord.y, neighbors[index].y)
				}
			}
		}
	}
}
