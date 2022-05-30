package heap

import "math"

type MinHeap struct {
	Size  int
	Elems []int
}

func NewMinHeap(maxSize int) *MinHeap {
	return &MinHeap{
		Size:  -1,
		Elems: make([]int, maxSize),
	}
}

func (h *MinHeap) Push(v int) {
	if h.Size == len(h.Elems)-1 {
		return
	}

	h.Size++
	i := h.Size
	// 与父节点进行比较，如果小于父节点，则向上冒泡；否则，停止；
	for i > 0 {
		parent := (i - 1) / 2
		if h.Elems[parent] <= v {
			break
		}
		h.Elems[i] = h.Elems[parent]
		i = parent
	}
	h.Elems[i] = v
}

func (h *MinHeap) Pop() int {
	if h.Size < 0 {
		return math.MinInt32
	}

	ret := h.Elems[0]
	// 将最后一个节点提到根节点，然后向下交换
	x := h.Elems[h.Size]
	i := 0
	for {
		l := 2*i + 1
		r := 2*i + 2

		if l >= h.Size {
			break
		}

		if r < h.Size && h.Elems[l] > h.Elems[r] {
			l = r
		}

		if x <= h.Elems[l] {
			break
		}

		h.Elems[i] = h.Elems[l]
		i = l
	}

	h.Size--
	h.Elems[i] = x

	return ret
}
