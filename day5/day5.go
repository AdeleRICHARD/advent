package day5

import (
	"fmt"
	"strconv"
	"strings"

	"main.go/day5/model"
	"main.go/util"
)

func Day5() {
	cratesStacksTxt := util.ReadFile("day5/crates_stacks.txt")
	crates := strings.Split(string(cratesStacksTxt), "\n\n")
	// Split the text by newline to get each row
	cratesArray, moves := strings.Split(crates[0], "\n"), strings.Split(crates[1], "\n")
	stacks := parseInput(cratesArray)

	for _, move := range moves {
		moveCrate := buildMoveObject(move)
		// Move the crate
		/*
			Then, the rearrangement procedure is given. In each step of the procedure, a quantity of crates is moved from one stack to a different stack. In the first step of the above rearrangement procedure, one crate is moved from stack 2 to stack 1
		*/
		// Get the crate from the stack
		start := stacks[moveCrate.From]
		end := stacks[moveCrate.To]
		start, end = moveCrates(moveCrate, start, end)
		stacks[moveCrate.From] = start
		stacks[moveCrate.To] = end
	}

	fmt.Printf("Stacks: %v\n", stacks)

	lastCrates := getLastCrateByColumn(stacks)
	fmt.Printf("Last: %v\n", lastCrates)
}

func buildMoveObject(move string) model.Move {
	moveCrate := model.Move{}
	movesToParse := strings.Split(move, " ")
	for i, elem := range movesToParse {
		switch elem {
		case "move":
			moveNb, err := strconv.Atoi(movesToParse[i+1])
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
			moveCrate.NbMove = moveNb
		case "from":
			fromNb, err := strconv.Atoi(movesToParse[i+1])
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
			moveCrate.From = fromNb
		case "to":
			toNb, err := strconv.Atoi(movesToParse[i+1])
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
			moveCrate.To = toNb
		}

	}
	return moveCrate
}

func parseInput(crates []string) map[int][]string {
	stacks := make(map[int][]string)
	columnNbs := strings.Split(crates[len(crates)-1:][0], "")
	fmt.Printf("Crates: %v\n", columnNbs)

	for _, crate := range crates {
		elements := strings.Split(crate, "")
		for i, elem := range elements {
			if elem == columnNbs[i] {
				continue
			}
			if elem == "[" || elem == "]" {
				continue
			}

			nbColumn, err := strconv.Atoi(columnNbs[i])
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}

			if stacks[nbColumn] == nil {
				if elem != " " {
					stacks[nbColumn] = []string{elem}
				} else {
					stacks[nbColumn] = []string{}
				}
			} else if elem != " " {
				stacks[nbColumn] = append(stacks[nbColumn], elem)
			}
		}
	}

	return stacks
}

func getLastCrateByColumn(stacks map[int][]string) map[int]string {
	lastCrates := make(map[int]string)
	for column, stack := range stacks {
		if len(stack) == 0 {
			continue
		}
		lastCrates[column] = stack[0]
	}
	return lastCrates
}

func moveCrates(moveToDo model.Move, start, end []string) ([]string, []string) {
	if len(start) < moveToDo.NbMove {
		moveToDo.NbMove = len(start)
	}
	for i := 0; i < moveToDo.NbMove; i++ {
		crate := start[i] // Get the top crate from the stack
		// Add at start of the stack
		end = append([]string{crate}, end...) // Add the crate to the new stack
	}
	start = start[moveToDo.NbMove:] // Remove the crates from the stack

	return start, end
}
