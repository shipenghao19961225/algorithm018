package leetcode

func groupAnagrams(strs []string) [][]string {
	// 关键点是用长度为26的数组作为键
	m := make(map[[26]int][]string)
	res := [][]string{}
	for _, str := range strs {
		Key := getKey(str)
		m[Key] = append(m[Key], str)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res

}

func getKey(s string) (array [26]int) {
	for i := 0; i < len(s); i++ {
		array[s[i]-'a']++
	}
	return
}
