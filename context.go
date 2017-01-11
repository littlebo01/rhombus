package rhombus

import (
	"sync"
)

type storeItem struct {
	key string
	value interface {}
	doneChan chan int
}

type Context struct {
	store      map[string]interface{}
	storeChan  chan *storeItem
	exitChan   chan int
	wg         sync.WaitGroup

	strategies []Strategy
}

func New(strategies []Strategy) *Context {
	return &Context{
		store: make(map[string]interface{}),
		storeChan: make(chan *storeItem),
		exitChan: make(chan int),
		strategies: strategies,
	}
}

func (c *Context) Do() {
	go c.storeWriter()

	for _, strategy := range c.strategies {
		strategy.Do(c)
	}

	c.exitChan <- 1
	c.wg.Wait()
	close(c.storeChan)
}

func (c *Context) Get(key string) interface{} {
	return c.store[key]
}

func (c *Context) Set(key string, value interface{}) {
	c.store[key] = value
}

func (c *Context) DoSet(key string, value interface{}) {
	c.wg.Add(1)
	doneChan := make(chan int, 1)
	item := &storeItem{
		key: key,
		value: value,
		doneChan: doneChan,
	}

	select {
	case c.storeChan <- item:
	}

	<-doneChan
	close(doneChan)

	c.wg.Done()
}

func (c *Context) storeWriter() {
	for {
		select {
		case item := <-c.storeChan:
			c.store[item.key] = item.value
			item.doneChan <- 1
		case <-c.exitChan:
			goto exit
		}
	}

exit:
	close(c.exitChan)
}
