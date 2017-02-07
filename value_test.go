package rhombus

import (
	"testing"
)

func TestValueWith(t *testing.T) {
	tasks := []Task{
		Multi(
			ValueWith("1", func(c *Context) interface{} { return 1 }),
			ValueWith("2", func(c *Context) interface{} { return 2 }),
			ValueWith("3", func(c *Context) interface{} { return 3 }),
			ValueWith("4", func(c *Context) interface{} { return 4 }),
		),
	}

	c := New(tasks)
	c.Do()

	contextAssert(t, c, 4)
}
