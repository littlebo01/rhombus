package rhombus

import (
	"testing"
)

type lazyIntTask struct {
	Id *LazyInt
}


func (t *lazyIntTask) Do(c *Context) {
	t.Id.Resolve(c)
}

func (t *lazyIntTask) Value() interface{} {
	return t.Id.Val
}

func newLazyIntTask(name string) Task {
	return &lazyIntTask{NewInt(name)}
}

func TestLazyInt(t *testing.T) {
	tasks := []Task{
		Value("1", newLazyIntTask("a")),
		Value("2", newLazyIntTask("b")),
		Value("3", newLazyIntTask("c")),
		Value("4", newLazyIntTask("d")),
	}

	c := New(tasks)
	c.Set("a", 1)
	c.Set("b", 2)
	c.Set("c", 3)
	c.Set("d", 4)

	c.Do()

	contextAssert(t, c, 1)
}

func TestLazyInts(t *testing.T) {
}


func TestLazyString(t *testing.T) {

}

func TestLazyStrings(t *testing.T) {

}
