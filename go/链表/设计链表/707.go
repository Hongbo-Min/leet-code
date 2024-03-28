package main

type MyListNode struct {
	Val  int
	Next *MyListNode
}

type MyLinkedList struct {
	Head *MyListNode
	Size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Head: &MyListNode{
			Val:  -1,
			Next: nil,
		},
		Size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if this == nil || index < 0 || index >= this.Size {
		return -1
	}
	cur := this.Head.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	if this == nil {
		return
	}
	preHead := this.Head.Next
	this.Head.Next = &MyListNode{
		Val:  val,
		Next: preHead,
	}
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newListNode := &MyListNode{Val: val}
	cur := this.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newListNode
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		index = 0
	} else if index > this.Size {
		return
	}
	newListNode := &MyListNode{Val: val}
	cur := this.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	newListNode.Next = cur.Next
	cur.Next = newListNode
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	cur := this.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}
	this.Size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
}
