package rhombus

import (
	"strconv"
	"testing"
)

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

func TestContextDiscardNames(t *testing.T) {
	tasks := []Task{
		newDiscardSetTask(1),
		newDiscardSetTask(2),
		newDiscardSetTask(3),
		newDiscardSetTask(4),
	}

	c := New(tasks)
	c.Do()

	if l := len(c.store); l != 0 {
		t.Fatalf("Store len was %d want %d.", l, 0)
	}
}

func TestContextDiscardValues(t *testing.T) {
	tasks := []Task{
		Value("1", &nilTask{}),
		Value("2", &nilTask{}),
		Value("3", &nilTask{}),
		Value("4", &nilTask{}),
	}

	c := New(tasks)
	c.Do()

	if l := len(c.store); l != 0 {
		t.Fatalf("Store len was %d want %d.", l, 0)
	}
}

func TestContextErrTaskWithoutAbort(t *testing.T) {
	task := Chain(
		newSetTask(1),
		newSetTask(2),
		&errTask{},
		newSetTask(3),
		newSetTask(4),
	)

	c := New([]Task{})

	if err := c.Sub(task); err != nil {
		t.Fatal("Sub tasks failed.")
	}

	contextAssert(t, c, 4)
}
