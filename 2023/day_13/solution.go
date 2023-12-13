package day13

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Part1() {
	file, _ := os.Open("./day_13/input.txt")
	defer file.Close()

	inputBytes, _ := io.ReadAll(file)
	input := string(inputBytes)
	ans := 0
	blocks := strings.Split(input, "\n\n")

	for _, block := range blocks {
		grid := makeGridFromBlockString(block)
		found := false

		for coli := 0; coli < len(grid[0])-1; coli++ {
			isReflectionCandidate := true
			reflectionCandidates := [2]int{}

			for rowi := range grid {
				if grid[rowi][coli] != grid[rowi][coli+1] {
					isReflectionCandidate = false
					break
				} else {
					reflectionCandidates[0], reflectionCandidates[1] = coli, coli+1
				}
			}

			if isReflectionCandidate {
				if checkReflectionRows(grid, reflectionCandidates) {
					ans += coli + 1
					found = true
					break
				}
			}
		}
		if found {
			continue
		}
		flippedGrid := flipGrid(grid)

		for coli := 0; coli < len(flippedGrid[0])-1; coli++ {
			isReflectionCandidate := true
			reflectionCandidates := [2]int{}

			for rowi := range flippedGrid {
				if flippedGrid[rowi][coli] != flippedGrid[rowi][coli+1] {
					isReflectionCandidate = false
					break
				} else {
					reflectionCandidates[0], reflectionCandidates[1] = coli, coli+1
				}
			}

			if isReflectionCandidate {
				if checkReflectionRows(flippedGrid, reflectionCandidates) {
					ans += (coli + 1) * 100
					found = true
					break
				}
			}
		}

	}
	fmt.Println(ans)
}

func checkReflectionRows(grid [][]string, reflectionCandidates [2]int) bool {
	for rowi := range grid {
		for i, j := reflectionCandidates[0], reflectionCandidates[1]; i >= 0 && j < len(grid[rowi]); i, j = i-1, j+1 {
			if grid[rowi][i] != grid[rowi][j] {
				return false
			}
		}
	}

	return true
}

func makeGridFromBlockString(block string) [][]string {
	rows := strings.Split(block, "\n")
	grid := make([][]string, len(rows))

	for y := range rows {
		splitted := strings.Split(rows[y], "")
		grid[y] = splitted
	}
	return grid
}

func flipGrid(grid [][]string) [][]string {
	rows := len(grid)
	if rows == 0 {
		return grid
	}

	cols := len(grid[0])
	flippedGrid := make([][]string, cols)

	for i := 0; i < cols; i++ {
		flippedGrid[i] = make([]string, rows)
		for j := 0; j < rows; j++ {
			flippedGrid[i][j] = grid[j][i]
		}
	}

	return flippedGrid
}
