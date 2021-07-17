package linked

import (
	"errors"
	"fmt"
)

// 双向链表

type dnode struct {
	data interface{}
	pre  *dnode
	next *dnode
}

type DLinked struct {
	head *dnode
	tail *dnode
	len  int
}

func TestDLinked() {
	dl := initDLinked()
	dl.AddHead(1)
	dl.AddHead(2)
	dl.AddTail(nil)
	dl.AddTail(4)
	dl.AddHead(5)
	dl.AddTail(6)
	dl.Print()
	fmt.Println(dl.RemoveHead())
	fmt.Println(dl.RemoveHead())
	fmt.Println(dl.RemoveTail())
	fmt.Println(dl.RemoveHead())
	fmt.Println(dl.RemoveTail())
	fmt.Println(dl.RemoveHead())
	fmt.Println(dl.RemoveTail())
	fmt.Println(dl.RemoveHead())
	fmt.Println(dl.RemoveTail())
	fmt.Println(dl.RemoveHead())
}

func initDLinked() *DLinked {
	return &DLinked{}
}

// 在尾部添加
func (dl *DLinked) AddTail(data interface{}) {
	dnd := &dnode{data: data}
	if dl.len == 0 {
		dl.head = dnd
		dl.tail = dnd
		dl.len++
		return
	}

	dl.tail.next = dnd
	dnd.pre = dl.tail
	dl.tail = dnd
	dl.len++
}

// 在头部添加
func (dl *DLinked) AddHead(data interface{}) {
	dnd := &dnode{data: data}
	if dl.len == 0 {
		dl.head = dnd
		dl.tail = dnd
		dl.len++
		return
	}
	dnd.next = dl.head
	dl.head.pre = dnd
	dl.head = dnd
	dl.len++
}

func (dl *DLinked) RemoveTail() (interface{}, error) {
	if dl.len == 0 {
		return nil, errors.New("链表为空")
	}

	tail := dl.tail.data
	dl.tail = dl.tail.pre
	dl.len--
	return tail, nil
}

func (dl *DLinked) RemoveHead() (interface{}, error) {
	if dl.len == 0 {
		return nil, errors.New("链表为空")
	}

	head := dl.head.data
	dl.head = dl.head.next
	dl.len--
	return head, nil
}

func (dl *DLinked) Print() {
	if dl.len == 0 {
		fmt.Println("链表为空")
		return
	}

	dnd := dl.head
	for {
		fmt.Printf(" %v", dnd.data)
		dnd = dnd.next
		if dnd == nil {
			fmt.Println("")
			return
		}
	}
}
