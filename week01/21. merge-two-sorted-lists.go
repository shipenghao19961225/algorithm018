package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归解法
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

//迭代解法

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	yummy := &ListNode{}
	tmp := yummy
	p1, p2 := l1, l2
	for {
		if p1 == nil {
			tmp.Next = p2
			break
		} else if p2 == nil {
			tmp.Next = p1
			break
		} else {
			if p1.Val <= p2.Val {
				tmp.Next = p1
				p1 = p1.Next

			} else {
				tmp.Next = p2
				p2 = p2.Next
			}
			tmp = tmp.Next
		}

	}
	return yummy.Next
}
