package main

import "fmt"

func main() {
	grid := [][]int{
		[]int{1, 2, 2, 3, 5},
		[]int{3, 2, 3, 4, 4},
		[]int{2, 4, 5, 3, 1},
		[]int{6, 7, 1, 4, 5},
		[]int{5, 1, 1, 2, 4},
	}

	fmt.Printf("%v\n", pacificAtlantic(grid))
	node7 := TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	node15 := TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}

	node20 := TreeNode{
		Val:   20,
		Left:  &node15,
		Right: &node7,
	}

	node9 := TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}

	node3 := TreeNode{
		Val:   3,
		Left:  &node9,
		Right: &node20,
	}

	ret := levelOrder(&node3)
	for _, r := range ret {
		fmt.Printf("%v\n", r)
	}
}

/*
  200. Number of Islands (Medium)
  Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

  Example 1:

  11110
  11010
  11000
  00000
  Answer: 1

  Example 2:

  11000
  11000
  00100
  00011
  Answer: 3
*/
func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				cnt++
				mark0(grid, i, j)
			}
		}
	}

	return cnt
}
func mark0(grid [][]byte, x, y int) {
	m := len(grid)
	n := len(grid[0])

	q := [][]int{}
	q = append(q, []int{x, y})

	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(q) != 0 {
		g := q[0]
		q = q[1:]

		grid[g[0]][g[1]] = '0'

		for _, d := range dir {
			r := d[0] + g[0]
			c := d[1] + g[1]

			if 0 <= r && r < m && 0 <= c && c < n && grid[r][c] == '1' {
				q = append(q, []int{r, c})
			}
		}
	}
}

/*
  130. Surrounded Regions (Medium)
  Given a 2D board containing 'X' and 'O' (the letter O), capture all regions
  surrounded by 'X'.

  A region is captured by flipping all 'O's into 'X's in that surrounded region.

  For example,
  X X X X
  X O O X
  X X O X
  X O X X
  After running your function, the board should be:

  X X X X
  X X X X
  X X X X
  X O X X
*/
func surroundingRegion(board [][]byte) {
	m := len(board)
	n := len(board[0])

	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			markI(board, i, 0)
		}
		if board[i][n-1] == 'O' {
			markI(board, i, n-1)
		}
	}

	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			markI(board, 0, j)
		}
		if board[m-1][j] == 'O' {
			markI(board, m-1, j)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'I' {
				board[i][j] = 'O'
			}
		}
	}
}

func markI(board [][]byte, i, j int) {
	m := len(board)
	n := len(board[0])

	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	q := [][]int{[]int{i, j}}
	for len(q) != 0 {
		next := q[0]
		q = q[1:]

		board[next[0]][next[1]] = 'I'

		for _, d := range dir {
			x := next[0] + d[0]
			y := next[1] + d[1]

			if 0 <= x && x < m && 0 <= y && y < n && board[x][y] == 'O' {
				q = append(q, []int{x, y})
			}
		}
	}
}

/*
  417. Pacific Atlantic Water Flow (Medium)
  Difficulty: Medium
  Given an m x n matrix of non-negative integers representing the height of each unit cell in a continent,
  the "Pacific ocean" touches the left and top edges of the matrix and the "Atlantic ocean" touches the right
  and bottom edges.

  Water can only flow in four directions (up, down, left, or right) from a cell to another one with height equal or lower.

  Find the list of grid coordinates where water can flow to both the Pacific and Atlantic ocean.

  Note:
  The order of returned grid coordinates does not matter.
  Both m and n are less than 150.
  Example:

  Given the following 5x5 matrix:

    Pacific ~   ~   ~   ~   ~
         ~  1   2   2   3  (5) *
         ~  3   2   3  (4) (4) *
         ~  2   4  (5)  3   1  *
         ~ (6) (7)  1   4   5  *
         ~ (5)  1   1   2   4  *
            *   *   *   *   * Atlantic

  Return:

  [[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]] (positions with parentheses in above matrix).
*/
func pacificAtlantic(heights [][]int) [][]int {
	m := len(heights)
	n := len(heights[0])

	p := make([][]int, m)
	for i := 0; i < m; i++ {
		p[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		pacificHeightBfs(heights, p, i, 0, 1)
		pacificHeightBfs(heights, p, i, n-1, 2)
	}

	for j := 0; j < n; j++ {
		pacificHeightBfs(heights, p, 0, j, 1)
		pacificHeightBfs(heights, p, m-1, j, 2)
	}

	ret := [][]int{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[i][j] == 3 {
				ret = append(ret, []int{i, j})
			}
		}
	}

	return ret
}

func pacificHeightBfs(h [][]int, p [][]int, x, y, v int) {
	m := len(h)
	n := len(h[0])
	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	q := [][]int{[]int{x, y}}

	for len(q) != 0 {
		next := q[0]
		q = q[1:]

		p[next[0]][next[1]] |= v

		for _, d := range dir {
			r := next[0] + d[0]
			c := next[1] + d[1]
			if 0 <= r && r < m && 0 <= c && c < n && h[next[0]][next[1]] <= h[r][c] && p[r][c]&v == 0 {
				q = append(q, []int{r, c})
			}
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 102. Binary Tree Level Order Traversal
func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		l := len(q)
		row := []int{}

		for l > 0 {
			n := q[0]
			q = q[1:]

			row = append(row, n.Val)

			if n.Left != nil {
				q = append(q, n.Left)
			}

			if n.Right != nil {
				q = append(q, n.Right)
			}
			l--
		}

		ret = append(ret, row)
	}

	return ret
}

// 103. Binary Tree Zigzag Level Order Traversal
func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}

	q := []*TreeNode{root}
	level := 0
	for len(q) > 0 {
		l := len(q)
		row := []int{}
		for l > 0 {
			n := q[0]
			q = q[1:]

			if level%2 == 0 {
				row = append(row, n.Val)
			} else {
				row = append([]int{n.Val}, row...)
			}

			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}

			l--
		}
		ret = append(ret, row)

		level++
	}

	return ret
}
