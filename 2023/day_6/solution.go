package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	file, _ := os.Open("./day_6/input.txt")
	scanner := bufio.NewScanner(file)
	times := []int{}
	distances := []int{}
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i, line := range lines {
		if i == 0 {
			_, sTimes, _ := strings.Cut(line, ":")
			parsed := parseIntSlice(strings.Fields(sTimes))
			times = append(times, parsed...)

		} else {
			_, sDistances, _ := strings.Cut(line, ":")
			parsed := parseIntSlice(strings.Fields(sDistances))
			distances = append(distances, parsed...)
		}
	}

	answer := 1

	for i := range times {
		answer *= waysToDistance(times[i], distances[i])
	}

	fmt.Println(answer)
}

func Part2() {
	file, _ := os.Open("./day_6/input.txt")
	scanner := bufio.NewScanner(file)
	var time int
	var distance int
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i, line := range lines {
		if i == 0 {
			_, sTimes, _ := strings.Cut(line, ":")
			time, _ = strconv.Atoi(strings.Join(strings.Fields(sTimes), ""))

		} else {
			_, sDistances, _ := strings.Cut(line, ":")
			distance, _ = strconv.Atoi(strings.Join(strings.Fields(sDistances), ""))
		}
	}

	fmt.Println(waysToDistance(time, distance))
}

func waysToDistance(t int, d int) int {
	ways := 0

	for x := 0; x <= t; x++ {
		reached := (t - x) * x
		if reached > d {
			ways++
		}
	}
	return ways
}

func parseIntSlice(s []string) []int {
	result := make([]int, len(s))

	for i, sNum := range s {
		num, _ := strconv.Atoi(sNum)
		result[i] = num
	}

	return result
}
