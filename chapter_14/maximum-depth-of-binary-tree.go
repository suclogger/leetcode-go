package main

import "math"

func maxDepth(root *TreeNode) int {
	return visit(root, 0)
}

func visit(root *TreeNode, dep int) int {
	if root == nil {
		return dep
	}
	return int(math.Max(float64(visit(root.Left, dep+1)), float64(visit(root.Right, dep+1))))
}
