package linked

import (
	"errors"
	"fmt"
)

type node struct {
	data interface{}
	next *node
}

// 单向链表栈
type DxLinked struct {
	current *node
	len     int
}

func TestDxLinked() {
	dl := initLinked()
	dl.Push(1)
	dl.Push(2)
	dl.Push(3)
	dl.Push(4)
	dl.Print()
	fmt.Println(dl.Pop())
	fmt.Println(dl.Pop())
	fmt.Println(dl.Pop())
	fmt.Println(dl.Pop())
	fmt.Println(dl.Pop())
	fmt.Println(dl.Pop())
}

func initLinked() *DxLinked {
	return &DxLinked{}
}

func (link *DxLinked) Push(data interface{}) {
	if link.len == 0 {
		link.current = &node{data: data}
		link.len++
		return
	}

	nd := &node{data: data}
	nd.next = link.current
	link.current = nd
	link.len++
}

func (link *DxLinked) Pop() (interface{}, error) {
	if link.len == 0 {
		return nil, errors.New("链表为空")
	}

	cur := link.current
	link.current = link.current.next
	link.len--
	return cur, nil
}

func (link *DxLinked) Print() {
	if link.len <= 0 {
		return
	}

	var cur *node
	cur = link.current
	for {
		fmt.Printf(" %v", cur.data)
		cur = cur.next
		if cur == nil {
			fmt.Println("")
			return
		}
	}
}
