package main

import (
	"math"
)

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

// maxPathSumFromRootToLeaf returns the maximum sum of a route from root to leaf
func maxPathSumFromRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxSum := math.MinInt64 // Initialize maxSum with the minimum integer value
	maxPathSumUtil(root, 0, &maxSum)
	return maxSum
}

func maxPathSumUtil(node *TreeNode, sum int, maxSum *int) {
	if node == nil {
		return
	}

	// Calculate the sum till this node
	sum += node.Value

	// If this node is a leaf node, update maxSum if necessary
	if node.Left == nil && node.Right == nil {
		if sum > *maxSum {
			*maxSum = sum
		}
		return
	}

	// Recursively calculate the maximum path sum for left and right subtrees
	maxPathSumUtil(node.Left, sum, maxSum)
	maxPathSumUtil(node.Right, sum, maxSum)
}
