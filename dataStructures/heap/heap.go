package heap

import (
	"errors"
	"fmt"
)

// 最大值堆, 第一版，有缺陷
type Heap struct {
	datas []int
}

func TestHeap() {
	h := initHeap()
	h.AddTail(10)
	h.AddTail(7)
	h.AddTail(9)
	h.AddTail(5)
	h.AddTail(1)
	h.AddTail(2)
	h.AddTail(8)
	h.AddTail(11)
	h.print()
}

func initHeap() *Heap {
	return &Heap{make([]int, 0, 16)}
}

// 移除最后一个节点
func (h *Heap) Remove() (int, error) {
	if len(h.datas) == 0 {
		return -1, errors.New("堆为空")
	}
	data := h.datas[len(h.datas)-1]
	h.datas = h.datas[:len(h.datas)-1]
	return data, nil
}

func (h *Heap) AddTail(data int) {
	h.datas = append(h.datas, data)
	h.Traverse(0)
}

// 移除第一个节点
func (h *Heap) RemoveHead() (int, error) {
	if len(h.datas) == 0 {
		return -1, errors.New("堆为空")
	}
	tail := h.datas[len(h.datas)-1]
	head := h.datas[0]
	h.datas[0] = tail
	h.datas = h.datas[:len(h.datas)-1]
	// todo shiftUp
	h.Traverse(0)
	return head, nil
}

func (h *Heap) Traverse(i int) {
	if 2*i+1 >= len(h.datas) {
		return
	}
	if 2*i+2 >= len(h.datas) {
		if h.datas[2*i+1] > h.datas[i] {
			da := h.datas[i]
			h.datas[i] = h.datas[2*i+1]
			h.datas[2*i+1] = da
		}
		return
	}
	h.Traverse(2*i + 1)
	h.Traverse(2*i + 2)

	cur := h.datas[i]
	left := h.datas[2*i+1]
	right := h.datas[2*i+2]

	var maxIndex int
	var max int
	if left > right {
		max = left
		maxIndex = 1
	} else {
		max = right
		maxIndex = 2
	}
	if max > cur {
		h.datas[i] = max
		h.datas[2*i+maxIndex] = cur
	}
}

// 移除元素后，重新调整堆
func (h *Heap) shiftUp() {

}

// 移除元素后，重新调整堆
func (h *Heap) shiftDown() {

}

func (h *Heap) print() {
	fmt.Println("")
	for i := 0; i < len(h.datas); i++ {
		fmt.Printf(" %d", h.datas[i])
	}
	fmt.Println("")
}
