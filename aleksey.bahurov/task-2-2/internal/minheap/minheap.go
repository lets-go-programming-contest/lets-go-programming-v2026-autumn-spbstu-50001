package minheap

import (
	"container/heap"
	"errors"
	"fmt"
)

var (
	errEmptyHeap      = errors.New("heap is empty")
	errUnexpectedType = errors.New("unexpected type in heap")
)

type rawHeap []int

func (h *rawHeap) Len() int {
	return len(*h)
}

func (h *rawHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *rawHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *rawHeap) Push(x any) {
	val, ok := x.(int)
	if !ok {
		panic(fmt.Sprintf("expected int, got %T", x))
	}

	*h = append(*h, val)
}

func (h *rawHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]

	return item
}

type minHeap struct {
	data *rawHeap
}

func NewMinHeap() *minHeap {
	h := &rawHeap{}
	heap.Init(h)

	return &minHeap{data: h}
}

func (m *minHeap) Len() int {
	return m.data.Len()
}

func (m *minHeap) Push(value int) {
	heap.Push(m.data, value)
}

func (m *minHeap) Pop() (int, error) {
	if m.data.Len() == 0 {
		return 0, errEmptyHeap
	}

	value, ok := heap.Pop(m.data).(int)
	if !ok {
		return 0, errUnexpectedType
	}

	return value, nil
}

func (m *minHeap) Peek() (int, error) {
	if m.data.Len() == 0 {
		return 0, errEmptyHeap
	}

	return (*m.data)[0], nil
}
