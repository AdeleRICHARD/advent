package day9

import (
	"slices"
	"strconv"
	"strings"

	"main.go/util"
)

func Day9() {
	// Part1()
	ropeMouves := util.ReadFile("day9/rope.txt")
	moves := fromListToMoves(strings.Split(string(ropeMouves), "\n"))

	rope := [][]string{}
	head := []int{0, 0}
	tail := []int{0, 0}

	nbTailPassed := 0

	for _, move := range moves {
		// If nothing in the rope, means that H,T and s are all at the same place
		// See only head
		rope[0][0] = "H"
		switch move.Step {
		case "R":
			// Move right
			rope, head, tail = moveRight(rope, head, tail, move.Quantity)

		case "U":
			// Move up
			rope, head, tail = moveUp(rope, head, tail, move.Quantity)
		case "L":
			// Move left
			rope, head, tail = moveLeft(rope, head, tail, -move.Quantity)
		case "D":
			rope, head, tail = moveDown(rope, head, tail, -move.Quantity)
		}

		if countTailPassed(rope) > nbTailPassed {
			nbTailPassed = countTailPassed(rope)
		}
	}

}

type Move struct {
	Step     string
	Quantity int
}

func fromListToMoves(list []string) []Move {
	var moves []Move
	for _, move := range list {
		letter, number := strings.Split(move, " ")[0], strings.Split(move, " ")[1]
		nb, _ := strconv.Atoi(number)
		moves = append(moves, Move{letter, nb})
	}
	return moves
}

/* func GetHeightOfTable(moves []Move) [][]string {

	totalUp := 0
	totalRight := 0
	for _, move := range moves {
		switch move.Step {
		case "U":
			totalUp += move.Quantity
		case "R":
			totalRight += move.Quantity
		}
	}
	return make([][]string, totalUp)
} */

func isTouchingHead(head, tail []int) bool {
	//Diagonal
	if head[0] == tail[0]+1 && head[1] == tail[1]+1 {
		return true
	}
	// Adjacent
	if head[0] == tail[0]+1 && head[1] == tail[1] {
		return true
	} else if head[0] == tail[0] && head[1] == tail[1]+1 {
		return true
	}

	// Same spot
	return head[0] == tail[0] && head[1] == tail[1]
}

func moveRight(rope [][]string, head, tail []int, nbMove int) ([][]string, []int, []int) {
	// indexes
	rowHead := head[0]
	rowTail := tail[0]

	colTail := tail[1]
	colHead := head[1]
	// Move head to the right

	rope[rowHead][colHead] = "" // Where H was, put a #
	head[1] += nbMove
	colHead = head[1]

	if len(rope[rowHead]) < colHead {
		slices.Grow(rope[rowHead], colHead)
		for i := 0; i < nbMove; i++ {
			rope[rowHead] = append(rope[rowHead], "#")
		}
		// Add new columns
		rope[rowHead][colHead] = "H"
	} else {
		// Move head to the next column
		rope[rowHead][colHead] = "H"
	}

	if !isTouchingHead(head, tail) {
		// Where T was, put a #
		rope[rowTail][colTail] = "#"
		for i := 0; i < nbMove; i++ {
			rope[rowHead][colTail+i] = "#"
		}
		// Move tail to the right
		colTail = colHead - 1
		// Put tail up to date
		tail[0], tail[1] = rowHead, colTail
		rope[rowHead][colTail] = "T"
	}

	return rope, head, tail
}

func moveUp(rope [][]string, head, tail []int, nbMove int) ([][]string, []int, []int) {
	// indexes
	rowHead := head[0]
	rowTail := tail[0]

	colTail := tail[1]
	colHead := head[1]

	// Move head up
	rope[rowHead][colHead] = "" // Where H was, put a #
	head[0] += nbMove
	rowHead = head[0]

	if len(rope) < rowHead {
		for i := 0; i < nbMove; i++ {
			newRow := make([]string, colHead+1)
			rope = append(rope, newRow)
			rope[i+1][colHead] = "#"
		}
		// Add new columns
		rope[rowHead][colHead] = "H"
	} else {
		// Move head to the next column
		rope[rowHead][colHead] = "H"
	}

	if !isTouchingHead(head, tail) {
		// Where T was, put a #
		rope[rowTail][colTail] = "#"
		// Move tail under head
		rowTail = rowHead - 1
		colTail = colHead
		// Put tail up to date
		tail[0] = rowTail
		tail[1] = colTail
		rope[rowTail][colTail] = "T"
	}

	return rope, head, tail
}

func moveLeft(rope [][]string, head, tail []int, nbMove int) ([][]string, []int, []int) {
	// indexes
	rowHead := head[0]
	rowTail := tail[0]

	colTail := tail[1]
	colHead := head[1]

	// Move left
	rope[rowHead][colHead] = "" // Where H was, put a #

	if len(rope[rowHead]) < nbMove {
		for i := 0; i < nbMove; i++ {
			rope[rowHead] = slices.Insert(rope[rowHead], 0, "#")
		}

		// Add new columns
		rope[rowHead][colHead] = "H"
		head[1] = slices.Index(rope[rowHead], "H")
		tail[1] = nbMove
		colTail = tail[1]
	} else {
		colHead -= nbMove
		head[1] = colHead
		// Move head to the next column
		rope[rowHead][colHead] = "H"
	}

	if !isTouchingHead(head, tail) {
		// Where T was, put a #
		rope[rowTail][colTail] = "#"
		// Move tail to the right
		colTail = colHead + 1
		// Put tail up to date
		tail[0], tail[1] = rowHead, colTail
		rope[rowHead][colTail] = "T"
	}

	return rope, head, tail
}

func moveDown(rope [][]string, head, tail []int, nbMove int) ([][]string, []int, []int) {
	//TODO

	return rope, head, tail
}

func countTailPassed(rope [][]string) int {
	pass := 0
	for _, row := range rope {
		for _, col := range row {
			if col == "#" {
				pass++
			}
		}
	}
	return pass
}
