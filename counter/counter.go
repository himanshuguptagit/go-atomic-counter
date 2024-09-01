package counter

import "sync"

type Counter struct {
	count int
	mu    sync.Mutex
}

func NewCounter() *Counter {

	return &Counter{}

}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c Counter) GetCount() int {
	return c.count
}
