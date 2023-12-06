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
	ans := 0

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i, line := range lines {
		winning, drawn, _ := strings.Cut(line, "|")
		indexColon := strings.Index(winning, ":")
		winning = winning[indexColon+2:]
		cardCount[i] += 1

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
			for m := 0; m < matches; m++ {
				cardCount[i+1+m] += cardCount[i]
			}
		}
	}
	fmt.Println(cardCount)
	for _, cardN := range cardCount {
		ans += cardN
	}

	fmt.Println(ans)
}
