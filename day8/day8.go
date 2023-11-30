package day8

import (
	"slices"
	"strings"

	"main.go/util"
)

func Day8() {
	trees := strings.Split(string(util.ReadFile("day8/trees.txt")), "\n")
	// Transform tree into an double array of int
	treesHeights := [][]int{}
	for _, tree := range trees {
		treeRows, _ := util.ConvertStringsToInts(strings.Split(tree, ""))
		treesHeights = append(treesHeights, treeRows)
	}
	countVisibleTrees := 0
	for i := 0; i < len(treesHeights); i++ {
		if i == 0 || i == len(treesHeights)-1 {
			countVisibleTrees += len(treesHeights[i])
			continue
		}
		for j := 0; j < len(treesHeights[i]); j++ {
			// Visible trees on the edges
			if j == 0 || j == len(treesHeights[i])-1 {
				countVisibleTrees++
				continue
			}
			// Check top
			currentTree := treesHeights[i][j]
			if checkTopAndBottom(treesHeights[:i], currentTree, j) &&
				checkTopAndBottom(treesHeights[i:], currentTree, j) &&
				checkAdjacentTrees(treesHeights[i][:j], currentTree) &&
				checkAdjacentTrees(treesHeights[i][j:len(treesHeights[i])], currentTree) {

				countVisibleTrees++
			}

		}
	}
	println("VISIBLE TREES : ", countVisibleTrees)
}

func checkTopAndBottom(trees [][]int, currentTree, colNb int) bool {
	max := 0
	for i := range trees {
		if trees[i][colNb] > max {
			max = trees[i][colNb]
		}
	}
	if currentTree > max {
		println(true)
	}
	return currentTree > max
}

func checkAdjacentTrees(adjacentTrees []int, currentTree int) bool {
	max := slices.Max(adjacentTrees)
	return currentTree == max
}
