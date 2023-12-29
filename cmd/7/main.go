package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type SyncMap[K comparable, V any] struct {
	mu sync.RWMutex
	m map[K]V
}

func New[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{
		m: make(map[K]V),
	}
}

func (sm *SyncMap[K, V]) Get(key K) (V, bool) {
	sm.mu.RLock()
	v, ok := sm.m[key]
	sm.mu.RUnlock()

	return v, ok
}

func (sm *SyncMap[K, V]) Set(key K, value V) {
	sm.mu.Lock()
	sm.m[key] = value
	sm.mu.Unlock()
}

// Реализовать конкурентную запись данных в map.
func main() {
	m := New[any, any]()

	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			m.Set(rand.Intn(10), rand.Intn(10))
		}()
	}

	wg.Wait()

	fmt.Printf("%v\n", m.m)
}
