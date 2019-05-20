package alo

import (
	"sync"
	"testing"
)

func TestAtomicLock(t *testing.T) {
	const expected = 255

	var wg sync.WaitGroup
	var mu AtomicLock
	counter := 0

	for i := 0; i < expected; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer mu.Unlock()
			mu.Lock()
			counter++
		}()
	}

	wg.Wait()
	if counter != expected {
		t.Errorf("expecting %d, got %d", expected, counter)
	}
}
