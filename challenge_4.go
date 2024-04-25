package main

type TreeNode_ struct {
	Left  *TreeNode_
	Right *TreeNode_
	Value int
}

func levelOrder(root *TreeNode_) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	queue := []*TreeNode_{root}

	for len(queue) > 0 {
		// Pop the front node
		node := queue[0]
		queue = queue[1:]

		// Add node value to the result
		result = append(result, node.Value)

		// Enqueue left child
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		// Enqueue right child
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}
