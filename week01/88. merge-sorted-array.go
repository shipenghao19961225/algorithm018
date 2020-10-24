package leetcode

// 技巧是在第一个数组，第二个数组设置指针，在第一个数组末尾设置指针进行操作
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p3 := m-1, n-1, m+n-1
	for p2 >= 0 {
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p3] = nums1[p1]
			p1--

		} else {
			nums1[p3] = nums2[p2]
			p2--
		}
		p3--
	}
}
