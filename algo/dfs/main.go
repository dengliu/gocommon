package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Graph with weight on edge: usually use DFS

func main() {
	input := []string{
		"BTC,USD,48000,49000",
		"BTC,USD,47000,48500",
		"USD,EUR,0.85,0.86",
		"BTC,EUR,39000,42000",
	}
	bestRate(input)

	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	//queries := [][]string{{"b", "a"}}

	ret := calcEquation(equations, values, queries)
	fmt.Printf("%v\n", ret)

	board := make([][]byte, 1)
	board[0] = []byte{'a'}

	fmt.Println(mahjong("11122"))
}

/*There is a list of currency exchange quotes. Each quote consists of <currency1, currency2,
bid _price, ask price>. One can convert from 1 unit of currency1 to get <bid price> units of
currency2. Inversely, one can convert from <ask_price> units of currency2 to get 1 unit of
currency1. Find the best exchange rate from a source currency and a target currency.
Example list of quotes:
BTC,USD,48000,49000
BTC,USD,47000,48500
USD,EUR,0.85,0.86
BTC,EUR,39000,42000
From the above example, the best exchange rate from BTC to EUR is 40800 (48000 * 0.85),
with the exchange path is BTC -> USD -> EUR.
*/

func bestRate(input []string) float64 {
	graph := make(map[string]map[string]float64)

	for _, r := range input {
		record := strings.Split(r, ",")
		src := record[0]
		dest := record[1]
		bid, _ := strconv.ParseFloat(record[2], 32)
		ask, _ := strconv.ParseFloat(record[3], 32)

		if graph[src] == nil {
			graph[src] = map[string]float64{dest: bid}
		} else if graph[src][dest] < bid {
			graph[src][dest] = bid
		}

		if graph[dest] == nil {
			graph[dest] = map[string]float64{src: 1.0 / ask}
		} else if graph[dest][src] < 1.0/ask {
			graph[dest][src] = 1.0 / ask
		}
	}

	bestRate := 0.0

	bestRateDFS(graph, make(map[string]bool), "BTC", "EUR", 1.0, &bestRate)

	fmt.Printf("best rate %v\n", bestRate)

	bestRate = 0
	prs := []pathRate{}

	bestRateDFSWithPath(graph, make(map[string]bool), "BTC", "EUR", []string{}, 1.0, &prs)
	fmt.Printf("%v\n", prs)

	return bestRate
}

func bestRateDFS(graph map[string]map[string]float64,
	visited map[string]bool,
	start, end string,
	curRate float64,
	bestRate *float64) {

	if start == end {
		if curRate > *bestRate {
			*bestRate = curRate
		}
		return
	}

	visited[start] = true

	for next, val := range graph[start] {
		if visited[next] {
			continue
		}

		bestRateDFS(graph, visited, next, end, curRate*val, bestRate)
	}

	visited[start] = false
}

type pathRate struct {
	path []string
	rate float64
}

func bestRateDFSWithPath(graph map[string]map[string]float64,
	visited map[string]bool,
	begin, end string,
	path []string,
	rate float64,
	prs *[]pathRate) {
	if begin == end {
		path = append(path, end)
		*prs = append(*prs, pathRate{
			path: path,
			rate: rate,
		})

		return
	}

	visited[begin] = true
	path = append(path, begin)

	for next, r := range graph[begin] {
		if visited[next] {
			continue
		}
		bestRateDFSWithPath(graph, visited, next, end, path, rate*r, prs)
	}
	visited[begin] = false
	path = path[0 : len(path)-1]
}

/*
  399. Evaluate Division
  Equations are given in the format A / B = k, where A and B are variables represented as strings, and k is a real number (floating point number). Given some queries, return the answers. If the answer does not exist, return -1.0.

Example:
Given a / b = 2.0, b / c = 3.0.
queries are: a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ? .
return [6.0, 0.5, -1.0, 1.0, -1.0 ].

The input is: vector<pair<string, string>> equations, vector<double>& values, vector<pair<string, string>> queries , where equations.size() == values.size(), and the values are positive. This represents the equations. Return vector<double>.

According to the example above:

equations = [ ["a", "b"], ["b", "c"] ],
values = [2.0, 3.0],
queries = [ ["a", "c"], ["b", "a"], ["a", "e"], ["a", "a"], ["x", "x"] ].


The input is always valid. You may assume that evaluating the queries will result in no division by zero and there is no contradiction.

graph, bi-direction graph with weight on edge, using map<String, Map<String, Double>>, point -> (point, weight) to represent it:
  - a -> (b -> value)
  - b -> (a -> 1.0 / value)

wordSearchDfs, visited (set), pass pre-calculated weight (multiply result)
  - avoid cycle
conditions to stop:
  - map doesn’t contains either a or b, return -1.0 (stop without finding result)
  - visited.size() == map.size(), return -1.0 (stop without finding result))
  - a.equals(b) return value; with result)
*/
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)

	// build graph
	for i, e := range equations {
		if graph[e[0]] == nil {
			graph[e[0]] = map[string]float64{}
		}

		graph[e[0]][e[1]] = values[i]

		if graph[e[1]] == nil {
			graph[e[1]] = map[string]float64{}
		}

		graph[e[1]][e[0]] = 1.0 / values[i]
	}

	ret := make([]float64, len(queries))
	for i, q := range queries {
		// wordSearchDfs
		r := calcEquationDFS(graph, make(map[string]bool), q[0], q[1], 1)
		fmt.Printf("%s->%s: %f\n", q[0], q[1], r)
		ret[i] = r
	}

	return ret
}

func calcEquationDFS(graph map[string]map[string]float64,
	visited map[string]bool,
	start, end string,
	tmp float64) float64 {
	if graph[start] == nil {
		return -1
	}

	if start == end {
		return tmp
	}

	visited[start] = true
	for next, val := range graph[start] {
		if visited[next] {
			continue
		}
		v := calcEquationDFS(graph, visited, next, end, tmp*val)
		if v != -1.0 {
			return v
		}
	}

	visited[start] = false

	return -1.0
}

/*79. Word Search (Medium)
Given a 2D board and a word, findRoot if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.

For example,
Given board =

[
['A','B','C','E'],
['S','F','C','S'],
['A','D','E','E']
]
word = "ABCCED", -> returns true,
word = "SEE", -> returns true,
word = "ABCB", -> returns false.

思路：

以board上的每个cell为出发点，利用depth first search向上下左右四个方向搜索匹配word。
搜索的时候要考虑board的边界，cell是否已经在DFS的路径上被用过，以及cell上的值是否与word的当前字符匹配。
为了跟踪DFS的路径避免同一个cell被访问多次，使用一个visited矩阵来代表board上任意的cell(i, j)是否已经被访问过。
*/
func wordSearch(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			visited := make([][]bool, m)
			for k := 0; k < m; k++ {
				visited[k] = make([]bool, n)
			}

			if wordSearchDfs(board, visited, word, i, j) {
				return true
			}
		}

	}
	return false

}

func wordSearchDfs(board [][]byte, visited [][]bool, word string, x, y int) bool {
	m := len(board)
	n := len(board[0])

	if word == "" {
		return true
	}

	if len(word) == 1 && word[0] == board[x][y] {
		return true
	}

	if word[0] != board[x][y] {
		return false
	}

	visited[x][y] = true

	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, d := range dir {
		r := x + d[0]
		c := y + d[1]

		if 0 <= r && r < m && 0 <= c && c < n && !visited[r][c] {
			if v := wordSearchDfs(board, visited, word[1:], r, c); v {
				return true
			}
		}

	}
	visited[x][y] = false

	return false

}

func letterCombinations(digits string) []string {
	dict := map[byte]string{}
	dict['1'] = ""
	dict['2'] = "abc"
	dict['3'] = "def"
	dict['4'] = "ghi"
	dict['5'] = "jkl"
	dict['6'] = "mno"
	dict['7'] = "pqrs"
	dict['8'] = "tuv"
	dict['9'] = "wxyz"

	ret := []string{}

	letterCombinationsDfs(dict, digits, "", &ret)

	return ret

}

func letterCombinationsDfs(dict map[byte]string, digits string, path string, results *[]string) {
	if digits == "" {
		if path != "" {
			*results = append(*results, path)
		}

		return
	}

	for _, c := range dict[digits[0]] {
		letterCombinationsDfs(dict, digits[1:], path+string(c), results)
	}
}

/*
设计一个类似麻将的游戏规则：
如果有a number of tiles,  0-9. The tiles can be grouped into pairs or triples of the same tile. 举一例，"33344466"有a triple of 3s, a triple of 4s, and a pair of 6s. 再举一例, "55555777" has a triple of 5s, a pair of 5s, and a triple of 7s.
A "complete hand" 就是说你的牌 可以group into any number of triples (zero or more) 和有且仅有一个 pair, and each tile is used in exactly one triple or pair.
写一个boolean function, 如果是complete hand, 返回true, 如果不是，返回false.
给了一堆test case:
tiles1 = "44433222"           # True. 444 33 222
tiles2 = "111333555"          # False. 多个triple 但是没有pair.
tiles3 = "00000111"           # True.
tiles4 = "13233121"           # True.
tiles6 = "99999999"           # True.
tiles7 = "999999999"          # False.
tiles9 = "77"                 # True.  One .
tiles12 = "332"               # False
tiles16 = "1111122222"        # False.


3. same as 1. but we can have 3 consecutives
          "56788822"          # True
*/

func mahjong1(s string) bool {
	m := make(map[int]int)
	single := []int{}

	pair := 0
	for _, c := range s {
		m[int(c)] = m[int(c)] + 1
	}

	for k, v := range m {
		if v%3 == 1 {
			single = append(single, k)
			//return false
		}
		if v%3 == 2 {
			pair++
		}
	}

	sort.Ints(single)

	for i := 0; i < len(single)-1; i++ {
		if single[i] > single[i+1] {
			return false
		}
	}

	if pair == 1 {
		return true
	}

	return false
}

func mahjong(s string) bool {
	state := make(map[int]int)

	// 统计每个数字出现的次数
	for _, c := range s {
		state[int(c)]++
	}

	return mahjongdfs(state, len(s), false)

}

func mahjongdfs(state map[int]int, len int, hasHead bool) bool {
	if len == 0 {
		return true
	}

	if !hasHead {
		//遍历个数大于2的数字，让它作为雀头，判断后面的规则是否通过
		for k, v := range state {
			if v >= 2 {
				state[k] -= 2
				if ok := mahjongdfs(state, len-2, true); ok {
					return true
				}
				state[k] += 2
			}
		}
	} else {
		for k, v := range state {
			//判断刻子
			if v >= 3 {
				state[k] -= 3
				if ok := mahjongdfs(state, len-3, true); ok {
					return true
				}
				state[k] += 3
			}

			//判断顺子
			if k+2 < '9' && state[k+1] >= 1 && state[k+2] >= 1 {
				state[k]--
				state[k+1]--
				state[k+2]--

				if ok := mahjongdfs(state, len-3, true); ok {
					return true
				}
				state[k]++
				state[k+1]++
				state[k+2]++
			}
		}
	}

	return false
}

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

/*
1120. Maximum Average Subtree

Given the root of a binary tree, find the maximum average value of any subtree of that tree.

 (A subtree of a tree is any node of that tree plus all its descendants.
 The average value of a tree is the sum of its values, divided by the number of nodes.)


 Example 1:

 Input: [5,6,1]
 Output: 6.00000
 Explanation:
     For the node with value = 5 we have an average of (5 + 6 + 1) / 3 = 4.
     For the node with value = 6 we have an average of 6 / 1 = 6.
     For the node with value = 1 we have an average of 1 / 1 = 1.
     So the answer is 6 which is the maximum.
*/

func maximumAverageSubtree(r *TreeNode) int {
	maxavg := 0

	maximumAverageSubtreeDFS(r, &maxavg)

	return maxavg
}

func maximumAverageSubtreeDFS(r *TreeNode, maxAve *int) (int, int) {
	if r == nil {
		return 0, 0
	}

	ltotal, lcount := maximumAverageSubtreeDFS(r.left, maxAve)
	rtotal, rcount := maximumAverageSubtreeDFS(r.right, maxAve)

	total := ltotal + r.val + rtotal
	val := lcount + rcount + 1

	avg := total / val
	if avg > *maxAve {
		*maxAve = avg
	}

	return total, val
}
