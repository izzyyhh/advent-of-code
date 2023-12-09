package day9

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1() {
	file, _ := os.Open("./day_9/input1.txt")
	scanner := bufio.NewScanner(file)
	ans := 0

	for scanner.Scan() {
		line := scanner.Text()
		history := parseInts(line)

		ans += jungleDiff(history)
	}

	fmt.Println(ans)
}

func Part2() {
	file, _ := os.Open("./day_9/input1.txt")
	scanner := bufio.NewScanner(file)
	ans := 0

	for scanner.Scan() {
		line := scanner.Text()
		history := parseInts(line)
		slices.Reverse(history)

		ans += jungleDiff(history)
	}

	fmt.Println(ans)
}

func jungleDiff(history []int) int {
	differences := make([]int, len(history)-1)
	for i := range differences {
		differences[i] = history[i+1] - history[i]
	}

	if !allZeros(differences) {
		return history[len(history)-1] + jungleDiff(differences)
	}
	return history[len(history)-1]
}

func allZeros(arr []int) bool {
	return !slices.ContainsFunc(arr, func(i int) bool { return i != 0 })
}

func parseInts(s string) []int {
	var result []int
	for _, v := range strings.Split(s, " ") {
		num, _ := strconv.Atoi(v)
		result = append(result, num)
	}
	return result
}
