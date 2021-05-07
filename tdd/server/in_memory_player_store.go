package main

import "sync"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}, sync.Mutex{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
	mu    sync.Mutex
}

func (m *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return m.store[name]
}

func (m *InMemoryPlayerStore) RecordWin(name string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.store[name]++
}

func (m *InMemoryPlayerStore) GetLeague() League {
	var league League

	for name, wins := range m.store {
		league = append(league, Player{name, wins})
	}
	return league
}
