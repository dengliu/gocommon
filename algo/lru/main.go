package main

import (
	"container/list"
	"fmt"
)

type KV struct {
	k string
	v any
}

type LRUCache struct {
	l   *list.List
	m   map[string]*list.Element
	cap int
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		l:   list.New(),
		m:   map[string]*list.Element{},
		cap: cap,
	}
}

func (c *LRUCache) Get(k string) (any, bool) {
	e, ok := c.m[k]
	if !ok {
		return nil, false
	}

	c.l.MoveToBack(e)

	return e.Value.(KV).v, true
}

func (c *LRUCache) Set(k string, v any) {
	e, ok := c.m[k]
	if ok {
		e.Value = KV{k, v}
		c.l.MoveToBack(e)
	} else {
		e = c.l.PushBack(KV{k, v})
		c.m[k] = e

		if c.l.Len() > c.cap {
			front := c.l.Front()
			c.l.Remove(front)
			delete(c.m, front.Value.(KV).k)
		}
	}
}

func (c *LRUCache) Del(k string) bool {
	e, ok := c.m[k]
	if !ok {
		return false
	} else {
		c.l.Remove(e)
		delete(c.m, k)

		return true
	}
}

func main() {
	cache := NewLRUCache(3)
	fmt.Println(cache.Get("a"))

	cache.Set("a", "a1")
	fmt.Println(cache.Get("a"))
	cache.Set("b", "b1")
	fmt.Println(cache.Get("b"))
	cache.Set("c", "c1")
	fmt.Println(cache.Get("c"))
	cache.Set("d", "d1")

	fmt.Println(cache.Get("a"))
	fmt.Println(cache.Get("b"))
	fmt.Println(cache.Get("c"))
	fmt.Println(cache.Get("d"))

	fmt.Println(cache.Get("b"))
	cache.Set("e", "e1")

	fmt.Println(cache.Get("b"))
	fmt.Println(cache.Get("c"))

	cache.Del("b")
	fmt.Println(cache.Get("b"))
}
