package rhombus

import (
	"testing"
)

func TestBatchDo(t *testing.T) {
	strategies := []Strategy{
		Batch(
			3,
			newSetStrategy(1),
			newSetStrategy(2),
			newSetStrategy(3),
			newSetStrategy(4),
		),
	}

	c := New(strategies)
	c.Do()
	defer c.Close()

	contextAssert(t, c, 4)
}
