package main

import (
	"fmt"
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
		createPiles(dimension, secondDimension, uint8(startingHeight))
		fmt.Print(getValidNeighbors(dimension, secondDimension, 0, 0))
	}
}

type Grid [][]uint8
type Coordinate struct {
	x, y int
}

//func (grid Grid) sift () () {
//	reference := copyGrid(grid)
//
//}

func getValidNeighbors(xMax, yMax, xPos, yPos int) (validNeighbors []Coordinate) {
	top := Coordinate{xPos, yPos + 1}
	bot := Coordinate{xPos, yPos - 1}
	left := Coordinate{xPos - 1, yPos}
	right := Coordinate{xPos + 1, yPos}
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

func copyGrid(original Grid) (duplicate Grid) {
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

func createPiles(xAxis, yAxis int, height uint8) (grid Grid) {
	grid = make([][]uint8, xAxis)
	for i := range grid {
		grid[i] = make([]uint8, yAxis)
		for cell := range grid[i] {
			print(" ")
			print(height)
			print(" ")
			grid[i][cell] = height
		}
		println()
	}
	return
}
