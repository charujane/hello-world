package main

import "fmt"

type AvlNode struct {
  left, right *AvlNode
  value int
}

func main() {
 var root AvlNode
 fmt.Println(root.left, root.right, root.value) 
}
