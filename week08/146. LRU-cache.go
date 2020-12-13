package leetcode

type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Next *Node
}
type LRUCache struct {
	Cap  int
	Map  map[int]*Node
	Head *Node
	Tail *Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Cap:  capacity,
		Map:  make(map[int]*Node, capacity),
		Head: &Node{},
		Tail: &Node{},
	}
	cache.Head.Next = cache.Tail
	cache.Tail.Pre = cache.Head
	return cache
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Map[key]
	if !ok {
		return -1
	}
	this.remove(node)
	this.setHeader(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.Map[key]
	if ok {
		this.remove(node)
		node.Val = value
		this.setHeader(node)
	} else {
		if len(this.Map) == this.Cap {
			delete(this.Map, this.Tail.Pre.Key)
			this.remove(this.Tail.Pre)
		}
		node = &Node{
			Key: key,
			Val: value,
		}
		this.setHeader(node)
		this.Map[key] = node
	}
}

func (this *LRUCache) setHeader(node *Node) {
	this.Head.Next.Pre = node
	node.Next = this.Head.Next
	this.Head.Next = node
	node.Pre = this.Head
}

func (this *LRUCache) remove(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
