package leetcode

// 解法一： 单向广度优先遍历
var dic map[string]bool
var visited map[string]bool

func ladderLength(beginWord string, endWord string, wordList []string) int {
	dic = make(map[string]bool)

	visited = make(map[string]bool)
	for _, word := range wordList {
		dic[word] = true
	}

	queue := []string{beginWord}
	visited[beginWord] = true
	step := 1
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curWord := queue[0]
			queue = queue[1:]
			if changeEveryChar(curWord, endWord, &queue) {
				return step + 1
			}
		}
		step++
	}
	return 0
}

func changeEveryChar(curWord, endWord string, queue *[]string) bool {
	charArray := []byte(curWord)
	for i := 0; i < len(charArray); i++ {
		originalChar := charArray[i]
		for char := 'a'; char <= 'z'; char++ {
			if byte(char) == originalChar {
				continue
			}
			charArray[i] = byte(char)
			newWord := string(charArray)

			if _, ok := dic[newWord]; ok {
				if newWord == endWord {
					return true
				}
				if _, ok1 := visited[newWord]; !ok1 {
					*queue = append(*queue, newWord)
					visited[newWord] = true
				}
			}
		}
		charArray[i] = originalChar
	}
	return false
}

// 双向广度优先遍历，抄袭weiwei的写法
var dic map[string]bool
var visited map[string]bool
var beginVisited map[string]bool
var endVisited map[string]bool

func ladderLength(beginWord string, endWord string, wordList []string) int {
	dic = make(map[string]bool)
	for _, word := range wordList {
		dic[word] = true
	}
	if _, ok := dic[endWord]; !ok {
		return 0
	}
	visited = make(map[string]bool)
	beginVisited = make(map[string]bool)
	beginVisited[beginWord] = true
	endVisited = make(map[string]bool)
	endVisited[endWord] = true
	step := 1
	for len(beginVisited) > 0 && len(endVisited) > 0 {
		if len(beginVisited) > len(endVisited) {
			beginVisited, endVisited = endVisited, beginVisited
		}
		newVisited := make(map[string]bool)
		for curWord, _ := range beginVisited {

			if changeEveryChar(curWord, newVisited) {
				return step + 1
			}
		}
		step++
		beginVisited = newVisited
	}
	return 0
}

func changeEveryChar(curWord string, newVisited map[string]bool) bool {
	charArray := []byte(curWord)
	for i := 0; i < len(charArray); i++ {
		originalChar := charArray[i]
		for char := 'a'; char <= 'z'; char++ {
			if byte(char) == originalChar {
				continue
			}
			charArray[i] = byte(char)
			newWord := string(charArray)

			if _, ok := dic[newWord]; ok {
				if _, ok := endVisited[newWord]; ok {
					return true
				}
				if _, ok := visited[newWord]; !ok {
					newVisited[newWord] = true
					visited[newWord] = true
				}
			}
		}
		charArray[i] = originalChar
	}
	return false
}
