package helperFunc

import (
	"math"
)

type BST struct {
	value int
	left  *BST
	right *BST
}

var globalDiff, closestValue int

func (tree *BST) FindClosestValue(target int) int {

	globalDiff = math.MaxInt32
	closestValueHelper(tree, target)
	return closestValue
}

func abs(a, b int) int {

	diff := a - b
	if diff < 0 {
		return b - a
	}
	return diff
}

func closestValueHelper(root *BST, target int) {
	if root == nil {
		return
	}

	currDiff := abs(root.value, target)

	if currDiff < globalDiff {
		globalDiff = currDiff
		closestValue = root.value
	}

	if root.value > target {
		closestValueHelper(root.left, target)
	} else {
		closestValueHelper(root.right, target)
	}
}
