package main

import "sync"

type InMemoryPlayerStore struct {
	store map[string]int
	mu    sync.Mutex
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}, sync.Mutex{}}
}

func (m *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return m.store[name]
}

func (m *InMemoryPlayerStore) RecordWin(name string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.store[name]++
}
