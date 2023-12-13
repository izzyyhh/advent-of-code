package day12

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	file, _ := os.Open("./day_12/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	ans := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		pattern, csvCountInfo := parts[0], parts[1]
		countInfo := strings.Split(csvCountInfo, ",")
		groupNums := []int{}

		for _, sNum := range countInfo {
			num, _ := strconv.Atoi(sNum)
			groupNums = append(groupNums, num)
		}

		ans += countArrangements(pattern, groupNums, 0)
	}

	fmt.Println(ans)
}

func isValidPattern(pattern string, groupNums []int) bool {
	count := 0
	occured := []int{}

	for _, ch := range pattern {
		if ch == '#' {
			count++
		} else if ch == '.' {
			if count != 0 {
				occured = append(occured, count)
				count = 0
			}
		} else {
			return false
		}
	}

	if count != 0 {
		occured = append(occured, count)
	}

	if len(occured) != len(groupNums) {
		return false
	}

	for i, occNum := range occured {
		if groupNums[i] != occNum {
			return false
		}
	}
	return true
}

func countArrangements(pattern string, groupNums []int, pi int) int {
	if pi == len(pattern) {
		if isValidPattern(pattern, groupNums) {
			return 1
		}
		return 0
	}

	if pattern[pi] == '?' {
		return countArrangements(pattern[:pi]+"."+pattern[pi+1:], groupNums, pi+1) + countArrangements(pattern[:pi]+"#"+pattern[pi+1:], groupNums, pi+1)
	} else {
		return countArrangements(pattern, groupNums, pi+1)
	}
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func binomialCoefficient(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	return factorial(n) / (factorial(k) * factorial(n-k))
}
