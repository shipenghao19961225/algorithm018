package leetcode

func minMutation(start string, end string, bank []string) int {
	dic := make(map[string]bool)
	for _, gene := range bank {
		dic[gene] = true
	}
	if _, ok := dic[end]; !ok {
		return -1
	}
	singleGroup := []byte{'A', 'G', 'C', 'T'}
	visited := make(map[string]bool)
	beginVisited := make(map[string]bool)
	beginVisited[start] = true
	endVisited := make(map[string]bool)
	endVisited[end] = true
	step := 0
	for len(beginVisited) > 0 && len(endVisited) > 0 {
		if len(beginVisited) > len(endVisited) {
			beginVisited, endVisited = endVisited, beginVisited
		}
		newVisited := make(map[string]bool)
		for curGene, _ := range beginVisited {
			// "AGCT"
			charArray := []byte(curGene)
			for i := 0; i < len(charArray); i++ {
				//AGCT -> GGCT
				oriChar := charArray[i]
				for _, char := range singleGroup {
					if char == oriChar {
						continue
					}
					charArray[i] = char
					newGene := string(charArray)
					if _, ok := dic[newGene]; ok {
						if _, ok := endVisited[newGene]; ok {
							return step + 1
						}
						if _, ok := visited[newGene]; !ok {
							newVisited[newGene] = true
							visited[newGene] = true
						}
					}
				}
				charArray[i] = oriChar
			}

		}
		step++
		beginVisited = newVisited
	}
	return -1
}


// 双向bfs模板

//
func () {
	visited := make(map[string]bool)
	beginVisited, endVisited := make(map[string]bool),make(map[string]bool)
	for len(beginVisited) > 0 && len(endVisited) > 0 {
		if len(beginVisited) > len(endVisited) {
			beginVisited, endVisited = endVisited, beginVisited
		}
		newVisited := map[string]bool
		for Ele in beginVisited {
			// process Ele, if processed Ele meet in endVisited 
			// return
			// else add the satisfied element to the newVisited
			}

		}
		beginVisited = newVisited
	} 
	 
}