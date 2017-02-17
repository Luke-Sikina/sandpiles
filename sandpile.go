package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func main() {
	core := runtime.NumCPU()
	fmt.Printf("There are %d cpus available. This uses one of them. Hmmm...\n", core)
	xDim, yDim, height := parseArgs(os.Args)
	fmt.Printf("Creating grid of size %dx%d and height %d\n", xDim, yDim, height)
	grid := createGrid(xDim, yDim, height)

	passes := 1
	for !grid.sift() {
		passes++
		if passes%100 == 0 {
			fmt.Printf("Pass #%d\n", passes)
		}
	}
	fmt.Printf("Finished in %d passes. Printing result\n", passes)
	grid.print()
}

type Grid [][]uint8

/**
Assumption: since we're only splitting the grid into 4 parts,
each subgrid has 2 borders shared with one other subgrid,
and one border shared with the other 3 subgrids
*/
type SuperGrid struct {
	northWest, northEast, southWest, southEast SubGrid
	northB, southB, eastB, westB, centerB      Border
}

type BorderCode int

const (
	north  BorderCode = iota
	south  BorderCode = iota
	east   BorderCode = iota
	west   BorderCode = iota
	center BorderCode = iota
)

type SubGrid struct {
	core    [][]uint8
	borders map[BorderCode]*Border
}

func (grid *SubGrid) incrementCell(x, y int) {
	grid[x][y]++
}

func (grid *SubGrid) siftCell(x, y int) {
	grid[x][y] = grid[x][y] - 4
}

/**
Visualization of the structure (+ = Border):
|-----++-----|
| sub ++ sub |
|     ++     |
|++++++++++++|
|     ++     |
| sub ++ sub |
|-----++-----|
*/
func NewSuperGrid(xAxis, yAxis int, height uint8) (grid *SuperGrid) {
	centerB := NewBorder(2, 2, height)
	northB := NewBorder(2, yAxis/2-2, height)
	southB := NewBorder(2, yAxis/2-2, height)
	eastB := NewBorder(xAxis/2-2, 2, height)
	westB := NewBorder(xAxis/2-2, 2, height)

	grid.northWest.borders[north] = northB
	grid.northWest.borders[west] = westB
	grid.northWest.borders[center] = centerB
	grid.northWest.core = createGrid(xAxis/2-1, yAxis/2-1, height)

	grid.northEast.borders[north] = northB
	grid.northEast.borders[east] = eastB
	grid.northEast.borders[center] = centerB
	grid.northEast.core = createGrid(xAxis/2-1, yAxis/2-1, height)

	grid.southWest.borders[south] = southB
	grid.southWest.borders[west] = westB
	grid.southWest.borders[center] = centerB
	grid.southWest.core = createGrid(xAxis/2-1, yAxis/2-1, height)

	grid.southEast.borders[south] = southB
	grid.southEast.borders[east] = eastB
	grid.southEast.borders[center] = centerB
	grid.southEast.core = createGrid(xAxis/2-1, yAxis/2-1, height)
	return
}

type Border struct {
	cells  Grid
	locker chan bool
}

/**
Thread safe borders for sub grid communication
*/
func NewBorder(xDim, yDim int, height uint8) *Border {
	grid := createGrid(xDim, yDim, height)
	lock := make(chan bool, 1)
	border := Border{cells: grid, locker: lock}
	return &border
}

func (border Border) lock() {
	border.locker <- true
}

func (border Border) release() {
	<-border.locker
}

func (border *Border) incrementCell(x, y int) {
	border.lock()
	border.cells[x][y]++
	border.release()
}

func (border *Border) siftCell(x, y int) {
	border.lock()
	border.cells[x][y] = border.cells[x][y] - 4
	border.release()
}

func (grid Grid) print() {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf(prettyPrintCell(cell))
		}
		fmt.Printf("\n")
	}
}

func (grid Grid) equals(compareTo Grid) (equals bool) {
	equals = len(grid) == len(compareTo) && len(grid[0]) == len(compareTo[0])
	for x, row := range grid {
		for y, height := range row {
			equals = equals && height == compareTo[x][y]
			if !equals {
				return
			}
		}
	}
	return
}

func (original Grid) sift() bool {
	reference := original.clone()
	for x, row := range reference {
		for y := range row {
			singleSift(x, y, &reference, original)
		}
	}
	return reference.equals(original)
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

func prettyPrintCell(cell uint8) (text string) {
	switch cell % 4 {
	case 0:
		text = "  "
	case 1:
		text = " o"
	case 2:
		text = " x"
	case 3:
		text = " *"
	}
	return
}
