package main

import (
	"strconv"
	"strings"
)

func main() {
	a := []int{10, 20, 30, 50, 60, 80, 110, 130, 140, 170}
	println(binarySearch(a, 85))
	println(findclosest(a, 9))
	println(findclosest(a, 172))

	A := [][]int{
		{1, 2},
		{1, 5},
		{2, 5, 2},
		{2, 6, 3},
		{2, 2, 1},
		{2, 3, 2},
	}
	println(obstacleSolution(A))

}

func binarySearch(a []int, k int) int {
	l := 0
	r := len(a) - 1
	for l <= r {
		m := l + (r-l)/2
		if a[m] == k {
			return m
		} else if a[m] < k {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}

func findclosest(a []int, k int) (int, int) {
	l := 0
	r := len(a) - 1

	lclose := -1
	rclose := r + 1

	for l <= r {
		m := l + (r-l)/2
		if a[m] == k {
			return m, m
		} else if a[m] < k {
			lclose = m
			l = m + 1
		} else {
			rclose = m
			r = m - 1
		}
	}

	return lclose, rclose
}

/*
Given an infinite number line, you would like to build few blocks and obstacles on it.
Specifically, you have to implement code which supports two types of operations:
  - {1, x} - builds an obstacle at coordinate x along the number line. It is guaranteed that coordinate x
             does not contain any obstacle when the operation is performed.
  - {2, x, size} - check whether it's possible to build a block occupying coordinates between
    x - size, x - size + 1, ... , x - 1 along the number line.
    Return 1 if it is possible, or return 0 otherwise.
    Please note that this operation does not actually build the block, it only checks whether a block can be built.
Given an array of operations, your task is to return a binary string representing the output for all {2, x, size} operations.

Example
For operations = [][]int{
		{1, 2},
		{1, 5},
		{2, 5, 2},
		{2, 6, 3},
		{2, 2, 1},
		{2, 3, 2},
	}
The output should be "1010"
           -3  -2   -1  0   1   2   3   4   5   6   7   8
------------o---o---o---o---o---X---o---o---X---o---o---o
                                    ^   ^

*/

func obstacleSolution(ops [][]int) string {
	a := []int{}

	ret := ""
	for _, op := range ops {
		if op[0] == 1 {
			l, r := findclosest(a, op[1])
			if l == -1 {
				a = append([]int{op[1]}, a...)
			} else if r == len(a) {
				a = append(a, op[1])
			} else {
				rightPart := a[r:]
				a = append(a, op[1])
				a = append(a, rightPart...)
			}
		} else {
			start := op[1] - op[2]
			end := op[1]
			l, _ := findclosest(a, start)

			if l+1 == len(a) {
				ret += "1"
			} else {
				if a[l+1] < end {
					ret += "0"
				} else {
					ret += "1"
				}
			}
		}
	}

	return ret
}

/*
  99. Recover Binary Search Tree
  Two elements of a binary search tree (BST) are swapped by mistake.

  Recover the tree without changing its structure.

  Note:
  A solution using O(n) space is pretty straight forward. Could you devise a constant space solution?

  Solution:
  it's like a swap in a sorted array.
  There could be one or two pairs of nodes are out of order:
  - case I, the swapped node are adjacent, one pair
  - case II: the swapped nodes are not adjacent, two pairs
  1) In-order traverse the tree have a prev and cur pointer point to previous and current node
     Use firstNode pointer to record the first node of the first out-of-order pair
     Use secondNode pointer to record the second node of the second out-of-order pair
  2) Swap the value of firstNode and secondNode
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var first, second, prev *TreeNode

	inorder(root, &prev, &first, &second)
	tmp := first.Val
	first.Val = second.Val
	second.Val = tmp
}

func inorder(r *TreeNode, prev, first, second **TreeNode) {
	if r == nil {
		return
	}

	inorder(r.Left, prev, first, second)
	if *prev != nil && (*prev).Val > r.Val {
		if *first == nil {
			*first = *prev
		}
		*second = r
	}

	*prev = r

	inorder(r.Right, prev, first, second)
}

/*
  297. Serialize and Deserialize Binary Tree (Hard)
  Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored
  in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or
  another computer environment.

  Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization
  algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be
  deserialized to the original tree structure.

  For example, you may serialize the following tree

      1
     / \
    2   3
       / \
      4   5
  as "[1,2,3,null,null,4,5]", just the same as how LeetCode OJ serializes a binary tree. You do not necessarily need to
  follow this format, so please be creative and come up with different approaches yourself.
  Note: Do not use class member/global/static variables to store states. Your serialize and deserialize algorithms should be stateless.

  Solution:
  The idea is simple: print the tree in pre-order traversal and use "X" to denote null node and split node with ",".
  We can use a StringBuilder for building the string on the fly. For deserializing, we use a Queue to store the pre-order
  traversal and since we have "X" as null node, we know exactly how to where to end building subtress.
*/
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	s := []string{}
	this.serializeHelper(root, &s)

	return strings.Join(s, ",")
}
func (this *Codec) serializeHelper(r *TreeNode, s *[]string) {
	if r == nil {
		*s = append(*s, "null")
		return
	}

	*s = append(*s, strconv.Itoa(r.Val))
	this.serializeHelper(r.Left, s)
	this.serializeHelper(r.Right, s)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(s string) *TreeNode {
	st := strings.Split(s, ",")
	return this.deserializeHelper(&st)
}
func (this *Codec) deserializeHelper(st *[]string) *TreeNode {
	if len(*st) == 0 {
		return nil
	}

	n, err := strconv.Atoi((*st)[0])
	*st = (*st)[1:]

	if err != nil {
		return nil
	}

	r := &TreeNode{Val: n}

	r.Left = this.deserializeHelper(st)
	r.Right = this.deserializeHelper(st)

	return r
}
