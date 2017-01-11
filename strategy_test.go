package rhombus

import (
	"testing"
)

func TestMutliNestStrategy(t *testing.T) {
	strategies := []Strategy{
		Multi(
			newSetStrategy("1", 1),
			newSetStrategy("2", 2),
			Batch(
				2,
				newSetStrategy("3", 3),
				newSetStrategy("4", 4),
				Multi(
					newSetStrategy("5", 5),
					newSetStrategy("6", 6),
				),
			),
		),
	}

	c := New(strategies)
	c.Do()
	defer c.Close()

	contextAssert(t, c, 6)
}



func TestChainNestStrategy(t *testing.T) {
	strategies := []Strategy{
		Chain(
			newSetStrategy("1", 1),
			newSetStrategy("2", 2),
			Multi(
				newSetStrategy("3", 3),
				newSetStrategy("4", 4),
				Chain(
					newSetStrategy("5", 5),
					newSetStrategy("6", 6),
				),
				Batch(
					2,
					newSetStrategy("7", 7),
					newSetStrategy("8", 8),
					newSetStrategy("9", 9),
					newSetStrategy("10", 10),
					newSetStrategy("11", 11),
					newSetStrategy("12", 12),
				),
			),
		),
	}

	c := New(strategies)
	c.Do()
	defer c.Close()

	contextAssert(t, c, 12)
}
