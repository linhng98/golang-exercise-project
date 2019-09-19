package main

import(
    "fmt"
    "os"
)

// Node struct
type Node struct {
	key    int
	left   *Node
	right  *Node
	height int
}

// InitNode init node value
func InitNode(val int) (newnd *Node){
    var nd Node

    nd.height=0
    nd.key=val
    nd.left=nil
    nd.right=nil

    return &nd
}

// PrintTree print value in tree
func (nd *Node) PrintTree(){
    if nd==nil{
        return
    }
    nd.left.PrintTree()
    nd.right.PrintTree()
    fmt.Printf("%d  ",nd.key);
}

// AddNode func add node to tree
func (nd *Node) AddNode(val int){

    // init value for node
    if(nd==nil){
        nd=InitNode(val)
        return
    }
    
    // 
}

func main() {

}
