# hello-world
first git repository

We are going to create an AVL tree implementation in GO lang.

We are new to GO lang, so let's make some assumptions for simplicity that will hopefully help us make some progress quickly given the lack of experience with GO. Our tree is going to store int values until we figure out how to make it more generic. So far, in the GO tutorial, there is no mention of OOP which is strange. Is GO not OOP language?!.

So first step our node structure.

Assumptions:
Nodes are int-valued
We will NOT allow duplicates in our tree. Node equality is based on node values.  

AVL trees are height-balanced binary search trees. We need to calculate balance factor for each node. 
BalanceFactor(Node) = Height(Left Node) - Height(Right Node)
For balanced AVL tree (and hence a O(logn) access time), the balance factor for each node is either -1, or 0, or 1.

Let's write a recursive function to find heights. Where do we store calculated heights? Inside the node structure?

"insert" function figures out new heights as it inserts new nodes. So we should not need extra traversal for updating heights. Scratching the previous calculateHeights function. 

So, tree rotations are....interesting. This helped.
chrome-extension://oemmndcbldboiebfnladdacbdfmadadm/http://eniac.cs.qc.cuny.edu/andrew/csci700-11/lecture7.pdf

Tree structure is wrong, jumps into a never ending loop while printing, so need to get rotations right. Also height is no more correct (Used to be before rotations).

Works now, atleast for my testcase. Need more testing and need to write a delete function.

For implementing delete, we need to lookup the node to be deleted, remove it from the tree, find a node to replace it, and rebalance the tree. 
The lookup function needs to return the parent of the lookedup node because the tree rooted at the parent will need to be rebalanced once our node is gone. So we need to track the parent. 

Deletion Algorithm:
1. Lookup node to be deleted. If found then go to step 2. Else go to step 6.
2. Simplest case: The node does not have any children, go to step 2a
   2a. Delete the node, If parent exists, go to step 5
   2b. If no parent, delete node and return nil.
3. If node has one child. Update its parent to point to the node's child. Go to step 5.
4. If node has two children. 
   4a. Find the minimum node on the right. To do that traverse to the left most node in the tree rooted at right.
   4b. Replace the node to be deleted with the minimum node in the right subtree. Go to step 5.
5. Update parent. Restore balance at the parent. 
6. Return root.
   
