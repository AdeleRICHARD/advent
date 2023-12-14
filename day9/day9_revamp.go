package day9

import (
	"fmt"
	"math"
	"strings"

	"main.go/util"
)

type Move struct {
	Step     string
	Quantity int
}

func Day9Revamp() {

	// Vérifier pour la taille du tableau
	// Part1()
	ropeMouves := util.ReadFile("day9/rope.txt")
	moves := fromListToMoves(strings.Split(string(ropeMouves), "\n"))
	fmt.Println("Part 1 - Nombre de cases visitées par la queue :", simulateMoves(moves))
	fmt.Println("Part 2 - Nombre de cases visitées par la queue :", simulateMovesWithKnots(moves))
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
			moveHead(&head, move)

			diffRow := head.row - tail.row
			diffCol := head.col - tail.col

			// Calcule la distance entre la tête et la queue
			// Si la distance est supérieure à 1, on déplace la queue
			if abs(diffRow) > 1 || abs(diffCol) > 1 {
				tail.row += sign(diffRow)
				tail.col += sign(diffCol)
				visited[tail] = true
			}
		}
	}
	return len(visited)
}

func moveHead(head *Position, move Move) {
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
}

func simulateMovesWithKnots(input []Move) int {
	knotsLength := 10
	// Slices of Positions with each knot
	rope := make([]Position, knotsLength)
	visited := make(map[Position]bool)
	visited[rope[len(rope)-1]] = true

	for _, move := range input {
		for i := 0; i < move.Quantity; i++ {
			moveHead(&rope[0], move)

			// Déplacer les autres nœuds
			for j := 1; j < knotsLength; j++ {
				updateKnotPosition(&rope[j], rope[j-1])
			}

			visited[rope[knotsLength-1]] = true // Marquer la position de la queue
		}
	}
	// 2739 too high
	// 2487 good
	return len(visited)
}

func updateKnotPosition(knot *Position, prevKnot Position) {
	// Calcule la différence entre la position du nœud précédent et la position du nœud actuel
	rowDiff := prevKnot.row - knot.row
	colDiff := prevKnot.col - knot.col

	// Si la différence est supérieure à 1, on déplace le nœud
	if abs(rowDiff) > 1 || abs(colDiff) > 1 {
		knot.row += sign(rowDiff)
		knot.col += sign(colDiff)
	}
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
