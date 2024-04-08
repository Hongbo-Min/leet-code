package main

import (
	"container/heap"
	"fmt"
)

type myNode struct {
	Value int
	Index int
}

type myHeap []myNode

func (h myHeap) Len() int {
	return len(h)
}

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i].Value < (*h)[j].Value
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Push(x interface{}) {
	node := x.(myNode)
	*h = append(*h, node)
}

func (h *myHeap) Pop() interface{} {
	old := *h
	n := len(old)
	node := old[n-1]
	*h = old[:n-1]
	return node
}

func topK(nums []int, k int) int {
	h := &myHeap{}
	heap.Init(h)
	for i, v := range nums {
		heap.Push(h, myNode{Value: v, Index: i})
		if len(*h) > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = (*h)[i].Value
	}
	return res[0]
}

func main() {
	nums := []int{1, 7, 2, 8, 9, 10, 3}
	k := 3
	fmt.Printf("topK(nums, k): %v\n", topK(nums, k))
}
