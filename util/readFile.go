package util

import (
	"os"
	"strconv"
)

func ReadFile(name string) []byte {
	txtRead, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return txtRead
}

// convertStringsToInts prend un slice de string et tente de les convertir en slice de int.
// Il retourne le slice de int et une erreur si une conversion échoue.
func ConvertStringsToInts(stringSlice []string) ([]int, error) {
	intSlice := make([]int, 0, len(stringSlice))
	for _, str := range stringSlice {
		// Convertit le string en int
		num, err := strconv.Atoi(str)
		if err != nil {
			// Gérer l'erreur si la conversion échoue
			return nil, err
		}
		// Ajouter le nombre converti au slice de int
		intSlice = append(intSlice, num)
	}
	return intSlice, nil
}
