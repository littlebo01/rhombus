package rhombus

import (
	"testing"
)

func TestMultiDo(t *testing.T) {
	strategies := []Strategy{
		Multi(
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
