package rhombus

import (
	"testing"
)

func TestChainDo(t *testing.T) {
	strategies := []Strategy{
		Chain(
			newSetStrategy("1", 1),
			newSetStrategy("2", 2),
			newSetStrategy("3", 3),
			newSetStrategy("4", 4),
		),
	}

	c := New(strategies)
	c.Do()
	go c.Close()

	contextAssert(t, c, 4)
}
