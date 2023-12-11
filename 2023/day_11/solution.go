package day11

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Galaxy struct {
	x int
	y int
}

func Part1() {
	file, _ := os.Open("./day_11/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	galaxies := []*Galaxy{}
	emptyRows := []int{}
	emptyCols := []int{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for y, line := range lines {
		isEmpty := true
		for x, ch := range line {
			if ch == '#' {
				isEmpty = false
				galaxies = append(galaxies, &Galaxy{x, y})
			}
		}
		if isEmpty {
			emptyRows = append(emptyRows, y)
		}
	}

	for x := range lines[0] {
		isEmptyCol := true
		for y := 0; y < len(lines); y++ {
			if lines[y][x] == '#' {
				isEmptyCol = false
				break
			}
		}
		if isEmptyCol {
			emptyCols = append(emptyCols, x)
		}
	}

	for _, g := range galaxies {
		g.expand(emptyRows, emptyCols, 1)
	}

	ans := 0
	for i := range galaxies {
		for j := i; j < len(galaxies); j++ {
			ans += distance(galaxies[i], galaxies[j])
		}
	}

	fmt.Println(ans)
}

func Part2() {
	file, _ := os.Open("./day_11/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	galaxies := []*Galaxy{}
	emptyRows := []int{}
	emptyCols := []int{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for y, line := range lines {
		isEmpty := true
		for x, ch := range line {
			if ch == '#' {
				isEmpty = false
				galaxies = append(galaxies, &Galaxy{x, y})
			}
		}
		if isEmpty {
			emptyRows = append(emptyRows, y)
		}
	}

	for x := range lines[0] {
		isEmptyCol := true
		for y := 0; y < len(lines); y++ {
			if lines[y][x] == '#' {
				isEmptyCol = false
				break
			}
		}
		if isEmptyCol {
			emptyCols = append(emptyCols, x)
		}
	}

	for _, g := range galaxies {
		g.expand(emptyRows, emptyCols, 999999)
	}

	ans := 0
	for i := range galaxies {
		for j := i; j < len(galaxies); j++ {
			ans += distance(galaxies[i], galaxies[j])
		}
	}

	fmt.Println(ans)
}

func distance(g1, g2 *Galaxy) int {
	return int(math.Abs(float64(g1.x-g2.x)) + math.Abs(float64(g1.y-g2.y)))
}

func (g *Galaxy) expand(emptyRows, emptyCols []int, expFactor int) {
	xExp, yExp := 0, 0
	for _, x := range emptyCols {
		if x < g.x {
			xExp++
		}
	}

	for _, y := range emptyRows {
		if y < g.y {
			yExp++
		}
	}

	g.x += xExp * expFactor
	g.y += yExp * expFactor
}
