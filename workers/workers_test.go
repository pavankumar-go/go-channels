package workers

import (
	"sync"
	"testing"
	"time"
)

func TestWorkers(t *testing.T) {
	wg := New(50) // runs 50 workers
	var mu sync.Mutex
	p := 0

	for i := 0; i < 1000; i++ {
		wg.Add(func() {
			time.Sleep(5 * time.Millisecond)
			mu.Lock()
			p++
			mu.Unlock()
		})
	}
	wg.Wait()

	if p != 1000 {
		t.Error("Expected 1000, but got ", p)
	}
}
