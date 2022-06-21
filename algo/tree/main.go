package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

func main() {
	n4 := TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	n5 := TreeNode{
		Val:   5,
		Left:  &n4,
		Right: nil,
	}

	n2 := TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}

	n3 := TreeNode{
		Val:   3,
		Left:  &n2,
		Right: &n5,
	}

	n1 := TreeNode{
		Val:   1,
		Left:  nil,
		Right: &n3,
	}

	n6 := TreeNode{
		Val:   6,
		Left:  &n1,
		Right: nil,
	}

	fmt.Printf("%v\n", binaryTreePaths(&n6))
}

//binary tree, 节点有左，右和父指针.How to traverse the tree with O(1) space?
func traverseTree(r *TreeNode) []int {
	ret := []int{}
	var prev *TreeNode
	cur := r

	for cur != nil {
		if prev == cur.Parent {
			if cur.Left != nil {
				prev = cur
				cur = cur.Left
			} else {
				ret = append(ret, cur.Val)
				prev = cur
				if cur.Right != nil {
					cur = cur.Right
				} else {
					cur = cur.Parent
				}
			}
		} else if prev == cur.Left {
			ret = append(ret, cur.Val)
			prev = cur
			if cur.Right != nil {
				cur = cur.Right
			} else {
				cur = cur.Parent
			}
		} else {
			cur = prev
			prev = cur.Parent
		}
	}

	return ret
}

// 98. Validate Binary Search Tree
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil && maxVal(root.Left) >= root.Val {
		return false
	}

	if root.Right != nil && minVal(root.Right) <= root.Val {
		return false
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
}

func minVal(r *TreeNode) int {
	for r.Left != nil {
		r = r.Left
	}

	return r.Val
}

func maxVal(r *TreeNode) int {
	for r.Right != nil {
		r = r.Right
	}

	return r.Val
}

// 108. Convert Sorted Array to Binary Search Tree
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2

	l := sortedArrayToBST(nums[0:mid])
	r := sortedArrayToBST(nums[mid+1:])

	return &TreeNode{
		Val:   nums[mid],
		Left:  l,
		Right: r,
	}
}

//173. Binary Search Tree Iterator
type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	s := []*TreeNode{}

	cur := root
	for cur != nil {
		s = append(s, cur)
		cur = cur.Left
	}

	return BSTIterator{stack: s}
}

func (this *BSTIterator) Next() int {
	cur := this.stack[len(this.stack)-1]
	this.stack = this.stack[0 : len(this.stack)-1]
	v := cur.Val

	cur = cur.Right
	for cur != nil {
		this.stack = append(this.stack, cur)
		cur = cur.Left
	}

	return v
}

func (this *BSTIterator) HasNext() bool {
	if len(this.stack) != 0 {
		return true
	}

	return false
}

//235. Lowest Common Ancestor of a Binary Search Tree
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	cur := root
	for cur != nil {
		if cur.Val > p.Val && cur.Val > q.Val {
			cur = cur.Right
		} else if cur.Val < p.Val && cur.Val < q.Val {
			cur = cur.Left
		} else {
			break
		}
	}

	return cur
}

// 257. Binary Tree Paths
func binaryTreePaths(root *TreeNode) []string {
	results := [][]string{}

	binaryTreepathDfs(root, []string{}, &results)

	ret := []string{}
	for _, r := range results {
		ret = append(ret, strings.Join(r, "->"))
	}

	return ret
}

// pre-order traverse
func binaryTreepathDfs(r *TreeNode, path []string, results *[][]string) {
	path = append(path, strconv.Itoa(r.Val))

	if r.Left == nil && r.Right == nil {
		// note: make a copy
		cp := make([]string, len(path))
		copy(cp, path)
		*results = append(*results, cp)
		return
	}

	if r.Left != nil {
		binaryTreepathDfs(r.Left, path, results)
	}
	if r.Right != nil {
		binaryTreepathDfs(r.Right, path, results)
	}
}

// 101. Symmetric Tree
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isMirror(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if (l == nil && r != nil) || (l != nil && r == nil) {
		return false
	}

	return (l.Val == r.Val) && isMirror(l.Left, r.Right) && isMirror(l.Right, r.Left)
}

// 230. Kth Smallest Element in a BST
func kthSmallest(root *TreeNode, k int) int {
	stk := []*TreeNode{}

	cur := root
	for cur != nil {
		stk = append(stk, cur)
		cur = cur.Left
	}

	var v int
	for k > 0 {
		cur = stk[len(stk)-1]
		stk = stk[:len(stk)-1]

		v = cur.Val

		cur = cur.Right
		for cur != nil {
			stk = append(stk, cur)
			cur = cur.Left
		}

		k--
	}

	return v
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
// Your Codec object will be instantiated and called as such:
// Codec codec = new Codec();
// codec.deserialize(codec.serialize(root));
type Codec struct {
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
	nodes := make([]*TreeNode, len(st))

	for i, n := range st {
		if n != "null" {
			v, _ := strconv.Atoi(n)
			nodes[i] = &TreeNode{Val: v}
		}
	}
	return this.deserializeHelper(&nodes)
}

func (this *Codec) deserializeHelper(nodes *[]*TreeNode) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}

	node := (*nodes)[0]
	*nodes = (*nodes)[1:]
	if node == nil {
		return nil
	}

	node.Left = this.deserializeHelper(nodes)
	node.Right = this.deserializeHelper(nodes)

	return node
}
