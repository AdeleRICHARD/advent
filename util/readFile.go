package util

import "os"

func ReadFile(name string) []byte {
	txtRead, err := os.ReadFile("day2/strategy_guide.txt")
	if err != nil {
		panic(err)
	}
	return txtRead
}
