package day2

import (
	"os"
	"strings"
)

func Day2() {
	txtStrategy, err := os.ReadFile("strategy_guide.txt")
	if err != nil {
		panic(err)
	}

	strategyGuide := strings.Split(string(txtStrategy), "\n")
	for _, line := range strategyGuide {
		println(line)
	}
}
