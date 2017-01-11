package rhombus

import (
	"testing"
)

func TestChainDo(t *testing.T) {
	strategies := []Strategy{
		Chain(
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
