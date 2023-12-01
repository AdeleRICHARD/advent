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
	countScenicCone := 0

	scenicCone := 0

	for row, trees := range treesHeights {
		// Part 1
		/* if row == 0 || row == len(treesHeights)-1 {
			countVisibleTrees += len(trees)
		} */
		for col, tree := range trees {
			// Visible trees on the edges
			// PART 1
			/* if col == 0 || col == len(trees)-1 {
				countVisibleTrees++
			} */

			// PART 2
			if row == 0 {
				scenicCone = countBottomTrees(treesHeights[row+1:], col, tree)
				scenicCone *= countLeftTrees(trees[:col], row, col, tree)
				scenicCone *= countRightTrees(trees[col+1:], row, col, tree)
			} else if row == len(treesHeights)-1 {
				scenicCone = countTopTrees(treesHeights[:row], col, tree)
				scenicCone *= countLeftTrees(trees[:col], row, col, tree)
				scenicCone *= countRightTrees(trees[col+1:], row, col, tree)
			} else if col == 0 {
				scenicCone = countTopTrees(treesHeights[:row], col, tree)
				scenicCone *= countBottomTrees(treesHeights[row+1:], col, tree)
				scenicCone *= countRightTrees(trees[col+1:], row, col, tree)
			} else if col == len(trees)-1 {
				scenicCone = countTopTrees(treesHeights[:row], col, tree)
				scenicCone *= countBottomTrees(treesHeights[row+1:], col, tree)
				scenicCone *= countLeftTrees(trees[:col], row, col, tree)
			} else {
				scenicCone = countTopTrees(treesHeights[:row], col, tree)
				scenicCone *= countBottomTrees(treesHeights[row+1:], col, tree)
				scenicCone *= countLeftTrees(trees[:col], row, col, tree)
				scenicCone *= countRightTrees(trees[col+1:], row, col, tree)
			}
			// Part 1
			/* if checkTop(treesHeights[:row], col, tree) ||
				checkBottom(treesHeights[row+1:], col, tree) ||
				checkLeft(trees[:col], row, col, tree) ||
				checkRight(trees[col+1:], row, col, tree) {
				countVisibleTrees++
			} */
			// Part 2
			if countScenicCone < scenicCone {
				countScenicCone = scenicCone
			}

		}
	}

	println("VISIBLE TREES : ", countVisibleTrees)
	println("SCENIC CONE : ", countScenicCone)
}

// ************** Part 2 ****************

func countTopTrees(trees [][]int, colNb, currentTree int) int {
	// Check top
	smallerTrees := 0
	for i := len(trees) - 1; i >= 0; i-- {
		treeRow := trees[i]
		if treeRow[colNb] >= currentTree {
			smallerTrees++
			break
		}
		smallerTrees++
	}

	return smallerTrees
}

func countBottomTrees(trees [][]int, colNb, currentTree int) int {
	// Check bottom
	smallerTrees := 0
	for _, treeRow := range trees {
		if treeRow[colNb] >= currentTree {
			smallerTrees++
			break
		}
		smallerTrees++
	}

	return smallerTrees
}
func countLeftTrees(trees []int, rowNb, colNb, currentTree int) int {
	// Check left
	smallerTrees := 0
	for i := len(trees) - 1; i >= 0; i-- {
		tree := trees[i]
		if tree >= currentTree {
			smallerTrees++
			break
		}
		smallerTrees++
	}
	return smallerTrees
}

func countRightTrees(trees []int, rowNb, colNb, currentTree int) int {
	// Check right
	smallerTrees := 0
	for _, tree := range trees {
		if tree >= currentTree {
			smallerTrees++
			break
		}
		smallerTrees++
	}

	return smallerTrees
}

// ***************** Part 1 *****************

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
