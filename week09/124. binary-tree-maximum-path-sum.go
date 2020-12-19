package leetcode

import "math"

var res int

func maxPathSum(root *TreeNode) int {
	res = math.MinInt32
	dfs(root)
	return res
}

// 递归的含义代表该所有该节点的子节点到该节点的路径最大值（包含该节点）
func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := max(0, dfs(root.Left))
	rightMax := max(0, dfs(root.Right))
	curMax := leftMax + root.Val + rightMax
	res = max(res, curMax)
	tmp := max(leftMax, rightMax) + root.Val
	return tmp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
