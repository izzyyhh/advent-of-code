package day15

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1() {
	fileContent, _ := os.ReadFile("./day_15/input.txt")
	input := string(fileContent)
	instructions := strings.Split(input, ",")

	ans := 0
	for _, instruction := range instructions {
		ans += hash(instruction)
	}

	fmt.Println(ans)
}

type Lense struct {
	label  string
	number int
}

func Part2() {
	fileContent, _ := os.ReadFile("./day_15/input.txt")
	input := string(fileContent)
	instructions := strings.Split(input, ",")
	boxes := make(map[int][]*Lense, 256)

	for _, instruction := range instructions {
		if instruction[len(instruction)-1] == '-' {
			label, _ := strings.CutSuffix(instruction, "-")

			box := hash(label)
			if boxContent, exists := boxes[box]; exists {
				i := slices.IndexFunc(boxContent, func(l *Lense) bool { return l.label == label })
				if i != -1 {
					boxes[box] = append(boxContent[:i], boxContent[i+1:]...)
				}
			}

		} else {
			label, snum, _ := strings.Cut(instruction, "=")
			box := hash(label)
			num, _ := strconv.Atoi(snum)
			newLense := Lense{label: label, number: num}

			if boxContent, exists := boxes[box]; exists {
				i := slices.IndexFunc(boxContent, func(l *Lense) bool { return l.label == newLense.label })
				if i != -1 {
					boxContent[i].number = newLense.number
				} else {
					boxContent = append(boxContent, &newLense)
					boxes[box] = boxContent
				}
			} else {
				boxes[box] = make([]*Lense, 1)
				new := append([]*Lense{}, &newLense)
				boxes[box] = new
			}
		}
	}

	ans := 0
	for x, box := range boxes {
		for i, lense := range box {
			slotNumber := i + 1
			ans += (x + 1) * slotNumber * lense.number
		}
	}
	fmt.Println(ans)
}

func hash(s string) int {
	value := 0

	for _, ch := range s {
		value += int(ch)
		value *= 17
		value = value % 256
	}

	return value
}
