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

/**
Explanation:

**Code Block 1: Defining the TreeNode struct**
```go
type TreeNode struct {
    Left  *TreeNode
    Right *TreeNode
    Value int
}
```
This defines a struct `TreeNode` to represent a node in a binary tree. Each node has three fields: `Left` and `Right` pointers to child nodes, and a `Value` integer.

**Code Block 2: maxPathSumFromRootToLeaf function**
```go
func maxPathSumFromRootToLeaf(root *TreeNode) int {
    if root == nil {
        return 0
    }
    maxSum := math.MinInt64 // Initialize maxSum with the minimum integer value
    maxPathSumUtil(root, 0, &maxSum)
    return maxSum
}
```
This function calculates the maximum sum of a route from the root to a leaf node in a binary tree. It takes a `root` node as input and returns the maximum sum.

* If the `root` node is `nil`, the function returns 0.
* Initialize `maxSum` with the minimum integer value using `math.MinInt64`.
* Call the `maxPathSumUtil` function to calculate the maximum sum.

**Code Block 3: maxPathSumUtil function**
```go
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
```
This function is a recursive utility function to calculate the maximum sum of a route from the root to a leaf node.

* If the `node` is `nil`, the function returns immediately.
* Calculate the sum till this node by adding the `Value` of the current node to the `sum`.
* If the current node is a leaf node (i.e., both `Left` and `Right` are `nil`), update `maxSum` if the current sum is greater than the current `maxSum`.
* Recursively call `maxPathSumUtil` for the left and right subtrees.

*/
