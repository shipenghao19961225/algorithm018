package leetcode

type Trie struct {
	Next [26]*Trie
	Word string
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, char := range word {
		if node.Next[char-'a'] == nil {
			node.Next[char-'a'] = new(Trie)
		}
		node = node.Next[char-'a']
	}
	node.Word = word
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, char := range word {
		if node.Next[char-'a'] == nil {
			return false
		}
		node = node.Next[char-'a']
	}
	return node.Word == word
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, char := range prefix {
		if node.Next[char-'a'] == nil {
			return false
		}
		node = node.Next[char-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
