package piscine

import (
	"math"
	"strconv"
)

func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, math.MinInt64, math.MaxInt64)
}

func isBST(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	nodeValue, err := strconv.Atoi(node.Data)
	if err != nil {
		return false
	}

	if nodeValue <= min || nodeValue >= max {
		return false
	}

	return isBST(node.Left, min, nodeValue) && isBST(node.Right, nodeValue, max)
}
