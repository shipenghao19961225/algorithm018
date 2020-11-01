package leetcode

// 这里的用切片实现队列的写法要学习一下，每次利用size遍历和切割
func levelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		size := len(queue)
		tmp := []int{}
		for i := 0; i < size; i++ {
			for _, child := range queue[i].Children {
				queue = append(queue, child)
			}
			tmp = append(tmp, queue[i].Val)
		}
		res = append(res, tmp)
		queue = queue[size:]
	}
	return res
}
