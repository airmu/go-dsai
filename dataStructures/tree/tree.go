package tree

import "fmt"

type TreeNode struct {
	Data   interface{}
	LChild *TreeNode // 左子节点
	RChild *TreeNode // 右子节点
}

// type Tree TreeNode

func TestTree() {
	root := &TreeNode{}
	root.Data = "A"
	tb := &TreeNode{Data: "B"}
	tc := &TreeNode{Data: "C"}
	root.LChild = tb
	root.RChild = tc

	td := &TreeNode{Data: "D"}
	te := &TreeNode{Data: "E"}
	tb.LChild = td
	tb.RChild = te

	tf := &TreeNode{Data: "F"}
	tg := &TreeNode{Data: "G"}
	tc.LChild = tf
	tc.RChild = tg

	th := &TreeNode{Data: "H"}
	ti := &TreeNode{Data: "I"}
	td.LChild = th
	td.RChild = ti

	tj := &TreeNode{Data: "J"}
	te.LChild = tj

	fmt.Println("前序遍历")
	root.PreTraverse(root)
	fmt.Println()
	fmt.Println("中序遍历")
	root.MiddleTraverse(root)
	fmt.Println()
	fmt.Println("后序遍历")
	root.EndTraverse(root)
	fmt.Println()
	fmt.Println("层序遍历")
	root.FlowTraverse(root)
	fmt.Println()
}

// 前序遍历
func (t *TreeNode) PreTraverse(nd *TreeNode) {
	if nd == nil {
		return
	}
	fmt.Printf(" %v", nd.Data)
	nd.PreTraverse(nd.LChild)
	nd.PreTraverse(nd.RChild)
}

// 中序遍历
func (t *TreeNode) MiddleTraverse(nd *TreeNode) {
	if nd == nil {
		return
	}
	nd.MiddleTraverse(nd.LChild)
	fmt.Printf(" %v", nd.Data)
	nd.MiddleTraverse(nd.RChild)
}

// 后序遍历
func (t *TreeNode) EndTraverse(nd *TreeNode) {
	if nd == nil {
		return
	}
	nd.EndTraverse(nd.LChild)
	nd.EndTraverse(nd.RChild)
	fmt.Printf(" %v", nd.Data)
}

// 层次遍历
func (t *TreeNode) FlowTraverse(nd *TreeNode) {
	nds := make([]*TreeNode, 0, 20)
	if nd == nil {
		return
	}
	nds = append(nds, nd)
	fmt.Printf(" %v", nd.Data)
	headIndex := 0
	tailIndex := 0
	for {
		if nds[headIndex].LChild != nil {
			nds = append(nds, nds[headIndex].LChild)
			fmt.Printf(" %v", nds[headIndex].LChild.Data)
			tailIndex++
		}
		if nds[headIndex].RChild != nil {
			nds = append(nds, nds[headIndex].RChild)
			fmt.Printf(" %v", nds[headIndex].RChild.Data)
			tailIndex++
		}
		if headIndex >= tailIndex {
			return
		}
		headIndex++
	}
}
