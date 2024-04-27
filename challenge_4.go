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
/**
Explanation:

**Code Block 1: Defining the TreeNode struct**
```go
type TreeNode_ struct {
    Left  *TreeNode_
    Right *TreeNode_
    Value int
}
```
This defines a struct `TreeNode_` to represent a node in a binary tree. Each node has three fields: `Left` and `Right` pointers to child nodes, and a `Value` integer.

**Code Block 2: levelOrder function**
```go
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
```
This function performs a level-order traversal of a binary tree and returns the values of the nodes in the order they are visited.

* If the `root` node is `nil`, the function returns an empty slice `[]int{}`.
* Initialize an empty slice `result` to store the node values.
* Initialize a queue `queue` with the `root` node.
* Loop until the queue is empty.
	+ Pop the front node from the queue and add its value to the `result` slice.
	+ Enqueue the left child node if it's not `nil`.
	+ Enqueue the right child node if it's not `nil`.
* Return the `result` slice.

*/
