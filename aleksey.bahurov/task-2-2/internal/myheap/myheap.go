package myheap

type MyHeap interface {
	Push(value int)
	Pop() (int, error)
	Peek() (int, error)
	Len() int
	Less(i, j int) (bool, error)
	Swap(i, j int) error
}
