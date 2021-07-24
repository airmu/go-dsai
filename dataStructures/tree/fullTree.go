package tree

// 完全二叉树可以用数组存储，更省空间，相当于一个队列了
type FullTree struct {
	datas []interface{}
}

func initFullTree() *FullTree {
	return &FullTree{make([]interface{}, 0, 10)}
}

func (ft *FullTree) AddNode(data interface{}) {
	ft.datas = append(ft.datas, data)
}
