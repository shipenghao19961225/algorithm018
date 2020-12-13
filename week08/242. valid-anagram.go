package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		m[t[i]]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

// 用数组速度快一点
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make([]int, 26)
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
		m[t[i]-'a']--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
