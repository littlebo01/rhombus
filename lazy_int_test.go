package rhombus

import (
	"testing"
)

type lazyValue interface {
	Resolve(c *Context)
}

type lazyTask struct {
	V lazyValue
}

func (t *lazyTask) Do(c *Context) {
	t.V.Resolve(c)
}

func (t *lazyTask) Value() interface{} {
	return t.V
}

func newLazyIntTask(name string) Task {
	return &lazyTask{NewInt(name)}
}

func TestLazyInt(t *testing.T) {
	tasks := []Task{
		Value("1", newLazyIntTask("a")),
		Value("2", newLazyIntTask("b")),
		Value("3", newLazyIntTask("c")),
		Value("4", newLazyIntTask("d")),
	}

	items := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}

	pairs := map[string]string {
		"a": "1",
		"b": "2",
		"c": "3",
		"d": "4",
	}

	c := New(tasks)
	for key, value := range(items) {
		c.Set(key, value)
	}

	c.Do()

	for key1, key2 := range(pairs)  {
		value := items[key1]
		except := c.Get(key2).(*LazyInt)

		if value != except.Val {
			t.Fatalf("Val (%g) %+v ; want (%+v)", value, value, except)
		}
	}
}
