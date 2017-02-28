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

func calculateHeight (node *AvlNode) int {
  if node == nil {
    return 0
  }
  node.height = max (calculateHeight(node.left), calculateHeight(node.right)) + 1
  return node.height
}

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

//Restores the balance at root and returns the new root 
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
//node, 1 --> value was found and lies on the right of node which is the parent
//node, -1 --> value was found and lies on the left of node which is the parent
//node, 0 --> value was found but is the root, node is the value node itself
//nil, -2 --> value not found.
func lookupParentOf (root *AvlNode, value int) (*AvlNode, int) {
  var node *AvlNode
  orientation := -2

  if (root != nil) {
    if value > root.value {
      node, orientation = lookupParentOf(root.right, value)
      if node != nil {
        if orientation == 0{ 
          return root, 1
        } 
      }
    } else if value < root.value {
      node, orientation = lookupParentOf(root.left, value)
      if node != nil {
        if orientation == 0{ 
          return root, -1
        } 
      }
    } else {
      return root, 0
    }
  }

  return node, orientation
}

//Returns the minimum value in the tree at root.
func findMinimum (root *AvlNode) int {
  //Keep moving to left until you reach the leaf node.
  if root.left != nil {
    return findMinimum (root.left)
  } else { 
    return root.value
  }
}

//Lookup "value" in tree rooted at "root", delete it, 
//restore balance, and return new root.
func delete (root *AvlNode, value int) *AvlNode{
  var nodeToBeDeleted *AvlNode
  
  parent, orientation := lookupParentOf (root, value) 
  if parent == nil {
     //fmt.Printf("The value %d, was not found in the tree rooted at node %s\n", value, root)
     return root
  }

  switch orientation {
  case 1:
    nodeToBeDeleted = parent.right
  case 0:
    nodeToBeDeleted = parent
  case -1:
    nodeToBeDeleted = parent.left
  default :
    //TODO: Throw assertion error of some kind, figure out exceptions in GO 
    fmt.Println("Should not get here 1\n")
  }

  var dNodeOnlyChild *AvlNode
  dNodeRchild := nodeToBeDeleted.right
  dNodeLchild := nodeToBeDeleted.left
  if  dNodeRchild == nil &&  dNodeLchild != nil {
    dNodeOnlyChild = dNodeLchild 
  } else if dNodeRchild != nil && dNodeLchild == nil {
    dNodeOnlyChild = dNodeRchild 
  } else if dNodeRchild != nil && dNodeLchild != nil {
    //nodeToBeDeleted has 2 children. Yikes.
    minVal := findMinimum(dNodeRchild)
    delete(root, minVal) //Should be simple leaf node deletion
    nodeToBeDeleted.value = minVal 
  } 

  //nodeToBeDeleted has only one child or none.
  if dNodeRchild == nil || dNodeLchild == nil {
    switch orientation {
      case 1:
        parent.right = dNodeOnlyChild
      case -1:
        parent.left = dNodeOnlyChild
      case 0:
        //nodeToBeDeleted is the root and has only one child
        root = dNodeOnlyChild
      default:
        //TODO: Throw assertion error of some kind, figure out exceptions in GO 
        fmt.Println("Should not get here 2")
    }
  }
 
  if root != nil { 
    calculateHeight(root)
    grandParent, parentOrientation := lookupParentOf (root, parent.value)
    switch parentOrientation {
    case 1:
      grandParent.right = restoreBalance(parent) 
    case -1:
      grandParent.left = restoreBalance(parent)
    case 0:
      //parent was the root
      root = restoreBalance(root)
    default: 
      //TODO: Throw assertion error of some kind, figure out exceptions in GO 
      fmt.Println("Should not get here 3")
    }
  }

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
  primes := [5]int{2, 3, 5, 11, 13}
  //primes := [2]int{2, 3}
  for _, prime := range primes {
    node = insert (node, prime)
  }  

  return node
}

func printTree (node *AvlNode, prefix string) {
 if node == nil {
   return
 }
 fmt.Printf ("%svalue %d, height %d\n", prefix, node.value, node.height)

 fmt.Printf ("%sLeft for %d: \n", prefix, node.value)
 printTree(node.left, prefix+"  ")
 fmt.Printf ("%sEnd of Left for %d\n", prefix, node.value)

 fmt.Printf ("%sRight for %d:\n ", prefix, node.value)
 printTree(node.right, prefix+"  ")
 fmt.Printf ("%sEnd of Right for %d\n", prefix, node.value)
}

func main() {
 //test
 node := &AvlNode {nil, nil, 8, 0}
 node = populateTree(node)
 printTree(node,"")

 //Vanishing tree. Delete all nodes one by one testing different conditions. 
 node= delete (node, 11) //Delete node with one child
 node= delete (node, 13) //Delete leaf on right
 node= delete (node, 5) //Delete leaf on left
 printTree(node,"")
 fmt.Println("---------------------------------------------------------------")
 node= delete (node, 3) //Delete new root and also node with two children!
 printTree(node,"")
 fmt.Println("---------------------------------------------------------------")
 node= delete (node, 2) //Delete new root
 node= delete (node, 13) //Delete leaf
 printTree(node,"")
 fmt.Println("---------------------------------------------------------------")
 node= delete (node, 13) //Delete non-existent node
 node= delete (node, 8) //Delete the last node
 printTree(node,"")
 fmt.Println("---------------------------------------------------------------")
}

