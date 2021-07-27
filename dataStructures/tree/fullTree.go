package tree

import "fmt"

// 完全二叉树可以用数组存储，更省空间，相当于一个队列了
type FullTree struct {
	datas []interface{}
}

func TestFullTree() {
	ft := initFullTree()
	ft.AddNode(1)
	ft.AddNode(2)
	ft.AddNode(3)
	ft.AddNode(4)
	ft.AddNode(5)
	ft.AddNode(6)
	ft.AddNode(7)
	ft.AddNode(8)
	fmt.Println("PreTraverse --- start")
	ft.PreTraverse(0)
	fmt.Println("\r\nPreTraverse --- end")
	fmt.Println("MiddleTraverse --- start")
	ft.MiddleTraverse(0)
	fmt.Println("\r\nMiddleTraverse --- end")
	fmt.Println("Traverse --- start")
	ft.EndTraverse(0)
	fmt.Println("\r\nTraverse --- end")
}

func initFullTree() *FullTree {
	return &FullTree{make([]interface{}, 0, 10)}
}

func (ft *FullTree) AddNode(data interface{}) {
	ft.datas = append(ft.datas, data)
}

// 前序遍历
func (ft *FullTree) PreTraverse(i int) {
	if i >= len(ft.datas) {
		return
	}
	fmt.Printf(" %v", ft.datas[i])
	ft.PreTraverse(2*i + 1)
	ft.PreTraverse(2*i + 2)
}

// 中序遍历
func (ft *FullTree) MiddleTraverse(i int) {
	if i >= len(ft.datas) {
		return
	}
	ft.MiddleTraverse(2*i + 1)
	fmt.Printf(" %v", ft.datas[i])
	ft.MiddleTraverse(2*i + 2)
}

// 后序遍历
func (ft *FullTree) EndTraverse(i int) {
	if i >= len(ft.datas) {
		return
	}
	ft.EndTraverse(2*i + 1)
	ft.EndTraverse(2*i + 2)
	fmt.Printf(" %v", ft.datas[i])
}

// 层次遍历
func (ft *FullTree) FlowTraverse(i int) {
	for i < len(ft.datas) {
		fmt.Printf(" %v", ft.datas[i])
	}
}
