package rhombus

import (
	"sync"
)

type Context struct {
	store map[string]interface{}
	strategies []Strategy
	wg *sync.WaitGroup
}

func New(strategies []Strategy) *Context {
	return &Context{
		store: make(map[string]interface{}),
		strategies: strategies,
		wg: &sync.WaitGroup{},
	}
}

func (c *Context) Do() {
	for _, strategy := range c.strategies {
		strategy.Do(c)
	}
}

func (c *Context) Get(key string) interface{} {
	return c.store[key]
}

func (c *Context) Set(key string, value interface{}) {
	c.store[key] = value
}

func (c *Context) Wait() {
	c.wg.Wait()
}

func (c *Context) Add(delta int) {
	c.wg.Add(delta)
}

func (c *Context) Done() {
	c.wg.Done()
}
