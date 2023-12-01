package day8

import (
	"strings"

	"main.go/util"
)

func Day8() {
	treeTxt := strings.Split(string(util.ReadFile("day8/trees.txt")), "\n")
	// Transform tree into an double array of int
	treesHeights := [][]int{}
	for _, tree := range treeTxt {
		treeRows, _ := util.ConvertStringsToInts(strings.Split(tree, ""))
		treesHeights = append(treesHeights, treeRows)
	}
	countVisibleTrees := 0

	for row, trees := range treesHeights {
		if row == 0 || row == len(treesHeights)-1 {
			countVisibleTrees += len(trees)
			continue
		}
		for col, tree := range trees {
			// Visible trees on the edges
			if col == 0 || col == len(trees)-1 {
				countVisibleTrees++
				continue
			}
			// Check top
			if checkTop(treesHeights[:row], col, tree) ||
				checkBottom(treesHeights[row+1:], col, tree) ||
				checkLeft(trees[:col], row, col, tree) ||
				checkRight(trees[col+1:], row, col, tree) {
				countVisibleTrees++
			}
		}
	}

	println("VISIBLE TREES : ", countVisibleTrees)
}

func checkTop(trees [][]int, colNb, currentTree int) bool {
	// Check top
	for _, treeRow := range trees {
		if treeRow[colNb] >= currentTree {
			return false
		}
	}
	return true
}

func checkBottom(trees [][]int, colNb, currentTree int) bool {
	// Check bottom
	for _, treeRow := range trees {
		if treeRow[colNb] >= currentTree {
			return false
		}
	}
	return true
}
func checkLeft(trees []int, rowNb, colNb, currentTree int) bool {
	// Check left
	for _, tree := range trees {
		if tree >= currentTree {
			return false
		}
	}
	return true
}

func checkRight(trees []int, rowNb, colNb, currentTree int) bool {
	// Check right
	for _, tree := range trees {
		if tree >= currentTree {
			return false
		}
	}
	return true
}
