package day1

import (
	"slices"
	"strconv"
	"strings"

	"main.go/util"
)

func Day1() {
	// Read calories.txt
	// For each values, calculate the calories
	// Add to a table the number of calories for each void in the file
	// Print the number of calories for the most calories
	txtCalories := util.ReadFile("day1/calories.txt")
	caloriesToCount := strings.Split(string(txtCalories), "\n")
	calories := make([]int, len(caloriesToCount))
	index := 0
	for _, calorie := range caloriesToCount {
		if calorie == "" {
			index++
			continue
		}
		calorieInt, err := strconv.Atoi(calorie)
		if err != nil {
			panic(err)
		}

		calories[index] += calorieInt
	}
	slices.Max(calories)
	println("max calories : ", calories[0])

	var top3Sum int

	for _, calorie := range calories[:3] {
		top3Sum += calorie
	}
	println("top 3 des elfs qui ont le plus de calories : ", top3Sum)

}
