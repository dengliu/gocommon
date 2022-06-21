package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

var (
	dirs = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
)

func main() {
	//moves := [][]int{{0,0}, {2,0},{1,1},{2,1},{2,2}}
	//moves := [][]int{{0,0}, {1,1},{0,1},{0,2},{1,0},{2,0}}
	//moves := [][]int{{0,0}, {1,1},{2,0},{1,0},{1,2},{2,1},{0,1},{0,2},{2,2}}
	println(isIsomorphic("badc", "baba"))
	fmt.Println(numOfWays([]int{1, 2, 3}, []int{2, 4}, [][]int{{1, 5}, {0, 0, 1}, {1, 5}}))

	b := figureUnderGravity([][]byte{
		{'F', 'F', 'F'},
		{'.', 'F', '.'},
		{'.', 'F', 'F'},
		{'#', 'F', '.'},
		{'F', 'F', '.'},
		{'.', '.', '.'},
		{'.', '.', '#'},
		{'.', '.', '.'},
	})

	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			fmt.Printf("%c ", b[i][j])
		}
		fmt.Println()
	}

	buggleGame(
		[][]int{{1, 2, 3},
			{1, 2, 4},
			{1, 2, 2}},
		[][]int{{0, 1}, {2, 1}})

	dict := make(map[string]int)
	dict["foo"] = 1
	fmt.Println(dict["foo"])
	bar := dict["bar"]
	fmt.Println(bar)

	dict2 := make(map[string]string)
	dict2["foo"] = "foo"

	f := dict2["foo"]
	bb := dict2["bar"]

	fmt.Println(f, bb)

	relevantCourse := map[string]bool{}
	fmt.Println(relevantCourse["c"])

	slice := []int{1}
	fmt.Printf("%v\n", slice)
	slice = slice[:len(slice)-1]
	fmt.Printf("%v\n", slice)

	str := "a"
	str1 := str[1:]
	fmt.Println(str1)

	fmt.Println(almostEqualNumbers([]int{1, 151, 241, 1, 9, 22, 351}))

	fmt.Println(numSub("303"))
	fmt.Println(isCyclic([]int{1, 2, 3, 6, 5, 4}, []int{3, 6, 5, 4, 1, 2}))

	fmt.Print(minChatter("chatchht"))

	g := [][]byte{
		{'c', 'c', 'x', 't', 'i', 'b'},
		{'c', 'c', 'a', 't', 'n', 'i'},
		{'a', 'c', 'n', 'n', 't', 't'},
		{'t', 'c', 's', 'i', 'p', 't'},
		{'a', 'o', 'o', 'o', 'a', 'a'},
		{'o', 'a', 'a', 'a', 'o', 'o'},
		{'k', 'a', 'i', 'c', 'k', 'i'},
	}
	fmt.Printf("%v\n", findword(g, "bit"))
	fmt.Println("%v\n", sums("99", "999"))
	fmt.Printf("%v\n", fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))

	message := [][]string{
		{"1", "Bob hello"},
		{"2", "Alice hi"},
		{"1", "How is your life"},
		{"1", "Better than before"},
		{"2", "M super"},
		{"1", "Wow pro"},
	}

	fmt.Printf("%v\n", rendermessage(message, 6, 15))

	fmt.Printf("%v\n", numsum("99", "999"))

	println(topGame2([]string{
		"1500000000,user1,1001,join",
		"1500000010,user1,1002,join",
		"1500000015,user1,1002,quit",
		"1500000025,user1,1001,quit"}))

	r := findPeak([]int{5, 4, 2, 4, 4, 3, 7, 7, 5, 6, 6})
	fmt.Printf("peak %v\n", r)

	fmt.Println(canAttendAll([][]int{{0, 30}, {15, 20}, {5, 10}}))
}

/*
   252. Meeting Rooms   Difficulty: Easy
  Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei),
  determine if a person could attend all meetings.

  For example,
  Given [[0, 30],[5, 10],[15, 20]],
  return false.

  Interval schedule
  这实际上就是求区间是否有交集的问题，我们可以先给所有区间排个序，用起始时间的先后来排，
  然后我们从第二个区间开始，如果开始时间早于前一个区间的结束时间，则说明会议时间有冲突，返回false，
  遍历完成后没有冲突，则返回true
*/
func canAttendAll(a [][]int) bool {
	sort.Slice(a, func(i, j int) bool {
		if a[i][0] < a[j][0] {
			return true
		}
		return false
	})

	for i := 1; i < len(a); i++ {
		if a[i][0] < a[i-1][1] {
			return false
		}
	}

	return true
}

/*
   253. Meeting Rooms II    Difficulty: Medium
   Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei),
   find the minimum number of conference rooms required.

    For example,
    Given [[0, 30],[5, 10],[15, 20]],
    return 2.

    Interval partition

    Sort both by start time and end time, and this problem can be convert to max number of Overlaps
    O(2n) space complexity
*/
func minMeetingRooms(input [][]int) int {
	n := len(input)
	start := make([]int, n)
	end := make([]int, n)

	for i, m := range input {
		start[i] = m[0]
		end[i] = m[1]
	}

	sort.Ints(start)
	sort.Ints(end)

	maxOverlap := 0
	curOverlap := 0

	startIdx := 0
	endIdx := 0
	for startIdx < n && endIdx < n {
		if startIdx < endIdx {
			curOverlap++
			if curOverlap > maxOverlap {
				maxOverlap = curOverlap
			}
		} else {
			endIdx++
			curOverlap--
		}
	}

	return maxOverlap
}

/*
419. Battleships in a Board
Given an m x n matrix board where each cell is a battleship 'X' or empty '.', return the number of the battleships on board.

Battleships can only be placed horizontally or vertically on board. In other words, they can only be made of the shape 1 x k (1 row, k columns) or k x 1 (k rows, 1 column), where k can be of any size. At least one horizontal or vertical cell separates between two battleships (i.e., there are no adjacent battleships).

Input: board = [["X",".",".","X"],
                [".",".",".","X"],
                [".",".",".","X"]]
Output: 2
Example 2:

Input: board = [["."]]
Output: 0


Constraints:

m == board.length
n == board[i].length
1 <= m, n <= 200
board[i][j] is either '.' or 'X'.

Solution 1:

Going over all cells, we can count only those that are the "first" cell of the battleship. First cell will be defined as the most top-left cell. We can check for first cells by only counting cells that do not have an 'X' to the left and do not have an 'X' above them.
*/
func countBattleships(board [][]byte) int {
	m := len(board)
	n := len(board[0])
	var cnt int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == byte('.') {
				continue
			}

			if i > 0 && board[i-1][j] == byte('X') {
				continue
			}
			if j > 0 && board[i][j-1] == byte('X') {
				continue
			}

			cnt++
		}
	}

	return cnt
}

/*
1275. Find Winner on a Tic Tac Toe Game
Tic-tac-toe is played by two players A and B on a 3 x 3 grid. The rules of Tic-Tac-Toe are:

Players take turns placing characters into empty squares ' '.
The first player A always places 'X' characters, while the second player B always places 'O' characters.
'X' and 'O' characters are always placed into empty squares, never on filled ones.
The game ends when there are three of the same (non-empty) character filling any row, column, or diagonal.
The game also ends if all squares are non-empty.
No more moves can be played if the game is over.
Given a 2D integer array moves where moves[i] = [rowi, coli] indicates that the ith move will be played on grid[rowi][coli]. return the winner of the game if it exists (A or B). In case the game ends in a draw return "Draw". If there are still movements to play return "Pending".

You can assume that moves is valid (i.e., it follows the rules of Tic-Tac-Toe), the grid is initially empty, and A will play first.
*/

func tictactoe(moves [][]int) string {
	var arow [3]int
	var acol [3]int
	var adiag1, adiag2 int

	for i, v := range moves {
		x := v[0]
		y := v[1]

		inc := 1
		if i%2 == 1 {
			inc = -1
		}

		arow[x] += inc
		acol[y] += inc

		if x == y {
			adiag1 += inc
		}

		if x+y == 2 {
			adiag2 += inc
		}
		if arow[x] == 3 || acol[y] == 3 || adiag1 == 3 || adiag2 == 3 {
			return "A"
		}

		if arow[x] == -3 || acol[y] == -3 || adiag1 == -3 || adiag2 == -3 {
			return "B"
		}
	}

	if len(moves) == 9 {
		return "Draw"
	}

	return "Pending"
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		v, ok := m[s[i]]
		if !ok {
			m[s[i]] = t[i]
			continue
		}

		if v != t[i] {
			return false
		}
	}

	return true
}

/*2. 题目：deleteMinimalPeaks

给一个 array，要求找出 minimal peak，然后把 minimal peak 放到新的
array 重新排一次

对 minimal peak 定义：`(A[i]>A[i+] 或 A[i+1]不存在)` 且 `(A[i]>A[i-1] 或 A[i-1] 不存在`

input : 4 2 5 3 7 9

output : 第一步找出最小 peak 是 4，然后把 4 移出，新的 array 变成
`2 5 3 7 9`，result 的 array 第一个数字是 4.

input :  1 4 5 3 8 6
[1,4,5,3,8,6], Min peak = 5, resultant = [1,4,3,8,6]
[1,4,3,8,6], Min peak = 4, resultant = [1,3,8,6]
[1,3,8,6], Min peak = 8, resultant = [1,3,6]
[1,3,6], Min peak = 6, resultant = [1,3]
[1,3], Min peak = 3, resultant = [1]
[1], min peak = 1
Output = [5,4,8,6,3,1]

如此循环一直到 sort 完一整个 input array
*/
func minPeak(a []int) []int {
	r := []int{}
	n := len(a)

	for i := 0; i < n; i++ {
		l := len(a)

		index := -1
		min := math.MaxInt32

		for j := 0; j < l; j++ {
			if l == 1 {
				index = 0
				min = a[j]
				break
			}
			if (j == 0 && a[j] > a[j+1]) ||
				(j+1 == l && a[j] > a[j-1]) ||
				(j > 0 && j+1 < l && a[j-1] < a[j] && a[j] > a[j+1]) {
				if a[j] < min {
					index = j
					min = a[j]
				}
			}
		}

		r = append(r, min)

		a = append(a[:index], a[index+1:]...)
	}

	return r
}

/*
goodTuples

Give an array and find the count of a pair number and a single number combination in a row of this array.
Target array is a[i - 1], a[i], a[i + 1]

Input: a = [1, 1, 2, 1, 5, 3, 2, 3]

Output: 3

Explain:

[1, 1, 2] -> two 1 and one 2(O)
[1, 2, 1] -> two 1 and one 2(O)
[2, 1, 5] -> one 2, one 1 and one five(X)
[1, 5, 3] -> (X)
[5, 3, 2] -> (X)
[3, 2, 3] -> (O)

Time: O(n)
*/
func goodTuples(a []int) int {
	var cnt int
	for i := 0; i < len(a)-2; i++ {
		if a[i] == a[i+1] && a[i] == a[i+2] {
			continue
		}
		if a[i] == a[i+1] || a[i] == a[i+2] || a[i+1] == a[i+2] {
			cnt++
		}
	}

	return cnt
}

/*
945. Minimum Increment to Make Array Unique
You are given an integer array nums. In one move, you can pick an index i where 0 <= i < nums.length and increment nums[i] by 1.

Return the minimum number of moves to make every value in nums unique.

Example 1:

Input: nums = [1,2,2]
Output: 1
Explanation: After 1 move, the array could be [1, 2, 3].
Example 2:

Input: nums = [3,2,1,2,1,7]
Output: 6
Explanation: After 6 moves, the array could be [3, 4, 1, 2, 5, 7].
It can be shown with 5 or less moves that it is impossible for the array to have all unique values.

Constraints:

1 <= nums.length <= 105
0 <= nums[i] <= 105
*/
func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	cnt := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			cnt += nums[i-1] - nums[i] + 1
			nums[i] = nums[i-1] + 1
		}
	}
	return cnt
}

/*
Rearrange an array so that arr[i] becomes arr[arr[i]] with O(1) extra space
Difficulty Level : Hard

Given an array arr[] of size n where every element is in range from 0 to n-1. Rearrange the given array so that arr[i] becomes arr[arr[i]]. This should be done with O(1) extra space.

Examples:

Input: arr[]  = {3, 2, 0, 1}
Output: arr[] = {1, 0, 3, 2}
Explanation:
In the given array
arr[arr[0]] is 1 so arr[0] in output array is 1
arr[arr[1]] is 0 so arr[1] in output array is 0
arr[arr[2]] is 3 so arr[2] in output array is 3
arr[arr[3]] is 2 so arr[3] in output array is 2

Input: arr[] = {4, 0, 2, 1, 3}
Output: arr[] = {3, 4, 2, 0, 1}
Explanation:
arr[arr[0]] is 3 so arr[0] in output array is 3
arr[arr[1]] is 4 so arr[1] in output array is 4
arr[arr[2]] is 2 so arr[2] in output array is 2
arr[arr[3]] is 0 so arr[3] in output array is 0
arr[arr[4]] is 1 so arr[4] in output array is 1

Input: arr[] = {0, 1, 2, 3}
Output: arr[] = {0, 1, 2, 3}
Explanation:
arr[arr[0]] is 0 so arr[0] in output array is 0
arr[arr[1]] is 1 so arr[1] in output array is 1
arr[arr[2]] is 2 so arr[2] in output array is 2
arr[arr[3]] is 3 so arr[3] in output array is 3
*/
func reArrangeArray(a []int) []int {
	b := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = a[a[i]]
	}

	return b
}

/*
3. Figure under figureUnderGravity 应该是下面这道题：

grid 上有一个不规则形状的图形，上面有三种格子：
` # 表示障碍`
` . 表示空`
` F 表示图形上的一个像素`

在重力作用下，图形从当前位置掉落后，新生成的 Grid 是什么样？

You are given a rectangular matrix of characters matrix, which represents a 2-dimensional field where each cell is either empty ('.'), contains an obstacle (#), or corresponds to a cell of a connected figure ('F').
Gravity makes the figure fall through the field, until one of its cells reaches the ground, or meets an obstacle. Your task is to return the state of the field after the figure has fallen.
Note that it is guaranteed that the figure is connected, ie. between any two cells of the figure there exists a path which goes through the cells' sides (not through corners).

假设 input 为：

[
['F','F','F'],
['.','F','.'],
['.','F','F'],
['#','F','.'],
['F','F','.'],
['.','.','.'],
['.','.','#'],
['.','.','.']
]

那么希望 Output 为:

[
['.','.','.'],
['.','.','.'],
['F','F','F'],
['#','F','.'],
['.','F','F'],
['.','F','.'],
['F','F','#'],
['.','.','.']
]


*/
func figureUnderGravity(a [][]byte) [][]byte {
	m := len(a)
	n := len(a[0])

	/*	Step1 - Find by how many levels we have to shift the array downwards (In the below code 'shiftBy'), which is equal to the minimum of all the distances between 'F' and a '#' lying in the same column below and the distance between the lowest 'F' and ground.*/
	minToGround := math.MaxInt32
	minToObs := math.MaxInt32

	for j := 0; j < n; j++ {
		dist := 0
		for i := 0; i < m; i++ {
			if a[i][j] == '.' {
				dist++
			} else if a[i][j] == '#' {
				minToObs = min(dist, minToObs)
			} else if a[i][j] == 'F' {
				dist = 0
			}
		}

		minToGround = min(minToGround, dist)
	}
	move := min(minToGround, minToObs)

	//Step2 - Shift the array downwards by 'shiftBy' levels.
	b := make([][]byte, m)
	for i := 0; i < m; i++ {
		b[i] = make([]byte, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] == '#' {
				b[i][j] = '#'
				continue
			}

			if i-move < 0 {
				b[i][j] = '.'
			} else {
				if a[i-move][j] == '#' {
					b[i][j] = '.'
				} else {
					b[i][j] = a[i-move][j]
				}
			}
		}
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numOfWays(a, b []int, query [][]int) [][]int {
	h := make(map[int]bool)
	for _, n := range a {
		h[n] = true
	}

	ret := [][]int{}

	for _, q := range query {
		if q[0] == 0 {
			b[q[1]] = q[2]
		} else {
			for _, x := range b {
				y := q[1] - x
				_, ok := h[y]
				if ok {
					ret = append(ret, []int{y, x})
				}
			}
		}
	}

	return ret
}

/* bubble game
模拟消除游戏，input1 是一个 2d array，代表一个类似三消的棋盘
e.g.:
[1,2,3],
[1,2,4],
[1,2,2]
每个数字可以想象成对应一个颜色，input2也是一个2d array，每个array代表进行消除操作的坐标
e.g.:
[0,1]
[2,1]
[0,1]意味着尝试消除第1行，第2列的元素，消除的规则如下：如果元素本身和4向相邻的元素中数字相同的数量超过3个，则将这些元素消除，并让上方的元素降落，如果数量少于3个则什么都不发生。如果尝试消除的元素为0则什么都不发生。以尝试[0,1]举例：
[1,2,3],
[1,2,4],
[1,2,2]
红色的2是【0,1】对应的元素，蓝色的2是相邻且值相同的元素，因为只有2个相同值的元素，不构成消除条件，所以不消除，无事发生。
再看第二个input [2,1]
[1,2,3],
[1,2,4],
[1,2,2]
同样，红色的是选中的元素，蓝色的是相邻且值相同的元素，所以这三个数字会被消除。之后上方的元素会降落下来，所以最终的board长这样：
[1,0,0],
[1,0,3],
[1,2,4]
所有的0是原本位置元素降落导致的留空。
要求你return 一系列操作后最终board的state*/

func buggleGame(input1, input2 [][]int) {
	for _, o := range input2 {
		buggleGameUtil(input1, o[0], o[1])
		fmt.Println(input1)
	}
}

func buggleGameUtil(input [][]int, x, y int) {
	m := len(input)
	n := len(input[0])
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	cnt := 1

	for _, d := range dirs {
		r := x + d[0]
		c := y + d[1]
		if 0 <= r && r < m &&
			0 <= c && c < n && input[r][c] == input[x][y] {
			cnt++
		}
	}

	if cnt < 3 {
		return
	}

	if y-1 >= 0 && input[x][y-1] == input[x][y] {
		for i := x; i > 0; i-- {
			input[i][y-1] = input[i-1][y-1]
		}
		input[0][y-1] = 0
	}

	if y+1 >= 0 && input[x][y+1] == input[x][y] {
		for i := x; i > 0; i-- {
			input[i][y+1] = input[i-1][y+1]
		}
		input[0][y+1] = 0
	}

	start := x
	l := 1
	if x+1 < m && input[x+1][y] == input[x][y] {
		start = x + 1
		l++
	}
	if x-1 >= 0 && input[x-1][y] == input[x][y] {
		l++
	}

	for ; start-l >= 0; start-- {
		input[start][y] = input[start-l][y]
	}
	for ; start >= 0; start-- {
		input[start][y] = 0
	}
}

/*
We want to know what the Top Game is, defined by: The Top Game is the game users spent the most time in.
Each line of the file has the following information (comma separated):
- timestamp in seconds (long)
- user id (string)
- game id (int)
- action (string, either "join" or "quit")
e.g.
[
"1000000000,user1,1001,join", // user 1 joined game 1001
"1000000005,user2,1002,join", // user 2 joined game 1002
"1000000010,user1,1001,quit", // user 1 quit game 1001 after 10 seconds
"1000000020,user2,1002,quit", // user 2 quit game 1002 after 15 seconds
];
In this log,
The total time spent in game 1001 is 10 seconds.
The total time spent in game 1002 is 15 seconds.
Hence game 1002 is the Top Game. -> 1002
This file could be missing some lines of data (e.g. the user joined, but then the app crashed).
If data for a session (join to quit) is incomplete, please discard the session.
To recover some data, we attempt to estimate session length with this rule:
If a user joined a game but did not leave, assume they spent the minimum of
  - time spent before they joined a different game; and
  - average time spent across the same user's gaming sessions (from join to leave)
  e.g.
  "1500000000,user1,1001,join"
  "1500000010,user1,1002,join"
  "1500000015,user1,1002,quit"
  The user spent 5 seconds in game 2, so we assume they spent 5 seconds in game 1.
Write a function that returns the top game ID, given an array of strings representing
each line in the log file.

To recover some data, we attempt to estimate session length with this rule:If a user joined a game but did not leave, assume they spent the minimum of
  - time spent before they joined a different game; and
  - average time spent across the same user's gaming sessions (from join to leave)
  e.g.
  "1500000000,user1,1001,join"
  "1500000010,user1,1002,join"
  "1500000015,user1,1002,quit"
  The user spent 5 seconds in game 2, so we assume they spent 5 seconds in game 1.
Write a function that returns the top game ID, given an array of strings representing
*/

func topGame2(input []string) (string, int) {
	userGameStart := map[string]map[string]int{}
	gameTime := map[string]int{}
	userTime := map[string][]int{} //1st element is  total game time, 2nd element is the total number of game sessions
	lostUserGameTime := map[string]map[string][]int{}

	for _, line := range input {
		items := strings.Split(line, ",")
		ts, _ := strconv.Atoi(items[0])
		user := items[1]
		game := items[2]
		op := items[3]

		if op == "join" {
			if userGameStart[user] == nil {
				userGameStart[user] = map[string]int{}
			}

			for g, st := range userGameStart[user] {
				// // handle lost quit log
				duration := ts - st
				if lostUserGameTime[user] == nil {
					lostUserGameTime[user] = map[string][]int{}
				}
				if lostUserGameTime[user][g] == nil {
					lostUserGameTime[user][g] = []int{}
				}

				lostUserGameTime[user][game] = append(lostUserGameTime[user][game], duration)
				delete(userGameStart[user], g)
			}

			userGameStart[user][game] = ts
		} else {
			if _, ok := userGameStart[user][game]; ok {
				duration := ts - userGameStart[user][game]
				gameTime[game] += duration

				if userTime[user] == nil {
					userTime[user] = []int{0, 0}
				}
				userTime[user][0] += duration
				userTime[user][1]++

				delete(userGameStart[user], game)
			} else {
				// handle lost join log
			}
		}
	}

	for user, gt := range lostUserGameTime {
		userAvg := userTime[user][0] / userTime[user][1]
		for game, t := range gt {
			for _, tt := range t {
				if tt < userAvg {
					gameTime[game] += tt
				} else {
					gameTime[game] += userAvg
				}
			}
		}
	}

	topgame := ""
	topScore := 0
	for game, duration := range gameTime {
		if duration > topScore {
			topgame = game
			topScore = duration
		}
	}

	return topgame, topScore
}

func topGame(input []string) (string, int) {
	userGameStart := map[string]map[string]int{}
	gameTime := map[string]int{}

	for _, line := range input {
		items := strings.Split(line, ",")
		ts, _ := strconv.Atoi(items[0])
		user := items[1]
		game := items[2]
		op := items[3]

		if op == "join" {
			if userGameStart[user] == nil {
				userGameStart[user] = map[string]int{}
			}
			userGameStart[user][game] = ts
		} else {
			duration := ts - userGameStart[user][game]
			gameTime[game] += duration
		}
	}

	topgame := ""
	topScore := 0
	for game, duration := range gameTime {
		if duration > topScore {
			topgame = game
			topScore = duration
		}
	}

	return topgame, topScore
}

/*connect 4 +2: 实现connect4。
实现两个function，drop (int column， color)
 checkWin(row, column, color)。
只需要实现horizontal和vertical的时候判断能不能赢就好了，可以忽略对角线的情况。
*/
type connect4 struct {
	board [][]int
}

func newConnect4(m, n int) connect4 {
	b := make([][]int, m)
	for i := 0; i < m; i++ {
		b[i] = make([]int, n)
	}

	return connect4{board: b}
}

func (c connect4) Drop(column, color int) error {
	m := len(c.board)
	n := len(c.board[0])

	if !(0 <= column && column < n) {
		return fmt.Errorf("invalid column %v", column)
	}
	if c.board[0][column] != 0 {
		return fmt.Errorf("colum %v is full", column)
	}

	i := 0
	for i < m {
		if c.board[i][column] == 0 {
			i++
		} else {
			break
		}
	}

	c.board[i][column] = color

	return nil
}

func (c connect4) Check(row, column, color int) bool {
	cnt := 1
	i := row - 1
	for i >= 0 {
		if c.board[i][column] == color {
			cnt++
			i--
		} else {
			break
		}
	}

	i = row + 1
	for i < len(c.board[0]) {
		if c.board[i][column] == color {
			cnt++
			i++
		} else {
			break
		}
	}

	if cnt >= 4 {
		return true
	}

	cnt = 1
	j := column - 1
	for j >= 0 {
		if c.board[row][j] == color {
			cnt++
			j--
		} else {
			break
		}
	}

	j = column + 1
	for j < len(c.board) {
		if c.board[row][j] == color {
			cnt++
			j++
		} else {
			break
		}
	}

	if cnt >= 4 {
		return true
	}

	return false
}

/*
You are given an array of integers numbers. Your task is to count the number of distinct pairs (i, j) such that numbers[i] and numbers[j] have the samenumber of digits, and only one of the digits differ between numbers[i] and numbers[j]

E.g for numbers := []int{1, 151, 241, 1, 9, 22, 351}, the output should be 3
numbers[0] = 1 differs from numbers[4] = 9
numbers[1] = 151 differs from numbers[6] = 351
numbers[3] = 1 differs from numbers[4] = 9

*/
func almostEqualNumbers(numbers []int) int {
	cnt := 0
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			cnt += almostEqual(numbers[i], numbers[j])
		}
	}

	return cnt
}

func almostEqual(a, b int) int {
	if a == b {
		return 0
	}

	astr := strconv.Itoa(a)
	bstr := strconv.Itoa(b)

	if len(astr) != len(bstr) {
		return 0
	}

	cnt := 0
	for i, _ := range astr {
		if astr[i] != bstr[i] {
			cnt++
		}
	}

	if cnt == 1 {
		return 1
	}

	return 0
}

func findPeak(n []int) [][]int {
	ret := [][]int{}

	i := 0
	for i < len(n) {
		j := i
		for j < len(n) && n[j] == n[i] {
			j++
		}

		left := math.MinInt32
		right := math.MinInt32

		if i-1 >= 0 {
			left = n[i-1]
		}
		if j < len(n) {
			right = n[j]
		}

		if left < n[i] && n[i] > right {
			row := []int{}
			for k := i; k < j; k++ {
				row = append(row, k)
			}
			ret = append(ret, row)
		}

		i = j
	}

	return ret
}

func reversePair(n int) int {
	s := strconv.Itoa(n)
	var d string
	for i := 0; i < len(s); i += 2 {
		d += string(s[i+1])
		d += string(s[i])
	}

	r, _ := strconv.Atoi(d)
	return r
}

func check(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 || len(nums) == 2 {
		return true
	}

	i := 0

	for i < len(nums)-1 {
		if nums[i] > nums[i+1] {
			break
		}
		i++
	}

	if i == len(nums) {
		return true
	}

	for j := 0; j < i-1; j++ {
		if nums[j] > nums[j+1] {
			return false
		}
	}

	for j := i + 1; j < len(nums)-1; j++ {
		if nums[j] > nums[j+1] {
			return false
		}
	}

	if nums[0] < nums[len(nums)-1] {
		return false
	}

	return true
}

func numSub(s string) int {
	if s == "0" {
		return 0
	}
	cnt := 0
	for i, _ := range s {
		//if c == '0' {
		//	continue
		//}
		for j := i + 1; j < len(s)+1; j++ {
			if s[i:j] == "0" || strings.HasPrefix(s[i:j], "00") {
				continue
			}
			n, _ := strconv.Atoi(s[i:j])

			if n%3 == 0 {
				cnt++
			}
		}
	}

	return cnt
}

/*
Count pairs (i, j) from arrays arr[] & brr[] such that arr[i] – brr[j] = arr[j] – brr[i]

Given two arrays arr[] and brr[] consisting of N integers, the task is to count the number of pairs (i, j) from both the array such that (arr[i] – brr[j]) and (arr[j] – brr[i]) are equal.

Examples:

Input: A[] = {1, 2, 3, 2, 1}, B[] = {1, 2, 3, 2, 1}
Output: 2
Explanation: The pairs satisfying the condition are:

(1, 5): arr[1] – brr[5] = 1 – 1 = 0, arr[5[ – brr[1] = 1 – 1 = 0
(2, 4): arr[2] – brr[4] = 2 – 2 = 0, arr[4] – brr[2] = 2 – 2 = 0
Input: A[] = {1, 4, 20, 3, 10, 5}, B[] = {9, 6, 1, 7, 11, 6}
Output: 4

Efficient Approach: The idea is to transform the given expression (a[i] – b[j] = a[j] – b[i]) into the form (a[i] + b[i] = a[j] + b[j]) and then calculate pairs satisfying the condition. Below are the steps:

Transform the expression, a[i] – b[j] = a[j] – b[i] ==> a[i] + b[i] = a[j] +b[j]. The general form of expression becomes to count the sum of values at each corresponding index of the two arrays for any pair (i, j).
Initialize an auxiliary array c[] to store the corresponding sum c[i] = a[i] + b[i] at each index i.
Now the problem reduces to find the number of possible pairs having same c[i] value.
Count the frequency of each element in the array c[] and If any c[i] frequency value is greater than one then it can make a pair.
Count the number of valid pairs in the above steps using formula:
*/

func countPair(a, b []int) int {
	c := make(map[int]int)

	for i := 0; i < len(a); i++ {
		sum := a[i] + b[i]
		c[sum] += c[sum] + 1
	}

	cnt := 0
	for _, v := range c {
		cnt += v * (v - 1) / 2
	}

	return cnt
}

/*
For two arrays a and b of the same length, let's say a is a cyclic shift of b, if it's possible for a to become equal to b by performing cyclic shift operations on a.
E.g. {1,2,3,4,5}, {3,4,5,1,2} => true
     {1, 2, 3, 6, 5, 4}, {3, 6, 5, 4, 1, 2} => true

*/
func isCyclic(a, b []int) bool {
	n := len(a)

	i := 0
	for i < n {
		k := 0
		j := 0
		for i+k < n {
			if a[i+k] != b[j] {
				break
			}
			k++
			j++
		}
		if i+k == n {
			l := 0
			for l < n-k {
				if a[l] != b[k+l] {
					break
				}

				l++
			}

			if l == n-k {
				return true
			}
		}

		i++
	}

	return false
}

/*
The string chatOfPlayers represents a combination of the string "chat" from
different players. Any player can say "chat", and if multiple players say
"chat" at the same time, the letters of "chat" can overlap. Return the minimum
number of different players to finish all the "chat"s in the given string.
A valid "chat" means a player is printing 4 letters ‘c’, ’h’, ’a’, ’t’
sequentially. The players have to print all four letters to finish a
chat. If the given string is not a combination of valid "chat" return -1.

Example 1:
Input: chatOfPlayers = "chatchat"
Output: 1
Explanation: One player yelling "chat" twice.

Example 2:
Input: chatOfPlayers = "chcathat"
Output: 2
Explanation: The minimum number of players is two.
The first player could yell "CHcAThat".
The second player could yell later "chCatHAT".

Example 3:
Input: chatOfPlayers = "chatchht"
Output: -1
Explanation: The given string is an invalid combination of "chat" from different players.

思路:
- maintain a array to represent the count of each letter,  scan the string and at any time check if the account
    c >= h >= a >= t
    In the end, c == h == a == t
- c means a chatter start, t means a chatter end. the counter difference between c and t means the concurrent numbers

*/
func minChatter(s string) int {
	// array     c  h  a  t
	cnt := []int{0, 0, 0, 0}
	ret := math.MinInt32

	for _, c := range s {
		switch c {
		case 'c':
			cnt[0]++
			d := cnt[0] - cnt[3]
			if d > ret {
				ret = d
			}
		case 'h':
			cnt[1]++
			if cnt[1] > cnt[0] {
				return -1
			}
		case 'a':
			cnt[2]++
			if cnt[2] > cnt[1] {
				return -1
			}
		case 't':
			cnt[3]++
			if cnt[3] > cnt[2] {
				return -1
			}
		}
	}

	if !(cnt[0] == cnt[1] && cnt[1] == cnt[2] && cnt[2] == cnt[3]) {
		return -1
	}

	return ret
}

func findword(g [][]byte, s string) [][][]int {
	m := len(g)
	n := len(g[0])

	ret := [][][]int{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			visited := make([][]bool, m)
			for k := 0; k < m; k++ {
				visited[k] = make([]bool, n)
			}
			findworddfs(g, visited, s, i, j, [][]int{}, &ret)
		}
	}

	return ret
}

func findworddfs(g [][]byte, visited [][]bool, s string, x, y int, path [][]int, ret *[][][]int) {
	m := len(g)
	n := len(g[0])
	//if len(*ret) == 1 {
	//	return
	//}
	if s[0] != g[x][y] {
		return
	}

	path = append(path, []int{x, y})

	if len(s) == 1 {
		p := make([][]int, len(path))
		copy(p, path)
		*ret = append(*ret, path)
		return
	}

	visited[x][y] = true

	for _, d := range dirs {
		r := x + d[0]
		c := y + d[1]
		if 0 <= r && r < m &&
			0 <= c && c < n && visited[r][c] == false {
			findworddfs(g, visited, s[1:], r, c, path, ret)
		}
	}

	visited[x][y] = false
	path = path[0 : len(path)-1]

}

/*
A nonogram is a logic puzzle, similar to a crossword, in which the player is given a blank grid and has to color it according to some instructions. Specifically, each cell can be either black or white, which we will represent as 'B' for black and 'W' for white.

WWWW
BWWW
BWBB
WWBW
BBWW

For each row and column, the instructions give the lengths of contiguous runs of black ('B') cells.
For example, the instructions for one row of [ 2, 1 ] indicate that there must be a run of two black cells,
followed later by another run of one black cell, and the rest of the row filled with white cells.

These are valid solutions: [ W, B, B, W, B ] and [ B, B, W, W, B ] and also [ B, B, W, B, W ]
This is not valid: [ W, B, W, B, B ] since the runs are not in the correct order.
This is not valid: [ W, B, B, B, W ] since the two runs of Bs are not separated by Ws.
Your job is to write a function to validate a possible solution against a set of instructions.
Given a 2D matrix representing a player's solution; and instructions for each row along with additional instructions for each column; return True or False according to whether both sets of instructions match.
Example instructions #1
matrix1 = [
   [ W, W, W, W ],
   [ B, W, W, W ],
   [ B, W, B, B ],
   [ W, W, B, W ],
   [ B, B, W, W ]]
rows1_1  =[], [1], [1,2], [1], [2]
columns1_1 =[2,1], [1], [2], [1]
validateNonogram(matrix1, rows1_1, columns1_1) => True
Example solution matrix:
matrix1 ->
           row
      +------------+ instructions
      | WWWW | <-- []
      | BWWW | <-- [1]
      | BWBB | <-- [1,2]
      | WWBW | <-- [1]
      | BBWW | <-- [2]
      +------------+
      ^^^^
      ||||
column   [2,1] | [2] |
instructions  [1] [1]
Example instructions #2
(same matrix as above)
rows1_2  =[], [], [1], [1], [1,1]
columns1_2 =[2], [1], [2], [1]
validateNonogram(matrix1, rows1_2, columns1_2) => False
The second and third rows and the first column do not match their respective instructions.
Example instructions #3
(same matrix as above)
rows1_3  = [], [1], [3], [1], [2]
columns1_3 = [3], [1], [2], [1]
validateNonogram(matrix1, rows1_3, columns1_3) => False
The third row and the first column do not match their respective instructions.
Example instructions #4
(same matrix as above)
rows1_4  =[], [1,1], [1,2], [1], [2]
columns1_4 =[2,1], [1], [2], [1]
validateNonogram(matrix1, rows1_4, columns1_4) => False
The second row and the first column do not match their respective instructions
Example instructions #5
matrix2 = [
[ W, W ],
[ B, B ],
[ B, B ],
[ W, B ]
]
rows2_1  = [], [2], [2], [1]
columns2_1 = [1, 1], [3]
validateNonogram(matrix2, rows2_1, columns2_1) => False
The black cells in the first column are not separated by white cells.
Example instructions #6
(same matrix as above)
rows2_2  = [], [2], [2], [1]
columns2_2 = [3], [3]
validateNonogram(matrix2, rows2_2, columns2_2) => False
The first column has the wrong number of black cells.
Example instructions #7
(same matrix as above)
rows2_3  = [], [], [], []
columns2_3 = [], []
validateNonogram(matrix2, rows2_3, columns2_3) => False
All of the instructions are empty
n: number of rows in the matrix
m: number of columns in the matrix
*/
//func validateNonogram(matrix [][]byte, rows, cols [][]int) bool {
//	mrow := make([]string, len(matrix))
//	for i, m := range matrix {
//		mrow[i] = string(m)
//	}
//
//	mcol := make([]string, len(matrix[0]))
//
//	for j := 0; j < len(matrix[0]); j++ {
//		col := []byte{}
//		for i := 0; i < len(matrix); i++ {
//			col = append(col, matrix[i][j])
//		}
//		mcol[j] = string(col)
//	}
//
//	for i, r := range rows {
//		str := ""
//		for _, num := range r {
//			for k := 0; k < num; k++ {
//				str = append(str)
//			}
//		}
//	}
//}

// 给2个numerical string，同个digit求和，输出string。比如"99"和"999"就是返回"91818"。同样一个loop，补上剩余的即可。
func sums(a, b string) string {
	la := len(a)
	lb := len(b)

	ret := ""

	i := 0
	for la-1-i >= 0 && lb-1-i >= 0 {
		na := a[la-1-i] - '0'
		nb := b[lb-1-i] - '0'
		ret = strconv.Itoa(int(na+nb)) + ret
		i++
	}

	for la-1-i >= 0 {
		ret = string(a[la-1-i]) + ret
		i++
	}

	for lb-1-i >= 0 {
		ret = string(b[lb-1-i]) + ret
		i++
	}

	return ret
}

/*
  289. Game of Life (Medium)
  According to the Wikipedia's article: "The Game of Life, also known simply as Life, is a cellular automaton devised
  by the British mathematician John Horton Conway in 1970."

  Given a board with m by n cells, each cell has an initial state live (1) or dead (0). Each cell interacts with its
  eight neighbors (horizontal, vertical, diagonal) using the following four rules (taken from the above Wikipedia article):

  Any live cell with fewer than two live neighbors dies, as if caused by under-population.
  Any live cell with two or three live neighbors lives on to the next generation.
  Any live cell with more than three live neighbors dies, as if by over-population..
  Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
  Write a function to compute the next state (after one update) of the board given its current state.

  Follow up:
  Could you solve it in-place? Remember that the board needs to be updated at the same time: You cannot update some cells
  first and then use their updated values to update other cells.
  In this question, we represent the board using a 2D array. In principle, the board is infinite, which would cause problems
  when the active area encroaches the border of the array. How would you address these problems?
面试官说他想要的时间上的优化，比如有一大堆0，只有一点点1，要怎么快速生成下一个棋盘
Idea: Use a hashmap to represent the live cell.
Scan the hashmap, and for each live cell, check the adjacent cell to determine the next state.
Also for each adjacent dead cell, add the live count by one.
At the end scan the recorded dead cells' live count and determine whether we want to make it live
*/

func gameOfLife(board [][]int) {
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, 1}, {1, -1}, {-1, -1}}
	m := len(board)
	n := len(board[0])

	lives := map[int]map[int]int{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 1 {
				if _, ok := lives[i]; !ok {
					lives[i] = map[int]int{}
				}
				lives[i][j] = 1
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			nlive := 0
			for _, d := range dirs {
				r := i + d[0]
				c := j + d[1]

				if 0 <= r && r < m && 0 <= c && c < n && lives[r][c] == 1 {
					if lives[r] == nil || lives[r][c] == 0 {
						continue
					}
					nlive++
				}
			}
			if nlive < 2 {
				board[i][j] = 0
			} else if nlive == 3 {
				board[i][j] = 1
				continue
			} else if nlive > 3 {
				board[i][j] = 0
			}
		}
	}
}

// 68. Text Justification
func fullJustify(words []string, maxWidth int) []string {
	ret := []string{}

	line := []string{}
	totalWordLen := 0

	for _, w := range words {
		newLen := totalWordLen + len(w) + len(line)
		if newLen > maxWidth {
			ret = append(ret, distribute(line, maxWidth-totalWordLen))
			line = []string{w}
			totalWordLen = len(w)
		} else {
			line = append(line, w)
			totalWordLen += len(w)
		}
	}

	lastLine := strings.Join(line, " ")
	lastLine = lastLine + strings.Repeat(" ", maxWidth-len(lastLine))
	ret = append(ret, lastLine)

	for _, v := range ret {
		fmt.Printf("%v\n", v)
	}
	return ret
}

func distribute(line []string, width int) string {
	if len(line) == 1 {
		return line[0] + strings.Repeat(" ", width)
	}

	for i := 0; i < width; i++ {
		idx := i % (len(line) - 1)

		line[idx] = line[idx] + " "
	}

	return strings.Join(line, "")
}

type Hit struct {
	ts    int
	count int
}

type HitCounter struct {
	q []Hit
}

func newHitCounter() *HitCounter {
	return &HitCounter{q: []Hit{}}
}
func (h *HitCounter) hit(t int) {
	if len(h.q) == 0 {
		h.q = append(h.q, Hit{ts: t, count: 1})
	} else if h.q[len(h.q)-1].ts == t {
		h.q[len(h.q)-1].count++
	} else {
		h.q = append(h.q, Hit{ts: t, count: 1})
	}
}

func (h HitCounter) getHits(t int) int {
	for i, v := range h.q {
		if v.ts > t-300 {
			h.q = h.q[i:]
			break
		}
	}

	cnt := 0
	for _, v := range h.q {
		cnt += v.count
	}
	return cnt
}

func rendermessage(message [][]string, userWidth, width int) []string {
	ret := []string{}

	for _, m := range message {
		words := strings.Split(m[1], " ")
		padright := true
		if m[0] == "2" {
			padright = false
		}

		l := adjust(words, userWidth, width, padright)
		ret = append(ret, l...)
	}

	for _, v := range ret {
		fmt.Printf("%v\n", v)
	}

	return ret
}

func adjust(words []string, userWidth, width int, padright bool) []string {
	ret := []string{}

	s := strings.Join(words, " ")

	for len(s) > 0 {
		end := userWidth
		if end > len(s) {
			end = len(s)
		}
		ret = append(ret, pad(s[0:end], width, padright))
		s = s[end:]
	}

	return ret
}

func pad(s string, width int, padright bool) string {
	n := width - len(s)
	pads := strings.Repeat(" ", n)

	if padright {
		return strings.TrimPrefix(s+pads, " ")
	} else {
		return strings.TrimSuffix(pads+s, " ")
	}
}

func numsum(a, b string) string {
	ret := ""
	i := 0
	for i < len(a) && i < len(b) {
		da := a[len(a)-i-1] - '0'
		db := b[len(b)-i-1] - '0'
		ret = strconv.Itoa(int(da+db)) + ret

		i++
	}

	if i < len(a) {
		ret = a[:len(a)-i] + ret
	} else {
		ret = b[:len(b)-i] + ret
	}

	return ret
}

/*
621. Task Scheduler Medium

Given a characters array tasks, representing the tasks a CPU needs to do, where each letter represents a different task. Tasks could be done in any order. Each task is done in one unit of time. For each unit of time, the CPU could complete either one task or just be idle.

However, there is a non-negative integer n that represents the cooldown period between two same tasks (the same letter in the array), that is that there must be at least n units of time between any two same tasks.

Return the least number of units of times that the CPU will take to finish all the given tasks.



Example 1:

Input: tasks = ["A","A","A","B","B","B"], n = 2
Output: 8
Explanation:
A -> B -> idle -> A -> B -> idle -> A -> B
There is at least 2 units of time between any two same tasks.
Example 2:

Input: tasks = ["A","A","A","B","B","B"], n = 0
Output: 6
Explanation: On this case any permutation of size 6 would work since n = 0.
["A","A","A","B","B","B"]
["A","B","A","B","A","B"]
["B","B","B","A","A","A"]
...
And so on.
Example 3:

Input: tasks = ["A","A","A","A","A","A","B","C","D","E","F","G"], n = 2
Output: 16
Explanation:
One possible solution is
A -> B -> C -> A -> D -> E -> A -> F -> G -> A -> idle -> idle -> A -> idle -> idle -> A
*/
type task struct {
	t   byte
	cnt int
}

func leastInterval(tasks []byte, n int) int {
	results := []byte{}
	total := 0
	taskCountMap := make(map[byte]int)

	for _, t := range tasks {
		taskCountMap[t]++
	}

	taskCount := []task{}
	for k, v := range taskCountMap {
		taskCount = append(taskCount, task{t: k, cnt: v})
	}

	for {
		// Sort the taskCount array in decreasing order for each cycle
		sort.Slice(taskCount, func(i, j int) bool {
			if taskCount[i].cnt > taskCount[j].cnt {
				return true
			}
			return false
		})
		cycle := []byte{}

		for j := 0; j < len(taskCount); j++ {
			if taskCount[j].cnt != 0 {
				cycle = append(cycle, taskCount[j].t)
				taskCount[j].cnt--
				total++
				if len(cycle) == n+1 {
					break
				}
			}
		}

		if total == len(tasks) {
			results = append(results, cycle...)
			break
		}

		for len(cycle) < n+1 {
			cycle = append(cycle, '#')
		}
		results = append(results, cycle...)
	}

	fmt.Printf("%s\n", results)
	return len(results)
}

/*
  93. Restore IP Addresses (medium)
  Given a string containing only digits, restore it by returning
  all possible valid IP address combinations.

  For example:
  Given "25525511135",

  return ["255.255.11.135", "255.255.111.35"]. (Order does not matter)

  一个有效的IP地址由4个数字组成，每个数字在0-255之间。
  对于其中的2位数或3位数，不能以0开头。
  所以对于以s[i]开头的数字有3种可能：

*/
func restoreIpAddresses(s string) []string {
	parts := []string{}
	results := []string{}
	restoreIpAddressesDFS(s, parts, &results)

	return results
}

func restoreIpAddressesDFS(s string, parts []string, results *[]string) {
	if s == "" {
		if len(parts) == 4 {
			*results = append(*results, strings.Join(parts, "."))
		}
		return
	}

	for i := 0; i < 3 && i < len(s); i++ {
		substr := s[:i+1]
		if isValidPart(substr) {
			parts = append(parts, substr)
			restoreIpAddressesDFS(s[i+1:], parts, results)
			parts = parts[:len(parts)-1]
		}
	}
}

func isValidPart(s string) bool {
	if s[0] == '0' {
		return len(s) == 1
	}

	n, _ := strconv.Atoi(s)

	return 1 <= n && n <= 255
}

/*
  435. Non-overlapping Intervals
  Difficulty: Medium
  Given a collection of intervals, findRoot the minimum number of intervals you need to remove to make the rest of the
  intervals non-overlapping.

  Note:
  You may assume the interval's end point is always bigger than its start point.
  Intervals like [1,2] and [2,3] have borders "touching" but they don't overlap each other.
  Example 1:
  Input: [ [1,2], [2,3], [3,4], [1,3] ]

  Output: 1

  Explanation: [1,3] can be removed and the rest of intervals are non-overlapping.
  Example 2:
  Input: [ [1,2], [1,2], [1,2] ]

  Output: 2

  Explanation: You need to remove two [1,2] to make the rest of intervals non-overlapping.
  Example 3:
  Input: [ [1,2], [2,3] ]

  Output: 0

  Explanation: You don't need to remove any of the intervals since they're already non-overlapping.

  Interval schedule

  solution: https://medium.com/swlh/non-overlapping-intervals-f0bce2dfc617
*/
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})

	cnt := 0
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			cnt++
			if end > intervals[i][1] {
				end = intervals[i][1]
			}
		} else {
			end = intervals[i][1]
		}
	}

	return cnt
}

/*
  307. Range Sum Query - Mutable
  http://www.cnblogs.com/grandyang/p/4985506.html
*/

/*
  304. Range Sum Query 2D - Immutable
  Given a 2D matrix matrix, findRoot the sum of the elements inside the rectangle defined by its
  upper left corner (row1, col1) and lower right corner (row2, col2).

    Example:
    Given matrix = [
      [3, 0, 1, 4, 2],
      [5, 6, 3, 2, 1],
      [1, 2, 0, 1, 5],
      [4, 1, 0, 1, 7],
      [1, 0, 3, 0, 5]
    ]

    sumRegion(2, 1, 4, 3) -> 8
    sumRegion(1, 1, 2, 2) -> 11
    sumRegion(1, 2, 2, 4) -> 12
    Note:
    You may assume that the matrix does not change.
    There are many calls to sumRegion function.
    You may assume that row1 ≤ row2 and col1 ≤ col2.
*/
type NumMatrix struct {
	DP [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := len(matrix[0])

	mtx := make([][]int, m)
	for i := 0; i < m; i++ {
		mtx[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			up := 0
			left := 0
			diag := 0
			if i-1 >= 0 {
				up = mtx[i-1][j]
			}
			if j-1 >= 0 {
				left = mtx[i][j-1]
			}

			if i-1 >= 0 && j-1 >= 0 {
				diag = mtx[i-1][j-1]
			}

			mtx[i][j] = up + left - diag + matrix[i][j]
		}
	}

	return NumMatrix{
		DP: mtx,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	up := 0
	left := 0
	diag := 0
	if col1-1 >= 0 {
		left = this.DP[row2][col1-1]
	}
	if row1-1 >= 0 {
		up = this.DP[row1-1][col2]
	}

	if col1-1 >= 0 && row1-1 >= 0 {
		diag = this.DP[row1-1][col1-1]
	}

	return this.DP[row2][col2] - left - up + diag
}

func shrinkNumberline(points []int, k int) int {
	res := []int{}
	return shrinkNumberdfs(points, 0, k, &res)
}

//Shrinking Number Line: https://www.jianshu.com/p/c64ba6694545
func shrinkNumberdfs(points []int, i int, k int, res *[]int) int {
	if i == len(points) {
		min := math.MaxInt
		max := math.MinInt

		for _, v := range *res {
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}

		return max - min
	}

	*res = append(*res, points[i]+k)
	minc := shrinkNumberdfs(points, i+1, k, res)
	*res = (*res)[:len(*res)-1]

	*res = append(*res, points[i]-k)
	mdec := shrinkNumberdfs(points, i+1, k, res)
	*res = (*res)[:len(*res)-1]

	if minc < mdec {
		return minc
	}
	return mdec
}

// 427. Construct Quad Tree

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func constructQuadTree(grid [][]int) *Node {
	return quadTreeDfs(grid, 0, len(grid), 0, len(grid[0]))
}

func quadTreeDfs(grid [][]int, rowStart, rowEnd, colStart, colEnd int) *Node {
	for i := rowStart; i < rowEnd; i++ {
		for j := colStart; j < colEnd; j++ {
			if grid[i][j] != grid[rowStart][colStart] {
				return &Node{
					Val:         grid[rowStart][colStart] == 1,
					IsLeaf:      false,
					TopLeft:     quadTreeDfs(grid, rowStart, (rowStart+rowEnd)/2, colStart, (colStart+colEnd)/2),
					TopRight:    quadTreeDfs(grid, rowStart, (rowStart+rowEnd)/2, (colStart+colEnd)/2, colEnd),
					BottomLeft:  quadTreeDfs(grid, (rowStart+rowEnd)/2, rowEnd, colStart, (colStart+colEnd)/2),
					BottomRight: quadTreeDfs(grid, (rowStart+rowEnd)/2, rowEnd, (colStart+colEnd)/2, colEnd),
				}
			}
		}
	}

	return &Node{
		Val:         grid[rowStart][colStart] == 1,
		IsLeaf:      true,
		TopLeft:     nil,
		TopRight:    nil,
		BottomLeft:  nil,
		BottomRight: nil,
	}
}
