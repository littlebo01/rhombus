package rhombus

import (
	"testing"
)

type lazyTask struct {
	V  Lazy
	id int
}

func (t *lazyTask) Do(c *Context) {
	t.V.Resolve(c, &t.id)
}

func (t *lazyTask) Value() interface{} {
	return t.id
}

func newLazyTask(name string) Task {
	return &lazyTask{V: Lazy(name)}
}

func TestLazy(t *testing.T) {
	tasks := []Task{
		Value("1", newLazyTask("a")),
		Value("2", newLazyTask("b")),
		Value("3", newLazyTask("c")),
		Value("4", newLazyTask("d")),
	}

	c := New(tasks)
	c.Set("a", 1)
	c.Set("b", 2)
	c.Set("c", 3)
	c.Set("d", 4)

	c.Do()

	contextAssert(t, c, 4)
}
