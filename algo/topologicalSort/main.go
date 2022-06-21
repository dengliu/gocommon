package main

import "fmt"

/*
PREREQUISITES = {
    "cs112": ["cs101", "cs102"],
    "cs132": ["cs112", "cs105"],
    "cs243": ["cs101", "cs112"],
    "cs229": ["cs101"],
    "cs224n": ["cs107", "cs229"],
    "cs230": ["cs229", "cs101"],
    "cs229t": ["cs101", "cs107"],
    "cs224u": ["cs224n"],
    "cs300" : ["cs230"]
}

NLP_COURSES = ["cs300", "cs229", "cs230", "cs224n", "cs229t", "cs224u"]

这个问题需要先正向 bfs 找到dependecy graph 中的 connected component，
再reverse graph 用 topological sort 找‍顺序，这里要注意有环的情况要单独判断
*/

func main() {
	preReq := map[string][]string{}
	preReq["cs112"] = []string{"cs101", "cs102"}
	preReq["cs132"] = []string{"cs112", "cs105"}
	preReq["cs243"] = []string{"cs101", "cs112"}
	preReq["cs229"] = []string{"cs101"}
	preReq["cs224n"] = []string{"cs107", "cs229"}
	preReq["cs230"] = []string{"cs229", "cs101"}
	preReq["cs229t"] = []string{"cs101", "cs107"}
	preReq["cs224u"] = []string{"cs224n"}
	preReq["cs300"] = []string{"cs230"}

	courses := []string{"cs300", "cs229", "cs230", "cs224n", "cs229t", "cs224u"}

	fmt.Printf("%v\n", generateOrder(preReq, courses))
	fmt.Println(alianDictionary([]string{"wrt",
		"wrf",
		"er",
		"ett",
		"rftt"}))
}

func generateOrder(prereq map[string][]string, courses []string) []string {
	inDegree := map[string]int{}
	graph := map[string][]string{}
	relevantCourse := map[string]bool{}

	// 1. BFS traverse dependence graph to get the relevant courses
	for len(courses) != 0 {
		c := courses[0]
		courses = courses[1:]

		relevantCourse[c] = true
		for _, p := range prereq[c] {
			if !relevantCourse[p] {
				courses = append(courses, p)
			}
		}
	}

	// 2. Build indegree and graph
	for c, _ := range relevantCourse {
		inDegree[c] = 0
		graph[c] = []string{}
	}

	for c, _ := range relevantCourse {
		for _, p := range prereq[c] {
			inDegree[c]++
			graph[p] = append(graph[p], c)
		}
	}

	// 3. Topological sort
	queue := []string{}
	for k, v := range inDegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	ret := []string{}
	for len(queue) != 0 {
		c := queue[0]
		queue = queue[1:]

		ret = append(ret, c)

		for _, next := range graph[c] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	return ret
}

// Leetcode 210 Course Schedule II
func findOrder(numCourses int, prerequisites [][]int) []int {
	inDegree := make([]int, numCourses)
	adjList := make([][]int, numCourses)

	for i := 0; i < numCourses; i++ {
		adjList[i] = []int{}
	}

	for _, v := range prerequisites {
		inDegree[v[0]]++
		adjList[v[1]] = append(adjList[v[1]], v[0])
	}

	q := []int{}
	for i, v := range inDegree {
		if v == 0 {
			q = append(q, i)
		}
	}

	ret := []int{}

	for len(q) != 0 {
		c := q[0]
		q = q[1:]
		ret = append(ret, c)

		for _, v := range adjList[c] {
			inDegree[v]--
			if inDegree[v] == 0 {
				q = append(q, v)
			}
		}
	}

	return ret
}

/*
269. Alien Dictionary

There is a new alien language which uses the latin alphabet. However, the order among letters are unknown to you.
You receive a list of words from the dictionary, where words are sorted lexicographically by the rules of this
new language. Derive the order of letters in this language.

For example,
Given the following words in dictionary,

[
"wrt",
"wrf",
"er",
"ett",
"rftt"
]

The correct order is: "wertf".

Note:

You may assume all letters are in lowercase.
If the order is invalid, return an empty string.
There may be multiple valid order of letters, return any one of them is fine.
*/
func alianDictionary(input []string) string {
	graph := map[byte][]byte{}
	indegree := map[byte]int{}

	for _, s := range input {
		for i := 0; i < len(s)-1; i++ {
			if graph[s[i]] == nil {
				graph[s[i]] = []byte{}
			}

			if s[i] != s[i+1] {
				graph[s[i]] = append(graph[s[i]], s[i+1])
				indegree[s[i+1]]++
			}
		}
		if _, ok := indegree[s[0]]; !ok {
			indegree[s[0]] = 0
		}
	}

	q := []byte{}
	for k, v := range indegree {
		if v == 0 {
			q = append(q, k)
		}
	}

	ret := ""
	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		ret += string(c)
		for _, next := range graph[c] {
			indegree[next]--
			if indegree[next] == 0 {
				q = append(q, next)
			}
		}
	}

	return ret
}
