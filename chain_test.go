package rhombus

import (
	"testing"
)

func TestChainDo(t *testing.T) {
	tasks := []Task{
		Chain(
			newSetTask(1),
			newSetTask(2),
			newSetTask(3),
			newSetTask(4),
		),
	}

	c := New(tasks)
	c.Do()
	defer c.Close()

	contextAssert(t, c, 4)
}
