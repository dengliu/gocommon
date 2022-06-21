package main

import (
	"fmt"

	"github.com/dghubble/trie"
)

func main() {
	rt := trie.NewRuneTrie()

	mp := map[rune][]rune{'o': []rune{'0'}, 'l': []rune{'1', 'I'}}
	words := []string{}
	dfs(mp, "fool", "", &words)
	dfs(mp, "silly", "", &words)

	fmt.Printf("%v\n", words)

	for _, w := range words {
		rt.Put(w, "")
	}

	fmt.Printf("%v\n", findIllegale(rt, "si11y"))
	fmt.Printf("%v\n", findIllegale(rt, "applefoolbanana"))
	fmt.Printf("%v\n", findIllegale(rt, "applef00lbanana"))
	fmt.Printf("%v\n", findIllegale(rt, "applebanana"))
	fmt.Printf("%v\n", findIllegale(rt, "applef00lbananasil1y"))
	fmt.Printf("%v\n", findIllegale(rt, "silIyapplesi11yf00lbananasil1y"))

}

func findIllegale(rt *trie.RuneTrie, s string) []string {
	i := 0
	illegalWords := []string{}
	for i < len(s) {
		var prefix string

		rt.WalkPath(s[i:], func(key string, value interface{}) error {
			prefix = key
			return nil
		})

		if prefix == "" {
			i++
		} else {
			i += len(prefix)
			illegalWords = append(illegalWords, prefix)
		}
	}

	return illegalWords
}

func dfs(mp map[rune][]rune, s string, word string, words *[]string) {
	if s == "" {
		*words = append(*words, word)
		return
	}

	dfs(mp, s[1:], word+string(s[0]), words)

	for _, v := range mp[rune(s[0])] {
		dfs(mp, s[1:], word+string(v), words)
	}
}

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children map[int32]*TrieNode
	word     string
}

func Constructor() Trie {
	return Trie{root: &TrieNode{children: make(map[int32]*TrieNode)}}
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for _, c := range word {
		if cur.children[c] == nil {
			cur.children[c] = &TrieNode{children: make(map[int32]*TrieNode)}
		}

		cur = cur.children[c]
	}

	cur.word = word
}

func (this *Trie) Search(word string) bool {
	cur := this.root
	for _, c := range word {
		if cur.children[c] == nil {
			return false
		}

		cur = cur.children[c]
	}

	return cur.word == word
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for _, c := range prefix {
		if cur.children[c] == nil {
			return false
		}

		cur = cur.children[c]
	}

	return true
}
