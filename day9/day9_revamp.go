package day9

import (
	"fmt"
	"math"
	"strings"

	"main.go/util"
)

func Day9Revamp() {

	// Vérifier pour la taille du tableau
	// Part1()
	ropeMouves := util.ReadFile("day9/rope.txt")
	moves := fromListToMoves(strings.Split(string(ropeMouves), "\n"))
	fmt.Println("Nombre de cases visitées par la queue :", simulateMoves(moves))
}

type Position struct {
	row, col int
}

func simulateMoves(input []Move) int {
	head := Position{0, 0}
	tail := Position{0, 0}
	visited := make(map[Position]bool)
	visited[tail] = true

	for _, move := range input {
		for i := 0; i < move.Quantity; i++ {
			switch move.Step {
			case "U":
				head.col++
			case "D":
				head.col--
			case "R":
				head.row++
			case "L":
				head.row--
			}

			diffRow := head.row - tail.row
			diffCol := head.col - tail.col

			if abs(diffRow) > 1 || abs(diffCol) > 1 {
				tail.row += sign(diffRow)
				tail.col += sign(diffCol)
				visited[tail] = true
			}
		}
	}
	return len(visited)
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}
