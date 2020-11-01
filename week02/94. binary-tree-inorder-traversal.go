package leetcode

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/die-dai-fa-by-jason-2/讲解的比较好
// 记住这个写法，慢慢理解
func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
		stack = stack[:len(stack)-1]
	}
	return res

}
