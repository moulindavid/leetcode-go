package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxLength := 0
	dfs(root, &maxLength)
	return maxLength
}

func dfs(node *TreeNode, maxLength *int) int {
	if node == nil {
		return 0
	}

	left := dfs(node.Left, maxLength)
	right := dfs(node.Right, maxLength)
	*maxLength = max(*maxLength, left+right)

	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
