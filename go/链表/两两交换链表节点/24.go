package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	virtualNode := &ListNode{
		Next: head,
	}
	pre := virtualNode
	for head != nil && head.Next != nil {
		temp := head.Next.Next
		pre.Next = head.Next
		head.Next.Next = head
		head.Next = temp
		pre = head
		head = temp
	}
	return virtualNode.Next
}

func main() {
}
