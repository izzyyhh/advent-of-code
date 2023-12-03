package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CubeColor string

const redCount = 12
const greenCount = 13
const blueCount = 14

var bag = map[CubeColor]int {
	"red": redCount,
	"green": greenCount,
	"blue": blueCount,
}

func Part1() {
	file, _ := os.Open("./day_2/input1.txt")
	scanner := bufio.NewScanner(file)
	lines := []string{}
	validGames := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)	
	}

	for i, line := range lines {
		gameId := i + 1
		_, setInfos, _ := strings.Cut(line, ":")
		gameSets := strings.Split(setInfos, ";")
		valid := true

		for i := range gameSets {
			cubeCounts := strings.Split(gameSets[i], ",")
			
			for j := range cubeCounts {
				cubeCount := strings.Trim(cubeCounts[j], " ")
				sCount, color, _ := strings.Cut(cubeCount, " ")
				count, _ := strconv.Atoi(sCount)
				bag[CubeColor(color)] -= count

				if !checkValidGame(bag) {
					valid = false
				}
				resetBag(bag)
			}
		}

		if valid {
			validGames = append(validGames, gameId)
		}

	}

	sum := 0

	for i := range validGames {
		sum += validGames[i]
	}

	fmt.Println(sum)
}

func Part2() {
	file, _ := os.Open("./day_2/input1.txt")
	scanner := bufio.NewScanner(file)
	lines := []string{}
	validGamePowers := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)	
	}

	for _, line := range lines {
		_, setInfos, _ := strings.Cut(line, ":")
		gameSets := strings.Split(setInfos, ";")
		maxVals := map[CubeColor]int{}
		
		for i := range gameSets {
			cubeCounts := strings.Split(gameSets[i], ",")
			
			for j := range cubeCounts {
				cubeCount := strings.Trim(cubeCounts[j], " ")
				sCount, color, _ := strings.Cut(cubeCount, " ")
				count, _ := strconv.Atoi(sCount)
				bag[CubeColor(color)] -= count
				maxVals[CubeColor(color)] = max(maxVals[CubeColor(color)], count)
			}
		}

		power := 1
		for _, val := range maxVals {
			power *= val
		}

		validGamePowers = append(validGamePowers, power)
	}

	sum := 0

	for i := range validGamePowers {
		sum += validGamePowers[i]
	}

	fmt.Println(sum)
}


func resetBag(bag map[CubeColor]int) {
	bag["red"] = redCount
	bag["green"] = greenCount
	bag["blue"] = blueCount
}

func checkValidGame(bag map[CubeColor]int) bool {
	for _, count := range bag {
		if count < 0 {
			return false 
		}
	}

	return true
}
