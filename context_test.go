package rhombus

import (
	"strconv"
	"testing"
)

type setTaskParams struct {
	value interface{}
}

type setTask struct {
	params *setTaskParams
}

type abortTask struct{}
type panicTask struct{}

func newSetTask(i int) Task {
	return Value(
		strconv.Itoa(i),
		&setTask{&setTaskParams{
			value: i,
		}},
	)
}

func (s *setTask) Do(c *Context) {}

func (s *setTask) Value() interface{} {
	return s.params.value
}

func (s *abortTask) Do(c *Context) {
	c.Abort("aborted")
}

func (s *abortTask) Value() interface{} {
	return nil
}

func (s *panicTask) Do(c *Context) {
	a := make([]int, 0, 0)
	_ = a[100]
}

func (s *panicTask) Value() interface{} {
	return nil
}

func contextAssert(t *testing.T, c *Context, top int) {
	for i := 1; i <= top; i++ {
		key := strconv.Itoa(i)
		val := c.Get(key)

		if val.(int) != i {
			t.Fatalf("Val(%g) = %d; want %d", val, val, i)
		}
	}
}

func TestContextSet(t *testing.T) {
	tasks := []Task{}
	c := New(tasks)
	c.Set("1", 1)
	c.Set("2", 2)
	c.Set("3", 3)
	c.Set("4", 4)

	contextAssert(t, c, 4)
}

func TestContextDoSet(t *testing.T) {
	tasks := []Task{
		newSetTask(1),
		newSetTask(2),
		newSetTask(3),
		newSetTask(4),
	}

	c := New(tasks)
	c.Do()

	contextAssert(t, c, 4)

	if c.err != nil {
		t.Fatal("Tasks no finished.")
	}
}

func TestContextAbort(t *testing.T) {
	tasks := []Task{
		newSetTask(1),
		newSetTask(2),
		Multi(
			&abortTask{},
			&abortTask{},
			&abortTask{},
			&abortTask{},
		),
		newSetTask(3),
		newSetTask(4),
	}

	c := New(tasks)
	if err := c.Do(); err == nil {
		t.Fatal("Tasks abort no finished")
	}

	contextAssert(t, c, 2)

	size := len(c.store)

	if size != 2 {
		t.Fatalf("Val (size) is %d; want %d", size, 2)
	}
}

func TestContextSub(t *testing.T) {
	task := Multi(
		newSetTask(1),
		newSetTask(2),
		newSetTask(3),
		newSetTask(4),
	)

	c := New([]Task{})
	if err := c.Sub(task); err != nil {
		t.Fatal("Sub tasks failed.")
	}
	contextAssert(t, c, 4)
}

func TestContextSubCatcher(t *testing.T) {
	task := Multi(
		&panicTask{},
		&panicTask{},
	)

	c := New([]Task{})

	if err := c.Sub(task); err == nil {
		t.Fatal("Err was nil want error.")
	}
}

func TestContextDoCatcher(t *testing.T) {
	tasks := []Task{
		Multi(
			&panicTask{},

			&panicTask{},
		),
	}

	c := New(tasks)

	if err := c.Do(); err == nil {
		t.Fatal("Err was nil want error.")
	}
}
