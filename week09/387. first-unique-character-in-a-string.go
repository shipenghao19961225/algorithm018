package leetcode

func firstUniqChar(s string) int {
	var m [128]int
	for _, char := range s {
		m[char]++
	}
	for i, char := range s {
		if m[char] == 1 {
			return i
		}
	}
	return -1
}
