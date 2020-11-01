package leetcode

// 写法和二叉树一致，便于记忆
func preorder(root *Node) []int {
	stack := []*Node{}
	res := []int{}
	if root == nil {
		return nil
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		topNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, topNode.Val)
		for i := len(topNode.Children) - 1; i >= 0; i-- {
			stack = append(stack, topNode.Children[i])
		}
	}
	return res
}
