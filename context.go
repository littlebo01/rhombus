package rhombus

import (
	"errors"
	"sync"
)

type Context struct {
	store      map[string]interface{}
	storeGuard sync.RWMutex
	err        error

	tasks []Task
}

func New(tasks []Task) *Context {
	return &Context{
		store: make(map[string]interface{}),
		tasks: tasks,
	}

}

func (c *Context) Do() error {
	for _, task := range c.tasks {
		task.Do(c)

		if c.err != nil {
			return c.err
		}
	}

	return nil
}

func (c *Context) Abort(msg interface{}) {
	switch err := msg.(type) {
	case error:
		c.err = err
	case string:
		c.err = errors.New(err)
	default:
		panic(errors.New("Unsupported type"))
	}
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
