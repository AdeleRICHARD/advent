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
			if checkTopAndBottom(treesHeights, currentTree, i, j) &&
				checkAdjacentTrees(treesHeights[i], currentTree) {

				countVisibleTrees++
			}

		}
	}
	println("VISIBLE TREES : ", countVisibleTrees)
}

func checkTopAndBottom(trees [][]int, currentTree, rowNb, colNb int) bool {
	for i := range trees {
		nb := trees[i][colNb]
		if nb > currentTree && i != rowNb {
			return false
		}

	}
	return true
}

func checkAdjacentTrees(adjacentTrees []int, currentTree int) bool {

	return currentTree > slices.Max(adjacentTrees)
}
