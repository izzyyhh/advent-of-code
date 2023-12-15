package day14

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func Part1() {
	fileContent, _ := os.ReadFile("./day_14/input.txt")
	input := string(fileContent)
	grid := makeGridFromInput(&input)
	trickleUpStones(grid)
	ans := 0

	for rowi := range grid {
		for coli := range grid[rowi] {
			if grid[rowi][coli] == "O" {
				ans += len(grid) - rowi
			}
		}
	}

	fmt.Println(ans)
}

func Part2() {
	fileContent, _ := os.ReadFile("./day_14/input.txt")
	input := string(fileContent)
	grid := makeGridFromInput(&input)
	ans := 0

	cache := make(map[string]int)
	const cycles = 1000000000

	for cycle := 0; cycle < cycles; cycle++ {
		for dir := 0; dir < 4; dir++ {
			grid = trickleUpStones(grid)
			grid = rotateGridClockwise(grid)
		}

		gridHash := hashGrid(grid)

		if prevCycle, exists := cache[gridHash]; exists {
			cycleLength := cycle - prevCycle
			amt := (cycles - cycle) / (cycleLength)
			//jonathan paulson solution
			cycle += amt * cycleLength
		}

		cache[gridHash] = cycle
	}

	for rowi := range grid {
		for coli := range grid[rowi] {
			if grid[rowi][coli] == "O" {
				ans += len(grid) - rowi
			}
		}
	}

	fmt.Println(ans)
}

func makeGridFromInput(input *string) [][]string {
	rows := strings.Split(*input, "\n")
	grid := make([][]string, len(rows))

	for rowi := range grid {
		grid[rowi] = strings.Split(rows[rowi], "")
	}

	return grid
}

func trickleUpStones(grid [][]string) [][]string {
	for rowi := range grid {
		for coli := range grid[rowi] {
			currRow := rowi
			for currRow > 0 && grid[currRow-1][coli] == "." && grid[currRow][coli] == "O" {
				grid[currRow][coli] = "."
				grid[currRow-1][coli] = "O"
				currRow--
			}
		}
	}

	return grid
}

func rotateGridClockwise(grid [][]string) [][]string {
	rows := len(grid)
	cols := len(grid[0])

	rotatedGrid := make([][]string, cols)
	for i := range rotatedGrid {
		rotatedGrid[i] = make([]string, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotatedGrid[j][rows-i-1] = grid[i][j]
		}
	}

	return rotatedGrid
}

func rotateGridCounterClockwise(grid [][]string) [][]string {
	rows := len(grid)
	cols := len(grid[0])

	rotatedGrid := make([][]string, cols)
	for i := range rotatedGrid {
		rotatedGrid[i] = make([]string, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotatedGrid[cols-j-1][i] = grid[i][j]
		}
	}

	return rotatedGrid
}

func flattenGrid(grid [][]string) string {
	var flatGrid string
	for _, row := range grid {
		for _, cell := range row {
			flatGrid += cell
		}
	}
	return flatGrid
}

func hashGrid(grid [][]string) string {
	gridString := flattenGrid(grid)
	hasher := sha256.New()
	hasher.Write([]byte(gridString))
	hashSum := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashSum)

	return hashString
}
