package main

import (
	"os"
	"strconv"
)

func main() {
	dimension, err := strconv.ParseInt(os.Args[1], 10, 8)
	if err != nil {
		println("Error: could not get the dimensions for grid from args")
	} else {
		secondDimension, err := strconv.ParseInt(os.Args[2], 10, 8)
		if err != nil {
			secondDimension = dimension
		}
		startingHeight, err := strconv.ParseInt(os.Args[3], 10, 8)
		if err != nil {
			startingHeight = 4
		}
		grid := createPiles(uint8(dimension), uint8(secondDimension), uint8(startingHeight))
		copyGrid(grid)
	}
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

type Grid [][]uint8

func createPiles(xAxis, yAxis, height uint8) (grid Grid) {
	grid = make([][]uint8, xAxis)
	for i, _ := range grid {
		grid[i] = make([]uint8, yAxis)
		for cell, _ := range grid[i] {
			grid[i][cell] = height
		}
	}
	return
}
