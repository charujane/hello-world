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

/*func calculateHeight (node *AvlNode) int {
  if node == nil {
    return 0
  }
  node.height = max (calculateHeight(node.left), calculateHeight(node.right)) + 1
  return node.height
}*/

func height (node *AvlNode) int {
  if node==nil {
    return 0
  }
  return node.height
}

func rightRotate (root *AvlNode) *AvlNode {
  if root == nil || root.left == nil {
    return root //We cannot do right rotation
  }
  
  node := root.left
  root.left = node.right
  node.right = root

  root.height = max (height(root.right), height(root.left)) + 1
  node.height = max (height(node.right), height(node.left)) + 1

  return node
}

func leftRotate (root *AvlNode) *AvlNode {
  if root == nil || root.right == nil {
    return root //We cannot do any left rotation
  }
 
  node := root.right
  root.right = node.left
  node.left = root

  root.height = max (height(root.right), height(root.left)) + 1
  node.height = max (height(node.right), height(node.left)) + 1
  
  return node
}

func leftRightRotate (root *AvlNode) *AvlNode {
  root.left = leftRotate (root.left)
  root = rightRotate(root)
  return root
}

func rightLeftRotate (root *AvlNode) *AvlNode {
  root.right = rightRotate(root.right)
  root = leftRotate(root)
  return root
}

func restoreBalance (root *AvlNode) *AvlNode {
  if balanceFactor(root) == -2 {
    if balanceFactor(root.right) == -1 {
      root = leftRotate(root)
    } else if balanceFactor(root.right) == 1{
      root = rightLeftRotate(root)
    }
  } else if balanceFactor(root) == 2 {
    if balanceFactor(root.left) == 1 {
      root = rightRotate(root)
    } else if balanceFactor(root.left) == -1 {
      root = leftRightRotate(root)
    }
  }
  return root
}

//Insert "value" in the tree rooted at "root".
func insert (root *AvlNode, value int) *AvlNode{

  if root == nil {
    root = &AvlNode {nil, nil, value, 0}

  } else if value > root.value { 
    root.right = insert (root.right, value)
    root = restoreBalance (root)

  } else if value < root.value { 
    root.left = insert (root.left, value)
    root = restoreBalance (root)
  }
  root.height = max (height(root.right), height(root.left)) + 1

  return root
}

//Returns the following combinations:
//non-nil node, false --> value was found and node is the parent of found node.
//non-nil node, true --> value was found but there is no parent because value was at root.
//nil, false --> value was not found in the tree.
func lookupParentOf (root *AvlNode, value int) (*AvlNode, bool) {
  if root == nil {
    return root, false
  }

  var found bool
  var node *AvlNode

  if value > root.value {
    node, found = lookupParentOf(root.right, value)
  } else if value < root.value {
    node, found = lookupParentOf(root.left, value)
  } else {
    return root, true
  }
  
  if node != nil {
    if found { 
      return root, false
    } 
  }

  return node, false
}

//Lookup "value" in tree rooted at "root", delete it, 
//restore balance, and return new root.
func delete (root *AvlNode, value int) *AvlNode{
  //Lookup value
  
  return root
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

func populateTree (node *AvlNode) *AvlNode{
  primes := [6]int{2, 3, 5, 7, 11, 13}
  //primes := [2]int{2, 3}
  for _, prime := range primes {
    node = insert (node, prime)
  }  

  return node
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
 node := populateTree(&root)
 printTree(node)
 fmt.Println(lookupParentOf(node, 11))
 fmt.Println(lookupParentOf(node, 13))
 fmt.Println(lookupParentOf(node, 7))
 fmt.Println(lookupParentOf(node, 2))
 fmt.Println(lookupParentOf(node, 3))
 fmt.Println(lookupParentOf(node, 20))
 fmt.Println(lookupParentOf(node, 0))  
}

