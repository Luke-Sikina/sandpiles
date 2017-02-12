package main

import (
	"os"
	"strconv"
)

func main() {
	dimension, err := strconv.Atoi(os.Args[1])
	if err != nil {
		println("Error: could not get the dimensions for grid from args")
	} else {
		secondDimension, err := strconv.Atoi(os.Args[2])
		if err != nil {
			secondDimension = dimension
		}
		startingHeight, err := strconv.ParseInt(os.Args[3], 10, 8)
		if err != nil {
			startingHeight = 4
		}
		xString := strconv.Itoa(int(dimension))
		yString := strconv.Itoa(int(secondDimension))
		heightString := strconv.Itoa(int(startingHeight))
		println("Creating grid of size " + xString + " x " + yString + " and height " + heightString)
		createGrid(dimension, secondDimension, uint8(startingHeight))
	}
}

type Grid [][]uint8
type Coordinate struct {
	x, y int
}

func (original Grid) sift() {
	reference := original.clone()
	for x, row := range reference {
		for y := range row {
			Coordinate{x, y}.sift(&reference, original)
		}
	}
}

func (center Coordinate) sift(reference *Grid, result Grid) {
	if (*reference)[center.x][center.y] > 3 {
		neighbors := center.getValidNeighbors(len(result), len(result[0]))
		result[center.x][center.y] = result[center.x][center.y] - 4
		for _, neighbor := range neighbors {
			result[neighbor.x][neighbor.y] = result[neighbor.x][neighbor.y] + 1
		}
	}
}

func (coordinate Coordinate) getValidNeighbors(xMax, yMax int) (validNeighbors []Coordinate) {
	top := Coordinate{coordinate.x, coordinate.y + 1}
	bot := Coordinate{coordinate.x, coordinate.y - 1}
	left := Coordinate{coordinate.x - 1, coordinate.y}
	right := Coordinate{coordinate.x + 1, coordinate.y}
	if top.isValidGridPosition(xMax, yMax) {
		validNeighbors = append(validNeighbors, top)
	}
	if bot.isValidGridPosition(xMax, yMax) {
		validNeighbors = append(validNeighbors, bot)
	}
	if left.isValidGridPosition(xMax, yMax) {
		validNeighbors = append(validNeighbors, left)
	}
	if right.isValidGridPosition(xMax, yMax) {
		validNeighbors = append(validNeighbors, right)
	}
	return
}

func (position Coordinate) isValidGridPosition(xMax, yMax int) bool {
	return position.x < xMax && position.x > -1 && position.y < xMax && position.y > -1
}

func (original Grid) clone() (duplicate Grid) {
	duplicate = make([][]uint8, len(original))
	for i := range original {
		duplicateRow := make([]uint8, len(original[0]))
		for j := range original[i] {
			duplicateRow[j] = original[i][j]
		}
		duplicate[i] = duplicateRow
	}
	return
}

func parseArgs(args []string) (xDim, yDim, startingHeight uint8) {
	xDim = parseNumOrUseDefault(args[1], 8)
	yDim = parseNumOrUseDefault(args[2], 8)
	startingHeight = parseNumOrUseDefault(args[3], 8)
	return
}

func parseNumOrUseDefault(toParse string, defaultNum uint8) (resultNum uint8) {
	parsed, err := strconv.ParseInt(toParse, 10, 8)
	if err == nil {
		resultNum = uint8(parsed)
	} else {
		resultNum = defaultNum
	}
	return
}

func createGrid(xAxis, yAxis int, height uint8) (grid Grid) {
	grid = make([][]uint8, xAxis)
	for i := range grid {
		grid[i] = make([]uint8, yAxis)
		for cell := range grid[i] {
			grid[i][cell] = height
		}
	}
	return
}
