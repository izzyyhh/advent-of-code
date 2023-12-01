package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Part1() {
	file, err := os.Open("day_1/input1.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	sum := 0
	scanner := bufio.NewScanner(file)

	
	for scanner.Scan() {
		indices := [2]int{0,0} 
		fFound, lFound := false, false
		line := scanner.Text()

		for i,j := 0, len(line) - 1; !fFound || !lFound ;{

			if unicode.IsDigit(rune(line[i])) {
				indices[0] = i
				fFound = true
			} else  {
				i++
			}

			if unicode.IsDigit(rune(line[j])) {
				indices[1] = j
				lFound = true
			} else {
				j--
			}
		}

		concatenatedDigits := string(line[indices[0]]) + string(line[indices[1]])
		num, _ := strconv.Atoi(concatenatedDigits)
		sum += num

	}
	fmt.Println(sum)
}

func Part2() {
	digitTable := map[string]string{
		"one": "1",
		"two": "2",
		"three": "3",
		"four": "4",
		"five": "5",
		"six": "6",
		"seven": "7",
		"eight": "8",
		"nine": "9",
	}

	keys := make([]string, 0, len(digitTable))
	for k := range digitTable {
		keys = append(keys, k)
	}

	file, err := os.Open("day_1/input1.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		digits := []string{}
		line := scanner.Text()

		for i, c := range line {
			if unicode.IsDigit(c) {
				digits = append(digits, string(c))
			}

			for _, key := range keys {
				if strings.HasPrefix(line[i:], key) {
					digits = append(digits, digitTable[key])
				}
			}
		}
		num, _ := strconv.Atoi(digits[0] + digits[len(digits) - 1])
		sum += num
	}

	fmt.Println(sum)
}
