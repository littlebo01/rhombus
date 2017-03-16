package rhombus

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
)

type Context struct {
	store      map[string]interface{}
	storeGuard sync.RWMutex
	err        error

	tasks []Task
}

const (
	stackSize = 4 << 10
)

func New(tasks []Task) *Context {
	return &Context{
		store: make(map[string]interface{}),
		tasks: tasks,
	}
}

func (c *Context) catcher() {
	if err := recover(); err != nil {
		stack := make([]byte, stackSize)
		size := runtime.Stack(stack, true)
		err2 := fmt.Errorf("[Task Recover] %s %s\n", err, stack[:size])

		c.Abort(err2)
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

func (c *Context) Sub(task Task) error {
	task.Do(c)

	if v, ok := task.(TaskError); ok {
		return v.Error()
	}

	return c.err
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

	c.store[key] = value

	c.storeGuard.Unlock()
}

func (c *Context) Del(key string) {
	if _, ok := c.store[key]; ok {
		c.storeGuard.Lock()
		delete(c.store, key)
		c.storeGuard.Unlock()
	}
}
