package main

import (
  "fmt"
)

type AvlNode struct {
  left, right *AvlNode
  value, height int
}

func max (a, b int) int {
  if a>b {
    return a
  }
  return b
}

func calculateHeight (node *AvlNode) int{
  if node == nil {
    return 0
  }
  node.height = max (calculateHeight(node.left), calculateHeight(node.right)) + 1
  return node.height
}

//Insert "value" in the tree rooted at "root".
func insert (root *AvlNode, value int) {
  if root == nil || value == root.value {
    return //Node identity is based on its value. This node already exists in the tree or root is nil.
  }

  if value>root.value { //traverse to right
    if root.right==nil { //We found an empty spot
      root.right = &AvlNode {nil, nil, value, 0}
      return
    }
    insert (root.right, value)
  } else { //Here because value is smaller than root's
    if root.left==nil {
      root.left = &AvlNode {nil, nil, value, 0}
      return
    }
    insert (root.left, value)
  }
}

//This function just returns difference of heights between left and right
func balanceFactor (root *AvlNode) int {
  var rootRightHeight, rootLeftHeight int
  if root.right != nil {
    rootRightHeight = root.right.height
  }
  if root.left != nil {
    rootLeftHeight = root.left.height
  }
  return rootLeftHeight-rootRightHeight
}

func populateTree (node *AvlNode) {
  primes := [6]int{2, 3, 5, 7, 11, 13}
  for _, prime := range primes {
    insert (node, prime)
  }
}

func printTree (node *AvlNode) {
 if node == nil {
   return
 }
 fmt.Printf ("value %d, height %d\n", node.value, node.height)

 fmt.Printf ("Left for %d: \n", node.value)
 printTree(node.left)
 fmt.Printf ("End of Left for %d\n", node.value)

 fmt.Printf ("Right for %d:\n ", node.value)
 printTree(node.right)
 fmt.Printf ("End of Right for %d\n", node.value)
}

func main() {
 var root AvlNode
 root.value = 8
 populateTree(&root)
 fmt.Println(calculateHeight(&root))
 printTree(&root)
}

                                                                                                                                                                                   1,2           Top
