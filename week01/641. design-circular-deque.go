package leetcode

type DLinkedNode struct {
	Val  int
	Pre  *DLinkedNode
	Next *DLinkedNode
}

type MyCircularDeque struct {
	Size   int
	Length int
	Head   *DLinkedNode
	Tail   *DLinkedNode
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	head := &DLinkedNode{
		-1,
		nil,
		nil,
	}
	tail := &DLinkedNode{
		-1,
		nil,
		nil,
	}
	head.Next = tail
	tail.Pre = head
	return MyCircularDeque{
		k,
		0,
		head,
		tail,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	newNode := new(DLinkedNode)
	newNode.Val = value
	tmp := this.Head.Next
	tmp.Pre = newNode
	newNode.Next = tmp
	this.Head.Next = newNode
	newNode.Pre = this.Head
	this.Length++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	newNode := new(DLinkedNode)
	newNode.Val = value
	tmp := this.Tail.Pre
	tmp.Next = newNode
	newNode.Pre = tmp
	this.Tail.Pre = newNode
	newNode.Next = this.Tail
	this.Length++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	tmp := this.Head.Next
	this.Head.Next = tmp.Next
	tmp.Next.Pre = this.Head
	tmp.Next = nil
	tmp.Pre = nil
	this.Length--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	tmp := this.Tail.Pre
	this.Tail.Pre = tmp.Pre
	tmp.Pre.Next = this.Tail
	tmp.Next = nil
	tmp.Pre = nil
	this.Length--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Head.Next.Val
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Tail.Pre.Val
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.Length == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.Length == this.Size
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
