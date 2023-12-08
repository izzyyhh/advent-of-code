package day8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Name  string
	Left  *Node
	Right *Node
}

func Part1() {
	file, _ := os.Open("./day_8/input.txt")
	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	instruction := lines[0]
	lines = lines[2:]

	nodes := buildNodes(lines)
	ans := followInstructions(nodes["AAA"], instruction)
	fmt.Println(ans)
}

func Part2() {
	file, _ := os.Open("./day_8/input2.txt")
	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	instruction := lines[0]
	lines = lines[2:]

	nodes := buildNodes(lines)
	starters := []*Node{}

	for _, node := range nodes {
		if strings.HasSuffix(node.Name, "A") {
			starters = append(starters, node)
		}
	}

	ans := followInstrctionsGhosts(starters, instruction)
	fmt.Println(ans)
}

func buildNodes(lines []string) map[string]*Node {
	nodeMap := make(map[string]*Node)

	for _, line := range lines {
		parts := strings.Split(line, " = ")
		name := parts[0]

		node := &Node{Name: name, Left: nil, Right: nil}
		nodeMap[node.Name] = node
	}

	for _, line := range lines {
		parts := strings.Split(line, " = ")
		name := parts[0]
		parts[1] = strings.Trim(parts[1], "()")
		children := strings.Split(parts[1], ", ")

		left := nodeMap[children[0]]
		right := nodeMap[children[1]]

		nodeMap[name].Left = left
		nodeMap[name].Right = right
	}

	return nodeMap
}

func followInstructions(start *Node, instructions string) int {
	currentNode := start
	steps := 0
	foundDestination := false

	for !foundDestination {
		for _, instruction := range instructions {
			steps++
			switch instruction {
			case 'L':
				currentNode = currentNode.Left
			case 'R':
				currentNode = currentNode.Right
			}

			if currentNode.Name == "ZZZ" {
				foundDestination = true
				break
			}
		}
	}

	return steps
}

func followInstructionStep(start *Node, instruction rune) *Node {
	currentNode := start

	switch instruction {
	case 'L':
		currentNode = currentNode.Left
	case 'R':
		currentNode = currentNode.Right
	}

	return currentNode
}

func followInstrctionsGhosts(ghosts []*Node, instructions string) int {
	stepsToCycleAll := []int{}

	for _, ghost := range ghosts {
		stepsToCycle := 0
		gotHome := false
		currentNode := ghost

		for !gotHome {
			for _, instruction := range instructions {
				stepsToCycle++
				destination := followInstructionStep(currentNode, instruction)
				currentNode = destination

				if strings.HasSuffix(currentNode.Name, "Z") {
					gotHome = true
					break
				}
			}
		}
		stepsToCycleAll = append(stepsToCycleAll, stepsToCycle)
	}

	return LCM(stepsToCycleAll)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(numbers []int) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}

	return result
}

func lcm(a, b int) int {
	return a * b / GCD(a, b)
}
