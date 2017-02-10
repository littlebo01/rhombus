package rhombus

import (
	"sync"
	"errors"
)

type Context struct {
	store      map[string]interface{}
	storeGuard sync.RWMutex
	done       bool
	Error      error

	tasks []Task
}

func New(tasks []Task) *Context {
	return &Context{
		store: make(map[string]interface{}),
		done: false,
		tasks: tasks,
	}

}

func (c *Context) Do() {
	for _, task := range c.tasks {
		task.Do(c)

		if c.done {
			break
		}
	}
}

func (c *Context) Abort(msg interface{}) {
	switch err := msg.(type) {
	case error:
		c.Error = err
	case string:
		c.Error = errors.New(err)
	default:
		panic(errors.New("Unsupported type"))
	}

	c.done = true
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
