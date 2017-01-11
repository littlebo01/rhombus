package rhombus

import (
	"sync"
)

func Batch(size int, strategies ...Strategy) Strategy {
	return &batchStrategy{
		size: size,
		strategies: strategies,
	}
}

type batchStrategy struct {
	size int
	strategies []Strategy
}

func (s *batchStrategy) Do(c *Context) {
	var wg sync.WaitGroup

	for i, strategy := range s.strategies {
		wg.Add(1)

		go strategy.Do(c)

		if i % s.size == 0 {
			wg.Wait()
		}
	}

	wg.Wait()
}
