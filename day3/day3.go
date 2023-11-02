package day3

import (
	"fmt"
	"strings"

	"main.go/util"
)

/*
	The first rucksack contains the items vJrwpWtwJgWrhcsFMMfFFhFp, which means its first compartment contains the items vJrwpWtwJgWr, while the second compartment contains the items hcsFMMfFFhFp. The only item type that appears in both compartments is lowercase p.
	The second rucksack's compartments contain jqHRNqRjqzjGDLGL and rsFMfFZSrLrFZsSL. The only item type that appears in both compartments is uppercase L.
	The third rucksack's compartments contain PmmdzqPrV and vPwwTWBwg; the only common item type is uppercase P.
	The fourth rucksack's compartments only share item type v.
	The fifth rucksack's compartments only share item type t.
	The sixth rucksack's compartments only share item type s.

	Lowercase item types a through z have priorities 1 through 26.
	Uppercase item types A through Z have priorities 27 through 52.
*/

func Day3() {
	txtRucksack := util.ReadFile("day3/rucksacks.txt")
	rucksacks := strings.Split(string(txtRucksack), "\n")

	commonElemPrioritySum := 0

	for _, rucksack := range rucksacks {
		// Split the rucksack in two compartments
		compartment1, compartment2 := rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]

		// Create a map of items for each compartment
		itemsMap1, itemsMap2 := fromSliceToMap(strings.Split(compartment1, "")), fromSliceToMap(strings.Split(compartment2, ""))

		// Compare the two maps and find the common item
		commonElement := compareMaps(itemsMap1, itemsMap2)
		if commonElement != "" {
			// Convert the common item to its priority and add to the total
			elementPriority := int(commonElement[0])
			if commonElement >= "a" && commonElement <= "z" {
				commonElemPrioritySum += elementPriority - int('a') + 1
			} else {
				commonElemPrioritySum += elementPriority - int('A') + 27
			}
		}
	}

	fmt.Println(commonElemPrioritySum)
}

func fromSliceToMap(slice []string) map[string]struct{} {
	itemsMap := make(map[string]struct{})
	for _, item := range slice {
		itemsMap[item] = struct{}{}
	}
	return itemsMap
}

func compareMaps(map1, map2 map[string]struct{}) string {
	for item := range map1 {
		if _, ok := map2[item]; ok {
			return item
		}
	}
	return ""
}
