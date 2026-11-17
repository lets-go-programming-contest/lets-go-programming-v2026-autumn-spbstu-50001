package minheap

import "fmt"

type MinHeap []int

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
	if i < 0 || i >= len(*h) || j < 0 || j >= len(*h) {
		panic(fmt.Sprintf("minheap.Less: index out of range [%d, %d] with length %d", i, j, len(*h)))
	}

	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Swap(i, j int) {
	if i < 0 || i >= len(*h) || j < 0 || j >= len(*h) {
		panic(fmt.Sprintf("minheap.Swap: index out of range [%d, %d] with length %d", i, j, len(*h)))
	}

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(x any) {
	value, ok := x.(int)
	if !ok {
		panic(fmt.Sprintf("minheap.Push: expected int, got %T", x))
	}

	*h = append(*h, value)
}

func (h *MinHeap) Pop() any {
	if len(*h) == 0 {
		return nil
	}

	oldHeap := *h
	oldLength := len(oldHeap)
	lastElement := oldHeap[oldLength-1]
	*h = oldHeap[0 : oldLength-1]

	return lastElement
}
