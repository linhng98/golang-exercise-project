package main

import (
	"fmt"
	"os"
	"strconv"
)

// Node struct
type Node struct {
	key   int
	left  *Node
	right *Node
}

// InitNode init node value
func InitNode(val int) *Node {
	nd := Node{key: val, left: nil, right: nil}
	return &nd
}

// PrintTree print value in tree
func (nd *Node) PrintTree() {
	if nd == nil {
		return
	}
	nd.left.PrintTree()
	fmt.Printf("%d(%d)   ", nd.key, GetHeight(nd))
	nd.right.PrintTree()
}

// AddNode func add node to tree
func AddNode(root *Node, val int) *Node {

	// init value for node
	if root == nil {
		return InitNode(val)
	}

	// recurse to find position to add node
	if val > root.key {
		root.right = AddNode(root.right, val)
	} else {
		root.left = AddNode(root.left, val)
	}

	// check this node balance or not
	bal := GetBalance(root)
	lbal := GetBalance(root.left)
	rbal := GetBalance(root.right)

	if bal == -2 && lbal == -1 { // left left case
		root = RightRotate(root)
	} else if bal == -2 && lbal == 1 { // left right case
		root.left = LeftRotate(root.left)
		root = RightRotate(root)
	} else if bal == 2 && rbal == -1 { // right left case
		root.right = RightRotate(root.right)
		root = LeftRotate(root)
	} else if bal == 2 && rbal == 1 { //right right case
		root = LeftRotate(root)
	}

	return root
}

// LeftRotate tree
func LeftRotate(root *Node) *Node {
	newroot := root.right
	root.right = newroot.left
	newroot.left = root

	return newroot
}

// RightRotate tree
func RightRotate(root *Node) *Node {
	newroot := root.left
	root.left = newroot.right
	newroot.right = root

	return newroot
}

// GetHeight func get heigh of tree
func GetHeight(root *Node) int {
	if root == nil {
		return 0
	}
	return 1 + GetMax(GetHeight(root.left), GetHeight(root.right))
}

// GetBalance func
func GetBalance(root *Node) int {
	if root == nil {
		return 0
	}

	return GetHeight(root.right) - GetHeight(root.left)
}

// GetMax func
func GetMax(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var root *Node

	for _, val := range os.Args[1:] {
		num, _ := strconv.Atoi(val)
		root = AddNode(root, num)
	}

	root.PrintTree()
	fmt.Printf("\n")
}
