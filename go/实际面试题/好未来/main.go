package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	len := len(old)
	res := old[len-1]
	*h = old[:len-1]
	return res
}

func topK(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}
	h := IntHeap{}
	heap.Init(&h)
	for _, num := range nums {
		heap.Push(&h, num)
		if h.Len() > k {
			heap.Pop(&h)
		}
	}
	return h[0]
}

func main() {
	nums := []int{1, 3, 6, 4, 2, 8, 9}
	k := 3
	fmt.Printf("topK(nums, k): %v\n", topK(nums, k))
}
