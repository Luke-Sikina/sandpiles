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

func TestCreateGrid(t *testing.T) {
	grid := createGrid(8, 8, 4)
	for xCoord := 0; xCoord < 8; xCoord++ {
		for yCoord := 0; yCoord < 8; yCoord++ {
			if grid[xCoord][yCoord] != uint8(4) {
				t.Errorf("Wrong value at row %d col %d expected %d was %d", xCoord, yCoord, 4, grid[xCoord][yCoord])
			}
		}
	}
}

func TestCloneGrid(t *testing.T) {
	grid := createGrid(8, 8, 4)
	clone := grid.clone()
	for rowIndex, row := range grid {
		for cellIndex, cell := range row {
			if cell != clone[rowIndex][cellIndex] {
				t.Errorf("Wrong value in clone at row %d col %d expected %d was %d", rowIndex, cellIndex, cell, clone[rowIndex][cellIndex])
			}
		}
	}
}

var singleSiftTestingData = []struct {
	startingPoint Coordinate
	startingGrid  Grid
	resultGrid    Grid
}{
	{Coordinate{1, 1},
		Grid([][]uint8{
			{4, 4, 4, 4},
			{4, 4, 4, 4},
			{4, 4, 4, 4},
			{4, 4, 4, 4}}),
		Grid([][]uint8{
			{4, 5, 4, 4},
			{5, 0, 5, 4},
			{4, 5, 4, 4},
			{4, 4, 4, 4}})},
	{Coordinate{0, 0},
		Grid([][]uint8{
			{4, 4, 4, 4},
			{4, 4, 4, 4},
			{4, 4, 4, 4},
			{4, 4, 4, 4}}),
		Grid([][]uint8{
			{0, 5, 4, 4},
			{5, 4, 4, 4},
			{4, 4, 4, 4},
			{4, 4, 4, 4}})},
	{Coordinate{3, 1},
		Grid([][]uint8{
			{4, 4, 4, 4},
			{4, 4, 4, 4},
			{4, 4, 4, 4},
			{4, 4, 4, 4}}),
		Grid([][]uint8{
			{4, 4, 4, 4},
			{4, 4, 4, 4},
			{4, 5, 4, 4},
			{5, 0, 5, 4}})},
}

func TestSingleSift(t *testing.T) {
	for _, data := range singleSiftTestingData {
		reference := data.startingGrid.clone()
		data.startingPoint.sift(&reference, data.startingGrid)
		for x, row := range data.resultGrid {
			for y, cell := range row {
				if cell != data.startingGrid[x][y] {
					t.Errorf("Wrong value in grid at %d, %d, expected %d was %d", x, y, cell, data.startingGrid[x][y])
				}
			}
		}
	}
}
