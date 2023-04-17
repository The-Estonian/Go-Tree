package main

import "fmt"

// Tree
type Node struct{
	key int
	left *Node
	right *Node
}
//Insert
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

//Search
func (n *Node) search(k int) bool{
	if n == nil {
		return false
	}
	if n.key < k {
		//move right
		return n.right.search(k)
	} else if n.key > k {
		//move left
		return n.left.search(k)
	}
	return true
}


func main() {
	tree := &Node{key: 100}
	tree.insert(200)
	tree.insert(300)
	tree.insert(233)
	tree.insert(23)
	tree.insert(15)
	tree.insert(500)
	fmt.Println(tree)
	fmt.Println(tree.right.right)
	fmt.Println(tree.search(233))
	fmt.Println(tree.search(234))
}
