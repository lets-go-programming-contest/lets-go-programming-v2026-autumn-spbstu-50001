package minheap

import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/Alehandro-de-bah/task-2-2/internal/myheap"
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
	*h = append(*h, x.(int))
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

func NewMinHeap() myheap.MyHeap {
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

func (m *minHeap) checkBounds(i, j int) error {
	if i < 0 || i >= m.Len() || j < 0 || j >= m.Len() {
		return fmt.Errorf("index out of range [%d, %d] with length %d", i, j, m.Len())
	}

	return nil
}

func (m *minHeap) Less(i, j int) (bool, error) {
	if err := m.checkBounds(i, j); err != nil {
		return false, fmt.Errorf("minheap.Less: %w", err)
	}

	return m.data.Less(i, j), nil
}

func (m *minHeap) Swap(i, j int) error {
	if err := m.checkBounds(i, j); err != nil {
		return fmt.Errorf("minheap.Swap: %w", err)
	}

	m.data.Swap(i, j)

	return nil
}
