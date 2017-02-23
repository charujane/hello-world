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


