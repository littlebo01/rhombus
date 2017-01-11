package rhombus

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
	for i, strategy := range s.strategies {
		c.Add(1)

		go strategy.Do(c)

		if i % s.size == 0 {
			c.Wait()
		}
	}

	c.Wait()
}
