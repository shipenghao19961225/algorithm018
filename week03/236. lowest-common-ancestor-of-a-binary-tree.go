package leetcode

// 当我们用递归去做这个题时不要被题目误导，应该要明确一点
// 这个函数的功能有三个：给定两个节点 p 和 q
// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/solution/c-jing-dian-di-gui-si-lu-fei-chang-hao-li-jie-shi-/
// 如果 p 和 q 都存在，则返回它们的公共祖先；
// 如果只存在一个，则返回存在的一个；
// 如果 p 和 q 都不存在，则返回NULL

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	leftAncestor := lowestCommonAncestor(root.Left, p, q)
	rightAncestor := lowestCommonAncestor(root.Right, p, q)
	if leftAncestor == nil && rightAncestor == nil {
		return nil
	}
	if leftAncestor == nil {
		return rightAncestor
	}
	if rightAncestor == nil {
		return leftAncestor
	}
	return root
}
