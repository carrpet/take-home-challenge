## Assumptions:

The assumptions made for the challenge are as follows:
1. The given tree is an arbitrary unordered binary tree.  The code could be easily extended to support n-ary trees.
2. The values of the node ids are integers. Generics could be used to relax this assumption and extend the code to support any orderable id type.
3. The function does not modify the given tree. 


## Algorithm and analysis:

The input tree representation was first transformed to a slice data structure with node data annotated with tree level information using a breadth-first pass.  The nodes were then sorted by id value in ascending order. A linear scan was performed to find the duplicates and their minimum associated levels.  The resulting slice of duplicates was then sorted by level value to find the minimum level duplicate.

## Time complexity:

Let n denote the number of nodes in the input tree.  The best sorting algorithms work in O(n log n) average and worst case time, which is assumed to be the time complexity for the Go standard library sort routines.  The initial tree traversal and linear scan to find duplicates each have a time complexity of O(n).  Thus the total time complexity is dominated by the sort routine, which has average time complexity of O(n log n). 

## Space complexity:

The function requires a buffer of length n to track the level information as the ids are sorted.  Thus the space complexity requirement is O(n).