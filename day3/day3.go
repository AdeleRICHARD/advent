package day3

import (
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
*/

func day3() {
	txtRucksack := util.ReadFile("day3/rucksack.txt")
	rucksacks := strings.Split(string(txtRucksack), "\n")
	for _, rucksack := range rucksacks {
	}
}
