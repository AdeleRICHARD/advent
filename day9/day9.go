package day9

import (
	"strconv"
	"strings"
)

// First try but not working
// Too complicated
/*
	 func Day9() {

		// Vérifier pour la taille du tableau
		// Part1()
		ropeMouves := util.ReadFile("day9/rope.txt")
		moves := fromListToMoves(strings.Split(string(ropeMouves), "\n"))

		gridHeight, gridWidth, startPosX, _ := calculateExtremes(moves)
		rope := make([][]string, gridHeight)
		for i := 0; i < gridHeight; i++ {
			rope[i] = make([]string, gridWidth)
		}
		head := []int{135, startPosX}
		knots := map[string][]int{
			"1": {135, startPosX},
			"2": {135, startPosX},
			"3": {135, startPosX},
			"4": {135, startPosX},
			"5": {135, startPosX},
			"6": {135, startPosX},
			"7": {135, startPosX},
			"8": {135, startPosX},
			"9": {135, startPosX},
		}
		tail := []int{135, startPosX}

		nbTailPassed := 0
		tailsPassed := make(map[string]int)
		rope[head[0]][head[1]] = "H"

		tailCount := 0
		for i, move := range moves {
			// If nothing in the rope, means that H,T and s are all at the same place
			// See only head
			if i == 168 {
				println("i: ", i)
			}
			switch move.Step {
			case "R":
				// Move right
				rope, head, tail = moveRight(rope, head, knots, move.Quantity, &tailCount)
			case "U":
				// Move up
				rope, head, tail = moveUp(rope, head, knots, move.Quantity, &tailCount)
			case "L":
				// Move left
				rope, head, tail = moveLeft(rope, head, knots, move.Quantity, &tailCount)
			case "D":
				rope, head, tail = moveDown(rope, head, knots, move.Quantity, &tailCount)
			}
			if i == 0 {
				rope[tail[0]][tail[1]] = "T"
			}
			// If tail is already in the map do not count it

			if countTailPassed(rope) > nbTailPassed {
				nbTailPassed = countTailPassed(rope)
			}

			if _, ok := tailsPassed[fmt.Sprintf("%v", tail)]; !ok {
				tailsPassed[fmt.Sprintf("%v", tail)] = 1
			}
			//printInFile(rope)
		}
		println("tailCount: ", tailCount)
		fmt.Printf("head: %v\n", head)
		fmt.Printf("tail: %v\n", tail)
		fmt.Printf("Nb tail passed %v: ", nbTailPassed)

		fmt.Printf("tails passed: %v\n", len(tailsPassed))
		// 1988 too low
		// 5041 too low
		// 5165 too low
		// 5174 not good
		// 7324 not good
	}
*/

/* func printInFile(rope [][]string) {
	file, err := os.Create("day9/ropeTest.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, row := range rope {
		for _, col := range row {
			if col != "" {
				fmt.Fprintf(file, "%s", col)
			}
		}
		fmt.Fprintf(file, "%s", "\n")
	}
} */

func fromListToMoves(list []string) []Move {
	var moves []Move
	for _, move := range list {
		letter, number := strings.Split(move, " ")[0], strings.Split(move, " ")[1]
		nb, _ := strconv.Atoi(number)
		moves = append(moves, Move{letter, nb})
	}
	return moves
}

/* func isTouchingHead(head, tail []int) bool {
	// Check for same position
	if head[0] == tail[0] && head[1] == tail[1] {
		return true
	}

	// Check for adjacent positions (up, down, left, right)
	if head[0] == tail[0] && (head[1] == tail[1]+1 || head[1] == tail[1]-1) {
		return true
	}
	if head[1] == tail[1] && (head[0] == tail[0]+1 || head[0] == tail[0]-1) {
		return true
	}

	// Check for diagonal positions
	if (head[0] == tail[0]+1 || head[0] == tail[0]-1) && (head[1] == tail[1]+1 || head[1] == tail[1]-1) {
		return true
	}

	return false
} */

/* func moveRight(rope [][]string, head []int, knots map[string][]int, nbMove int, count *int) ([][]string, []int, []int) {
	// indexes
	rowHead := head[0]
	colHead := head[1]
	// Move head to the right

	rope[rowHead][colHead] = "" // Where H was, put a #
	head[1] += nbMove
	colHead = head[1]

	// Move head to the next column
	rope[rowHead][colHead] = "H"
	for !isTouchingHead(head, tail) {
		// Where T was, put a #
		if rope[rowTail][colTail] != "#" {
			rope[rowTail][colTail] = "#"
			*count++
		}
		// Move tail to the right
		colTail = colTail + 1
		// Put tail up to date
		rowTail = rowHead
		tail[0], tail[1] = rowTail, colTail
		rope[rowTail][colTail] = "T"
	}

	return rope, head, tail
} */

/* func moveUp(rope [][]string, head, tail []int, nbMove int, count *int) ([][]string, []int, []int) {
	// indexes
	rowHead := head[0]
	rowTail := tail[0]

	colTail := tail[1]
	colHead := head[1]

	// Move head up
	fmt.Printf("rowHead: %d, colHead: %d\n", rowHead, colHead)
	println("nbMove: ", nbMove)

	rope[rowHead][colHead] = "" // Where H was, put a ""
	head[0] -= nbMove
	rowHead = head[0]

	// Move head to the next column
	rope[rowHead][colHead] = "H"

	for !isTouchingHead(head, tail) {
		// Where T was, put a #
		if rope[rowTail][colTail] != "#" {
			rope[rowTail][colTail] = "#"
			*count++
		}
		// Move tail under head
		rowTail = rowTail - 1
		colTail = colHead
		// Put tail up to date
		tail[0], tail[1] = rowTail, colTail
		rope[rowTail][colTail] = "T"
	}

	return rope, head, tail
} */

/* func moveLeft(rope [][]string, head, tail []int, nbMove int, count *int) ([][]string, []int, []int) {
	// indexes
	rowHead := head[0]
	rowTail := tail[0]

	colTail := tail[1]
	colHead := head[1]

	// Move left
	rope[rowHead][colHead] = "" // Where H was, put a #
	println("colHead: ", colHead)
	println("rowHead: ", rowHead)
	colHead -= nbMove
	head[1] = colHead
	// Move head to the next column
	rope[rowHead][colHead] = "H"

	if !isTouchingHead(head, tail) {
		// Where T was, put a #
		if rope[rowTail][colTail] != "#" {
			rope[rowTail][colTail] = "#"
			*count++
		}
		rope[rowTail][colTail] = "#"

		// Move tail to the right
		colTail = colHead + 1
		// Put tail up to date
		rowTail = rowHead
		tail[0], tail[1] = rowTail, colTail
		rope[rowHead][colTail] = "T"
	}

	return rope, head, tail
} */

/* func moveDown(rope [][]string, head, tail []int, nbMove int, count *int) ([][]string, []int, []int) {
	// indexes
	rowHead := head[0]
	rowTail := tail[0]

	colTail := tail[1]
	colHead := head[1]

	// Move head down
	rope[rowHead][colHead] = "" // Where H was, nothing
	head[0] += nbMove
	rowHead = head[0]

	// Move head to the next column
	rope[rowHead][colHead] = "H"

	for !isTouchingHead(head, tail) {
		// Where T was, put a #
		if rope[rowTail][colTail] != "#" {
			rope[rowTail][colTail] = "#"
			*count++
		}

		// Move tail
		rowTail = rowTail + 1
		colTail = colHead
		// Put tail up to date
		tail[0] = rowTail
		tail[1] = colTail
		rope[rowTail][colTail] = "T"
	}

	return rope, head, tail
} */

/* func countTailPassed(rope [][]string) int {
	pass := 0
	t := 0
	for _, row := range rope {
		for _, col := range row {
			if col == "#" {
				pass++
			} else if col == "T" {
				t++
			}
		}
	}
	if t != 1 {
		fmt.Println("There should be only one T")
	}
	return pass
}

func calculateExtremes(moves []Move) (int, int, int, int) {
	maxUp := 0
	maxDown := 0
	maxRight := 0
	maxLeft := 0

	currentVertical := 0
	currentHorizontal := 0

	for _, move := range moves {
		switch move.Step {
		case "U":
			currentVertical += move.Quantity
			if currentVertical > maxUp {
				maxUp = currentVertical
			}
		case "D":
			currentVertical -= move.Quantity
			if -currentVertical > maxDown {
				maxDown = -currentVertical
			}
		case "R":
			currentHorizontal += move.Quantity
			if currentHorizontal > maxRight {
				maxRight = currentHorizontal
			}
		case "L":
			currentHorizontal -= move.Quantity
			if -currentHorizontal > maxLeft {
				maxLeft = -currentHorizontal
			}
		}
	}

	// The total dimensions needed for the grid
	gridHeight := maxUp + maxDown + 1
	gridWidth := maxRight + maxLeft + 1

	// Starting position within the grid
	startPosX := maxLeft
	startPosY := maxDown

	return gridHeight, gridWidth, startPosX, startPosY
} */
