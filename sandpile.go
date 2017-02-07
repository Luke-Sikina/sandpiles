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
		var startingHeight uint8 = 4
		createPiles(uint8(dimension), uint8(dimension), startingHeight)
	}
}

type Grid [][]uint8

func createPiles(xAxis, yAxis, height uint8) (grid Grid) {
	grid = make([][]uint8, xAxis)
	for i, _ := range grid {
		grid[i] = make([]uint8, yAxis)
		for cell, _ := range grid[i] {
			grid[i][cell] = height
			print(height)
		}
		println()
	}
	return
}
