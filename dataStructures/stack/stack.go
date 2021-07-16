package stack

import (
	"errors"
	"fmt"
)

func TestStack() {
	s := initStack(10)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Print()
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

type Stack struct {
	datas []interface{}
}

func initStack(cap int) *Stack {
	return &Stack{make([]interface{}, 0, cap)}
}

func (s *Stack) Push(data interface{}) {
	s.datas = append(s.datas, data)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.datas) == 0 {
		return nil, errors.New("栈为空")
	}
	data := s.datas[len(s.datas)-1]
	s.datas = s.datas[:len(s.datas)-1]
	return data, nil
}

func (s *Stack) Print() {
	fmt.Println(s.datas...)
}
