package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root.Right)
			res = append(res, root.Val)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return res
}

// 这里偏向于第二种解法，因为更具有通用性

func preorderTraversal2(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	if root == nil {
		return res
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		topNode := stack[len(stack)-1]
		res = append(res, topNode.Val)
		stack = stack[:len(stack)-1]
		if topNode.Right != nil {
			stack = append(stack, topNode.Right)
		}
		if topNode.Left != nil {
			stack = append(stack, topNode.Left)
		}
	}
	return res
}
