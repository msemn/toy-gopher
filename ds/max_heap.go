package main

import "fmt"

type MaxHeap struct {
	array []int
}

func NewMaxHeap(inputArray []int) *MaxHeap {
	newHeap := &MaxHeap{}
	for _, v := range inputArray {
		newHeap.Insert(v)
	}
	return newHeap
}

func (h *MaxHeap) Insert(value int) {
	h.array = append(h.array, value)
	if h.last() != 0 {
		h.siftUp(h.last())
	}
}

func (h *MaxHeap) Max() int {
	if len(h.array) == 0 {
		return -1
	}

	result := h.array[0]
	h.array[0] = h.array[h.last()]
	h.array = h.array[:len(h.array)-1]
	if len(h.array) > 1 {
		h.siftDown()
	}
	return result
}

func (h *MaxHeap) siftUp(i int) {
	for i > 0 && h.array[i] > h.array[h.parent(i)] {
		h.swap(i, h.parent(i))
		i = h.parent(i)
	}
}

func (h *MaxHeap) siftDown() {
	i := 0
	for i < h.last() {
		if h.left(i) < 0 && h.right(i) < 0 {
			return
		}

		if h.right(i) > 0 && h.array[h.left(i)] > h.array[h.right(i)] {
			if h.array[i] < h.array[h.left(i)] {
				h.swap(i, h.left(i))
				i = h.left(i)
			} else {
				return
			}
		} else if h.right(i) > 0 && h.array[h.left(i)] < h.array[h.right(i)] {
			if h.array[i] < h.array[h.right(i)] {
				h.swap(i, h.right(i))
				i = h.right(i)
			} else {
				return
			}
		} else {
			return
		}
	}
}

func (h *MaxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MaxHeap) left(i int) int {
	left := (i * 2) + 1
	if left > h.last() {
		return -1
	}
	return left
}

func (h *MaxHeap) right(i int) int {
	right := (i * 2) + 2
	if right > h.last() {
		return -1
	}
	return right
}

func (h *MaxHeap) last() int {
	return len(h.array) - 1
}

func (h *MaxHeap) String() string {
	return fmt.Sprintf("%q", h.array)
}

func (h *MaxHeap) swap(i1 int, i2 int){
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	//inputArray := []int{12, 5, 1, 8, 15, 2, 7, 3, 4, 10}
	heap := &MaxHeap{}
	heap.Insert(12)
	heap.Insert(5)
	heap.Insert(1)
	heap.Insert(8)
	heap.Insert(15)
	heap.Insert(2)

	for i := 0; i < 8; i++ {
		fmt.Printf("%d, ", heap.array)
		fmt.Printf("My Max is: %d\n", heap.Max())
	}
}
