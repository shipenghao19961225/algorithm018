学习笔记

## 刷题表



|                   | 1     | 2     | 3    | 4    | 5    |
| ----------------- | ----- | ----- | ---- | ---- | ---- |
| 208. 实现Trie     | 11.29 | 11.29 | 12.6 | 12.6 |      |
| 78. 单词搜索      | 11.30 | 11.30 | 12.6 | 12.6 |      |
| 212. 单词搜索ii   | 11.30 | 11.30 | 12.6 | 12.6 |      |
| 547. 朋友圈       | 12.2  | 12.2  | 12.6 | 12.6 |      |
| 200. 岛屿数量     | 12.2  | 12.2  | 12.3 | 12.6 |      |
| 130. 被包围的区域 | 12.3  | 12.4  | 12.6 |      |      |
| 22.括号生成       | 12.4  | 12.4  | 12.6 |      |      |
| 70. 爬楼梯        | 12.6  |       |      |      |      |
| 36. 有效的数独    | 12.6  | 12.6  |      |      |      |
| 37. 解数独        | 12.6  | 12.6  |      |      |      |
| 127. 单词接龙     | 12.6  | 12.6  |      |      |      |
| 433. 最小基因变化 | 12.6  | 12.6  |      |      |      |

## 单词搜索时间复杂度分析

### dfs



如果是对于每个单词在每一个格都通过dfs搜索的方式，时间复杂度为

M * N * 4 ^ k 

M 代表单元格数，N代表单词数， k代表单词的平均长度



### Trie

因为已经构建了Trier树了，所以不需要每个单词都去找一遍，只要遍历单元格就可以了

时间复杂度为

M * 4  ^ k

M代表单元格数， k代表单词的平均长度

## 双向bfs模板

```go

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
```





## AVL TREE

### 旋转操作

- 左旋
- 右旋
- 左右旋
- 右左旋

