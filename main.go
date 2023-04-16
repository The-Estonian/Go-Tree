package main

import "fmt"

type Node struct{
	key int
	left *Node
	right *Node
}

func (n *Node) insert(k int) {
	if n.key < k {
		//move right
		if n.right == nil {
			n.right = &Node{key:k}
		} else {
			n.right.insert(k)
		}
	} else if n.key > k {
		//move left
		if n.left == nil {
			n.left = &Node{key:k}
		} else {
			n.left.insert(k)
		}
	}
}

func main() {
	tree := &Node{key: 100}
	tree.insert(200)
	tree.insert(300)
	fmt.Println(tree)
}
