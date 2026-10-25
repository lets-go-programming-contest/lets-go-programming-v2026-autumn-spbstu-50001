package main

import (
	"container/heap"
	"errors"
	"fmt"
)

var (
	ErrDishesCount     = errors.New("invalid dishes count")
	ErrPreferenceCount = errors.New("invalid preference number")
)

type minHeap []int

func (h *minHeap) Len() int {
	return len(*h)
}

func (h *minHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *minHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *minHeap) Push(x any) {
	value, ok := x.(int)
	if !ok {
		return
	}

	*h = append(*h, value)
}

func (h *minHeap) Pop() any {
	oldHeap := *h
	oldLength := len(oldHeap)
	lastElement := oldHeap[oldLength-1]
	*h = oldHeap[0 : oldLength-1]

	return lastElement
}

func main() {
	var dishesCount int
	if _, err := fmt.Scan(&dishesCount); err != nil || dishesCount < 1 {
		fmt.Println("Error:", ErrDishesCount)

		return
	}

	ratings := make([]int, dishesCount)
	for i := 0; i < dishesCount; i++ {
		if _, err := fmt.Scan(&ratings[i]); err != nil {
			fmt.Println("Error: invalid rating format")

			return
		}
	}

	var preference int
	if _, err := fmt.Scan(&preference); err != nil || preference < 1 || preference > dishesCount {
		fmt.Println("Error:", ErrPreferenceCount)

		return
	}

	dishHeap := &minHeap{}
	heap.Init(dishHeap)

	for _, rating := range ratings {
		heap.Push(dishHeap, rating)

		if dishHeap.Len() > preference {
			heap.Pop(dishHeap)
		}
	}

	fmt.Println((*dishHeap)[0])
}
