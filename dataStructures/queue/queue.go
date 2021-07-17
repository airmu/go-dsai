package queue

import (
	"errors"
	"fmt"
)

// FIFO队列

type node struct {
	data interface{}
	next *node
}

type Queue struct {
	head *node
	tail *node
	len  int
}

func TestQueue() {
	q := initQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Printf("size -> %d\r\n", q.len)
	q.Print()
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}

func initQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(data interface{}) {
	if q.len == 0 {
		nd := &node{data: data}
		q.head = nd
		q.tail = nd
		q.len++
		return
	}

	nd := &node{data: data}
	q.tail.next = nd
	q.tail = nd
	q.len++
}

// 出队,空队列时返回nil
func (q *Queue) Pop() (interface{}, error) {
	if q.len == 0 {
		return nil, errors.New("队列为空")
	}
	nd := q.head.data
	q.head = q.head.next
	q.len--
	return nd, nil
}

func (q *Queue) Size() int {
	return q.len
}

func (q *Queue) Print() {
	if q.len == 0 {
		fmt.Println("队列为空")
		return
	}

	nd := q.head
	for {
		fmt.Printf(" %v", nd.data)
		nd = nd.next
		if nd == nil {
			fmt.Println("")
			return
		}
	}
}
