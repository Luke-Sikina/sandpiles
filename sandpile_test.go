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
		actual := isValidGridPosition(scenario.xCoord, scenario.yCoord, scenario.xSize, scenario.ySize)
		if actual != scenario.expected {
			t.Errorf("Error for isValidGridPosition: for coordinate %d %d, in grid of size %dx%d, expected %b", scenario.xCoord, scenario.yCoord, scenario.xSize, scenario.ySize, scenario.expected)
		}
	}
}

var getValidNeighborsTestData = []struct {
	xMax, yMax, startingX, startingY       int
	expectedNeighborsX, expectedNeighborsY []int
}{
	{10, 10, 1, 1, []int{1, 1, 0, 2}, []int{2, 0, 1, 1}},
	{10, 10, 0, 1, []int{0, 0, 1}, []int{2, 0, 1}},
	{10, 10, 0, 0, []int{0, 1}, []int{1, 0}},
	{10, 10, 9, 9, []int{9, 8}, []int{8, 9}},
}

func TestGetValidNeighbors(t *testing.T) {
	for scenario, data := range getValidNeighborsTestData {
		neighborsX, neighborsY := getValidNeighbors(data.startingX, data.startingY, data.xMax, data.yMax)
		if len(neighborsX) != len(data.expectedNeighborsX) {
			t.Error("Too few/many neighbors")
		} else {
			for index, _ := range data.expectedNeighborsX {
				if data.expectedNeighborsX[index] != neighborsX[index] {
					t.Errorf("Scenario %d: Wrong x value for coord number %d. Expected %d got %d", scenario, index, data.expectedNeighborsX[index], neighborsX[index])
				}
				if data.expectedNeighborsY[index] != neighborsY[index] {
					t.Errorf("Scenario %d: Wrong y value for coord number %d. Expected %d got %d", scenario, index, data.expectedNeighborsY[index], neighborsY[index])
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
	startingX, startingY     int
	startingGrid, resultGrid Grid
}{
	{1, 1,
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
	{0, 0,
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
	{3, 1,
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
		singleSift(data.startingX, data.startingY, &reference, data.startingGrid)
		for x, row := range data.resultGrid {
			for y, cell := range row {
				if cell != data.startingGrid[x][y] {
					t.Errorf("Wrong value in grid at %d, %d, expected %d was %d", x, y, cell, data.startingGrid[x][y])
				}
			}
		}
	}
}

var fullSiftTestingData = []struct {
	startingGrid, resultGrid Grid
}{
	{
		Grid([][]uint8{
			{4, 4},
			{4, 4}}),
		Grid([][]uint8{
			{2, 2},
			{2, 2}})},
	{
		Grid([][]uint8{
			{4, 4, 4, 4},
			{4, 4, 4, 4},
			{4, 4, 4, 4},
			{4, 4, 4, 4}}),
		Grid([][]uint8{
			{2, 3, 3, 2},
			{3, 4, 4, 3},
			{3, 4, 4, 3},
			{2, 3, 3, 2}})},
}

func TestFullSift(t *testing.T) {
	for _, data := range fullSiftTestingData {
		data.startingGrid.sift()
		for x, row := range data.resultGrid {
			for y, cell := range row {
				if cell != data.startingGrid[x][y] {
					t.Errorf("Wrong value in grid at %d, %d, expected %d was %d", x, y, cell, data.startingGrid[x][y])
				}
			}
		}
	}
}

func TestIncrementBorder(t *testing.T) {
	border := NewBorder(15, 2, 8)

	asyncIncrement := func(signal chan bool, object *Border, x, y int) {
		object.incrementCell(x, y)
		signal <- true
	}
	channel := make(chan bool, 10)
	calls := 9
	for count := 0; count < calls; count++ {
		go asyncIncrement(channel, border, 14, 0)
		<-channel
	}
	if border.cells[14][0] != uint8(calls + 8) {
		t.Errorf("Wrong cell value. Expected %d got %d", calls, border.cells[14][0])
	}
}
