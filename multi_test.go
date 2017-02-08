package rhombus

import (
	"testing"
)

func TestMultiDo(t *testing.T) {
	tasks := []Task{
		Multi(
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
