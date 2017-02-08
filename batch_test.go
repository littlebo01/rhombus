package rhombus

import (
	"testing"
)

func TestBatchDo(t *testing.T) {
	tasks := []Task{
		Batch(
			3,
			newSetTask(1),
			newSetTask(2),
			newSetTask(3),
			newSetTask(4),
		),
	}

	c := New(tasks)
	c.Do()

	contextAssert(t, c, 4)
}
