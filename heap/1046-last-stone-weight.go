package heap

import (
	"container/heap"
	"fmt"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	minHeap *IntHeap
	k       int
}

func lastStoneWeight(_stones []int) int {
	for s := 0; s < len(_stones); s++ {
		_stones[s] *= -1
	}
	stones := IntHeap(_stones)
	heap.Init(&stones)

	for len(stones) > 1 {
		first, _ := heap.Pop(&stones).(int)
		second, _ := heap.Pop(&stones).(int)
		if second > first {
			heap.Push(&stones, first-second)
		}
	}
	stones = append(stones, 0)
	return abs(stones[0])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func lastStoneWeightOpt(stones []int) int {
	sort.Ints(stones)
	for len(stones) > 1 {
		fmt.Println(stones)
		e1, e2 := stones[len(stones)-1], stones[len(stones)-2]
		if e1 == e2 {
			stones = stones[:len(stones)-2]
		} else {
			stones[len(stones)-2] = stones[len(stones)-1] - stones[len(stones)-2]
			stones = stones[:len(stones)-1]
		}
		sort.Ints(stones)
	}
	if len(stones) > 0 {
		return stones[0]
	}
	return 0
}
