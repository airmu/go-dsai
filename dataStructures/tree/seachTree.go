package tree

import "fmt"

type SeachTreeNode struct {
	data       interface{}
	leftChild  *SeachTreeNode
	rightChild *SeachTreeNode
	key        int
}

type BiSeachTree struct {
	root *SeachTreeNode
}

func TestBiSeachTree() {
	bt := initSeachTree()
	bt.Insert("A", 23)
	bt.Insert("B", 46)
	bt.Insert("C", 34)
	bt.Insert("D", 75)
	bt.Insert("E", 45)
	bt.Insert("F", 25)
	bt.Insert("G", 88)
	bt.Insert("H", 53)
	bt.Insert("I", 44)
	bt.Insert("J", 43)
	fmt.Println("===")
	bt.MiddleTranvers(bt.root)
	fmt.Println("")
	fmt.Println("===")
	bt.Remove(45)
	fmt.Println("===")
	bt.MiddleTranvers(bt.root)
	fmt.Println("")
	fmt.Println("===")
	nd := bt.Seach(75)
	if nd == nil {
		fmt.Println("未找到节点")
	} else {
		fmt.Printf("(data -> %v, key -> %v)", nd.data, nd.key)
	}
	fmt.Println("")
}

func initSeachTree() *BiSeachTree {
	return &BiSeachTree{}
}
func (bt *BiSeachTree) Seach(key int) *SeachTreeNode {
	return bt.seachNode(bt.root, key)
}

func (bt *BiSeachTree) seachNode(node *SeachTreeNode, key int) *SeachTreeNode {
	if node == nil {
		return nil
	}
	if key > node.key {
		return bt.seachNode(node.rightChild, key)
	} else if key < node.key {
		return bt.seachNode(node.leftChild, key)
	} else {
		return node
	}
}

func (bt *BiSeachTree) MiddleTranvers(node *SeachTreeNode) {
	if node == nil {
		return
	}
	bt.MiddleTranvers(node.leftChild)
	fmt.Printf(" (%v, %v)", node.data, node.key)
	bt.MiddleTranvers(node.rightChild)
}

func (bt *BiSeachTree) Insert(data interface{}, key int) {
	bt.root = bt.insertNode(bt.root, data, key)
}
func (bt *BiSeachTree) insertNode(node *SeachTreeNode, data interface{}, key int) *SeachTreeNode {
	if node == nil {
		return &SeachTreeNode{data: data, key: key}
	}
	if key > node.key {
		node.rightChild = bt.insertNode(node.rightChild, data, key)
	} else if key == node.key {

	} else {
		node.leftChild = bt.insertNode(node.leftChild, data, key)
	}
	return node
}

func (bt *BiSeachTree) Remove(key int) {
	bt.root = bt.removeNode(bt.root, key)
}
func (bt *BiSeachTree) removeNode(node *SeachTreeNode, key int) *SeachTreeNode {
	if node == nil {
		return nil
	}
	if key > node.key {
		node.rightChild = bt.removeNode(node.rightChild, key)
	} else if key < node.key {
		node.leftChild = bt.removeNode(node.leftChild, key)
	} else {
		if node.rightChild == nil && node.leftChild == nil {
			node = nil
		} else if node.rightChild != nil && node.leftChild == nil {
			node = node.rightChild
		} else if node.rightChild == nil && node.leftChild != nil {
			node = node.leftChild
		} else {
			// 寻找待删除节点左子树最大值（最大值一定是左子树最右边的节点）
			tempTree := node.leftChild
			tempPtTree := node
			for tempTree.rightChild != nil {
				tempPtTree = tempTree
				tempTree = tempTree.rightChild
			}
			nodeRightChild := node.rightChild
			nodeLeftChild := node.leftChild
			node = tempTree
			tempPtTree.rightChild = tempTree.leftChild
			node.leftChild = nodeLeftChild
			node.rightChild = nodeRightChild
		}
	}
	return node
}
