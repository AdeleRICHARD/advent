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
	commonElemPrioritySumPart2 := 0
	commonThreeElements := make([]string, 0, 3)
	for _, rucksack := range rucksacks {
		// Split the rucksack in two compartments
		compartment1, compartment2 := rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]

		// Create a map of items for each compartment
		itemsMap1, itemsMap2 := fromSliceToMap(strings.Split(compartment1, "")), fromSliceToMap(strings.Split(compartment2, ""))

		// Compare the two maps and find the common item
		commonElement := compareMapsOneByOne(itemsMap1, itemsMap2)
		commonThreeElements = append(commonThreeElements, rucksack)
		if len(commonThreeElements) == 3 {
			// Compare the three maps and find the common item
			commonElement = compareMapsThreeByThree(commonThreeElements)
			commonElemPrioritySumPart2 += calculatePriority(commonElement)
			commonThreeElements = make([]string, 0, 3)
		}

		commonElemPrioritySum += calculatePriority(commonElement)

	}

	fmt.Println(commonElemPrioritySum)
	fmt.Println("Total of 3 by 3 : ", commonElemPrioritySumPart2)
}

func fromSliceToMap(slice []string) map[string]struct{} {
	itemsMap := make(map[string]struct{})
	for _, item := range slice {
		itemsMap[item] = struct{}{}
	}
	return itemsMap
}

func compareMapsOneByOne(map1, map2 map[string]struct{}) string {
	for item := range map1 {
		if _, ok := map2[item]; ok {
			return item
		}
	}
	return ""
}

func compareMapsThreeByThree(threeItems []string) string {
	item1, item2, item3 := fromSliceToMap(strings.Split(threeItems[0], "")), fromSliceToMap(strings.Split(threeItems[1], "")), fromSliceToMap(strings.Split(threeItems[2], ""))

	for item := range item1 {
		if _, ok := item2[item]; ok {
			if _, ok := item3[item]; ok {
				return item
			}
		}
	}
	return ""
}

func calculatePriority(commonElement string) int {
	if commonElement != "" {
		// Convert the common item to its priority and add to the total
		elementPriority := int(commonElement[0])
		if commonElement >= "a" && commonElement <= "z" {
			return elementPriority - int('a') + 1
		} else {
			return elementPriority - int('A') + 27
		}
	}
	return 0
}
