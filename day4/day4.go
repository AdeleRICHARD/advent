package day4

import (
	"strconv"
	"strings"

	"main.go/util"
)

func Day4() {
	txtAssignment := util.ReadFile("day4/assignments.txt")
	assignments := strings.Split(string(txtAssignment), "\n")
	countContainAllShifts := 0
	for _, assignment := range assignments {
		pair1, pair2 := strings.Split(assignment, ",")[0], strings.Split(assignment, ",")[1]
		if contains(pair1, pair2) || contains(pair2, pair1) {
			countContainAllShifts++
		}
	}
	println("Shifts contained in other shifts : ", countContainAllShifts)
}

func parseRange(r string) (int, int) {
	parts := strings.Split(r, "-")
	start, err1 := strconv.Atoi(parts[0])
	end, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		panic("Invalid range")
	}
	return start, end
}

func contains(pair1, pair2 string) bool {
	start1, end1 := parseRange(pair1)
	start2, end2 := parseRange(pair2)
	return start1 <= start2 && end1 >= end2
}
