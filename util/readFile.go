package util

import "os"

func ReadFile(name string) []byte {
	txtRead, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return txtRead
}
