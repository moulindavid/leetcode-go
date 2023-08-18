package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS recursive
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// DFS iterative
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(root)
	level := 0
	for queue.Len() > 0 {
		n := queue.Len()
		for i := 0; i < n; i++ {
			e := queue.Front()
			queue.Remove(e)
			node := e.Value.(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		level++
	}
	return level
}
