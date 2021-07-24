package heap

import "fmt"

type HeapNode struct {
	data interface{} // 具体存的值
	key  int         // 对象重要度
}

// 最大堆
type Heap2 struct {
	datas []*HeapNode
}

func TestHeap2() {
	h := initHeap2(16)
	h.InsertTail(&HeapNode{"A", 22})
	h.InsertTail(&HeapNode{"B", 12})
	h.InsertTail(&HeapNode{"C", 14})
	h.InsertTail(&HeapNode{"D", 6})
	h.InsertTail(&HeapNode{"E", 19})
	h.InsertTail(&HeapNode{"F", 18})
	h.InsertTail(&HeapNode{"G", 23})
	h.InsertTail(&HeapNode{"H", 27})
	h.InsertTail(&HeapNode{"I", 2})
	h.InsertTail(&HeapNode{"J", 10})
	h.InsertTail(&HeapNode{"K", 9})
	h.Print()
}

func initHeap2(cap int) *Heap2 {
	if cap <= 16 {
		cap = 16
	}
	return &Heap2{make([]*HeapNode, 0, cap)}
}

// 向堆尾插入元素
func (h *Heap2) InsertTail(nd *HeapNode) {
	h.datas = append(h.datas, nd)
	h.shiftUp(len(h.datas) - 1)
}

// 移除堆头元素
func (h *Heap2) removeHead() *HeapNode {
	if len(h.datas) == 0 {
		return nil
	}
	nd := h.datas[0]
	if len(h.datas) == 1 {
		// 长度1时移除堆头，就剩空堆，无需shiftDown操作
		h.datas = h.datas[1:]
		return nd
	}
	h.datas[0] = h.datas[len(h.datas)-1]
	h.datas = h.datas[0 : len(h.datas)-1]
	h.shiftDown(0)
	return nd
}

// 上拉
func (h *Heap2) shiftUp(i int) {
	for {
		// 上拉到顶节点，或者比父节点的key小就不再上拉
		if i <= 0 || h.datas[i].key <= h.datas[h.parentId(i)].key {
			return
		}
		nd := h.datas[h.parentId(i)]
		h.datas[h.parentId(i)] = h.datas[i]
		h.datas[i] = nd
		i = h.parentId(i)
	}
}

// 下拉
func (h *Heap2) shiftDown(i int) {
	for {
		// 下拉到当前下标为叶子节点为止
		if i >= len(h.datas)/2 {
			return
		}
		var maxIndex int //较大节点的下标
		if h.datas[h.leftChild(i)].key > h.datas[h.rightChild(i)].key {
			// 左侧子节点较大
			maxIndex = h.leftChild(i)
		} else {
			maxIndex = h.rightChild(i)
		}
		temp := h.datas[maxIndex]
		h.datas[maxIndex] = h.datas[i]
		h.datas[i] = temp
		i = maxIndex
	}
}

// 获取父节点下标
func (h *Heap2) parentId(i int) int {
	return (i - 1) / 2
}

// 获取左子节点下标
func (h *Heap2) leftChild(i int) int {
	return i*2 + 1
}

// 获取右子节点下标
func (h *Heap2) rightChild(i int) int {
	return i*2 + 2
}

func (h *Heap2) Print() {
	fmt.Println("-----")
	for i := 0; i < len(h.datas); i++ {
		fmt.Printf(" %v", h.datas[i])
	}
	fmt.Println("")
	fmt.Println("-----")
}
