package EventBus

import "sync"


type Array[T any] struct {
	m []T
	mu sync.RWMutex
}

func (a *Array[T]) Len() int {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return len(a.m)
}

func (a *Array[T]) Range(f func(index int, value T) bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for i, v := range a.m {
		if !f(i, v) {
			break
		}
	}
}

func (a *Array[T]) Get(index int) (T, bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if index < 0 || index >= len(a.m) {
		var zero T
		return zero, false
	}
	return a.m[index], true
}

func (a *Array[T]) Append(values ...T) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.m = append(a.m, values...)
}

func (a *Array[T]) Copy() Array[T] {
	a.mu.RLock()
	defer a.mu.RUnlock()
	cp := make([]T, len(a.m))
	copy(cp, a.m)
	return Array[T]{m: cp}
}

func (a *Array[T]) Delete(index int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if index < 0 || index >= len(a.m) {
		return
	}
	a.m = append(a.m[:index], a.m[index+1:]...)
}