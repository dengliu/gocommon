package main

import "sync"

type kvmap map[string]interface{}

type TNXKVStore struct {
	queue []kvmap
	mu    sync.Mutex
}

func NewTxnKVStore() *TNXKVStore {
	return &TNXKVStore{
		queue: []kvmap{kvmap{}},
	}
}

func (s *TNXKVStore) Get(k string) (interface{}, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	val, ok := s.queue[len(s.queue)-1][k]

	return val, ok
}

func (s *TNXKVStore) Set(k string, v interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.queue[len(s.queue)-1][k] = v
}

func (s *TNXKVStore) Del(k string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.queue[len(s.queue)-1], k)
}

func (s *TNXKVStore) BeginTxn() {
	snapshot := kvmap{}
	for k, v := range s.queue[len(s.queue)-1] {
		snapshot[k] = v
	}

	s.queue = append(s.queue, snapshot)
}

func (s *TNXKVStore) EndTxn() {
	if len(s.queue) == 1 {
		return
	}

	s.queue = s.queue[0 : len(s.queue)-1]
}

func (s *TNXKVStore) RollBack() {
	s.queue[len(s.queue)-1] = kvmap{}

	for k, v := range s.queue[len(s.queue)-2] {
		s.queue[len(s.queue)-1][k] = v
	}
}

func (s *TNXKVStore) Commit() {
	if len(s.queue) == 1 {
		return
	}

	for k, v := range s.queue[len(s.queue)-1] {
		s.queue[len(s.queue)-2][k] = v
	}
}

func main() {
	kv := NewTxnKVStore()

	kv.Set("a", 2)
	v, _ := kv.Get("a")
	println(v.(int))

	kv.Set("a", 3)
	v, _ = kv.Get("a")
	println(v.(int))

	kv.BeginTxn()
	kv.Set("a", 4)
	v, _ = kv.Get("a")
	println(v.(int))

	kv.RollBack()
	v, _ = kv.Get("a")
	println(v.(int))

	kv.BeginTxn()
	kv.Set("a", 4)
	v, _ = kv.Get("a")
	println(v.(int))

	kv.Commit()
	v, _ = kv.Get("a")
	println(v.(int))

	kv.BeginTxn()
	kv.Set("a", 5)
	v, _ = kv.Get("a")
	println(v.(int))

	kv.EndTxn()
	v, _ = kv.Get("a")
	println(v.(int))

}
