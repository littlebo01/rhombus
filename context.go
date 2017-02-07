package rhombus

import (
	"sync"
)

type Context struct {
	store      map[string]interface{}
	storeGuard sync.RWMutex
	running    bool

	tasks []Task
}

func New(tasks []Task) *Context {
	return &Context{
		store: make(map[string]interface{}),
		running: true,
		tasks: tasks,
	}

}

func (c *Context) Do() {
	for _, task := range c.tasks {
		if c.running {
			task.Do(c)
		}
	}
}

func (c *Context) Abort() {
	c.running = false
}

func (c *Context) Close() {
}

func (c *Context) Get(key string) interface{} {
	c.storeGuard.RLock()
	defer c.storeGuard.RUnlock()

	return c.store[key]
}

func (c *Context) Set(key string, value interface{}) {
	c.storeGuard.Lock()
	defer c.storeGuard.Unlock()

	c.store[key] = value
}
