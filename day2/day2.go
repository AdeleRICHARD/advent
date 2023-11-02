package day2

import (
	"strings"

	"main.go/util"
)

func Day2() {
	// A for Rock, B for Paper, and C for Scissors
	//  X for Rock, Y for Paper, and Z for Scissors
	// The winner of the whole tournament is the player with the highest score.
	// Your total score is the sum of your scores for each round.
	// The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).
	txtStrategy := util.ReadFile("day2/strategy_guide.txt")

	strategyGuide := strings.Split(string(txtStrategy), "\n")
	score := 0
	scorePart2 := 0
	for _, line := range strategyGuide {
		if line == "" {
			continue
		}
		s := strings.Split(line, " ")
		player1 := GetType(s[0])
		player2 := GetType(s[1])
		score += getWinner(player1, player2) + player2.Number()
		// Part 2
		newPlayer2 := GetMoves(player1, s[1])
		scorePart2 += getWinner(player1, newPlayer2) + newPlayer2.Number()
	}

	println("Mon score : ", score)
	println("Mon score part 2 : ", scorePart2)
}

func getWinner(player1, player2 TypeOfPlay) int {
	// draw
	if player1 == player2 {
		return 3
	}

	if player2 == ROCK {
		if player1 == SCISSORS {
			return 6
		}
	}

	if player2 == PAPER {
		if player1 == ROCK {
			return 6
		}
	}

	if player2 == SCISSORS {
		if player1 == PAPER {
			return 6
		}
	}
	return 0
}
