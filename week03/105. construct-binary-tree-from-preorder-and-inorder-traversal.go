package leetcode

func buildTree(preorder []int, inorder []int) *TreeNode {
	// termination
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// process
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	// drill down
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:1+i], inorder[:i]),
		Right: buildTree(preorder[1+i:], inorder[i+1:]),
	}
}
