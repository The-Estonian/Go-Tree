package main

import (
	"fmt"
	"The-Estonian/Go-Tree/node"
)

func main() {
	tree := &node.Node{Key: 100}
	tree.Insert(200)
	tree.Insert(50)
	fmt.Println(tree)
}
