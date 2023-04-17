package node

type Node struct {
	Key int
	Left *Node
	Right *Node
}

func (n *Node) Insert(k int){
	if n.Key < k {
		//right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}
