package day4

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func Part1() {
	file, _ := os.Open("./day_4/input.txt")
	scanner := bufio.NewScanner(file)
	ans := 0

	for scanner.Scan() {
		winning, drawn, _ := strings.Cut(scanner.Text(), "|")
		winning = winning[strings.Index(winning, ":")+2:]

		winningSlice := strings.Split(winning, " ")
		drawnSlice := strings.Split(drawn, " ")

		winningSlice = slices.DeleteFunc(winningSlice, func(s string) bool {
			return s == " " || s == ""
		})

		drawnSlice = slices.DeleteFunc(drawnSlice, func(s string) bool {
			return s == " " || s == ""
		})
		drawnMap := make(map[string]bool)

		for _, val := range drawnSlice {
			drawnMap[val] = true
		}
		matches := 0

		for _, winningNum := range winningSlice {
			if drawnMap[winningNum] {
				matches++
			}
		}
		if matches != 0 {
			ans += int(math.Pow(2, float64(matches-1)))
		}
	}

	fmt.Println(ans)
}

func Part2() {
	file, _ := os.Open("./day_4/input.txt")
	scanner := bufio.NewScanner(file)
	lines := []string{}
	cardCount := map[int]int{}

	gameId := 1

	for scanner.Scan() {
		lines = append(lines, scanner.Text())

		cardCount[gameId] = 1
		gameId++
	}

	for _, line := range lines {
		// do the same but get gameid, and iterate over matches and ++ the value in map
		// at the end accumulate the map
		fmt.Println(line)
	}
}
