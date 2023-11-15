package day6

import (
	"slices"
	"strings"

	"main.go/util"
)

func Day6() {
	communication := util.ReadFile("day6/communication.txt")
	// Split the text by newline to get each row
	communications := strings.Split(string(communication), "\n")
	differentLetters := []string{}
	/* 	var previousLetter string
	 */
	for _, line := range communications {
		for i, letter := range strings.Split(line, "") {
			if slices.Contains(differentLetters, letter) {
				indexLetter := slices.Index(differentLetters, letter)
				differentLetters = differentLetters[indexLetter+1:]
			}
			differentLetters = append(differentLetters, letter)
			// part 1
			/*
				if len(differentLetters) == 14 {
					println(i + 1)
					break
				}
			*/
			// part 2
			if len(differentLetters) == 14 {
				println(i + 1)
				break
			}
		}
	}
}
