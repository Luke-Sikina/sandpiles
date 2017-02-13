package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	xDim, yDim, height := parseArgs(os.Args)
	fmt.Printf("Creating grid of size %dx%d and height %d", xDim, yDim, height)
	createGrid(xDim, yDim, height)
}

type Grid [][]uint8

func (original Grid) sift() {
	reference := original.clone()
	for x, row := range reference {
		for y := range row {
			singleSift(x, y, &reference, original)
		}
	}
}

func singleSift(x, y int, reference *Grid, result Grid) {
	if (*reference)[x][y] > 3 {
		neighborsX, neighborsY := getValidNeighbors(x, y, len(result), len(result[0]))
		result[x][y] = result[x][y] - 4
		for i := range neighborsX {
			result[neighborsX[i]][neighborsY[i]] = result[neighborsX[i]][neighborsY[i]] + 1
		}
	}
}

func getValidNeighbors(x, y, xMax, yMax int) (validNeighborsX, validNeighborsY []int) {
	topX := x
	topY := y + 1
	botX := x
	botY := y - 1
	leftX := x - 1
	leftY := y
	rightX := x + 1
	rightY := y
	if isValidGridPosition(topX, topY, xMax, yMax) {
		validNeighborsX = append(validNeighborsX, topX)
		validNeighborsY = append(validNeighborsY, topY)
	}
	if isValidGridPosition(botX, botY, xMax, yMax) {
		validNeighborsX = append(validNeighborsX, botX)
		validNeighborsY = append(validNeighborsY, botY)
	}
	if isValidGridPosition(leftX, leftY, xMax, yMax) {
		validNeighborsX = append(validNeighborsX, leftX)
		validNeighborsY = append(validNeighborsY, leftY)
	}
	if isValidGridPosition(rightX, rightY, xMax, yMax) {
		validNeighborsX = append(validNeighborsX, rightX)
		validNeighborsY = append(validNeighborsY, rightY)
	}
	return
}

func isValidGridPosition(xPos, yPos, xMax, yMax int) bool {
	return xPos < xMax && xPos > -1 && yPos < xMax && yPos > -1
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

func parseArgs(args []string) (xDim, yDim int, startingHeight uint8) {
	switch len(args) {
	case 0:
		fallthrough
	case 1:
		xDim = 8
		yDim = 8
		startingHeight = uint8(8)
	case 2:
		xDim = parseNumOrUseDefault(args[1], 8)
		yDim = 8
		startingHeight = uint8(8)
	case 3:
		xDim = parseNumOrUseDefault(args[1], 8)
		yDim = parseNumOrUseDefault(args[2], 8)
		startingHeight = uint8(8)
	default:
		xDim = parseNumOrUseDefault(args[1], 8)
		yDim = parseNumOrUseDefault(args[2], 8)
		startingHeight = uint8(parseNumOrUseDefault(args[3], 8))
	}
	return
}

func parseNumOrUseDefault(toParse string, defaultNum int) (resultNum int) {
	parsed, err := strconv.Atoi(toParse)
	if err == nil {
		resultNum = parsed
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
