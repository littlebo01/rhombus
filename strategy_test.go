package rhombus

import (
	"testing"
)

func TestMutliNestStrategy(t *testing.T) {
	strategies := []Strategy{
		Multi(
			newSetStrategy(1),
			newSetStrategy(2),
			Batch(
				2,
				newSetStrategy(3),
				newSetStrategy(4),
				Multi(
					newSetStrategy(5),
					newSetStrategy(6),
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
			newSetStrategy(1),
			newSetStrategy(2),
			Multi(
				newSetStrategy(3),
				newSetStrategy(4),
				Chain(
					newSetStrategy(5),
					newSetStrategy(6),
				),
				Batch(
					2,
					newSetStrategy(7),
					newSetStrategy(8),
					newSetStrategy(9),
					newSetStrategy(10),
					newSetStrategy(11),
					newSetStrategy(12),
				),
			),
		),
	}

	c := New(strategies)
	c.Do()
	defer c.Close()

	contextAssert(t, c, 12)
}

func TestNestStrategy(t *testing.T) {
	strategies := []Strategy{
		newSetStrategy(1),
		newSetStrategy(2),
		newSetStrategy(3),
		Batch(
			2,
			Multi(
				newSetStrategy(4),
				newSetStrategy(5),
			),
			Chain(
				newSetStrategy(16),
				newSetStrategy(17),
			),
			Multi(
				newSetStrategy(6),
				newSetStrategy(7),
			),
			Chain(
				newSetStrategy(18),
				newSetStrategy(19),
			),
			Multi(
				newSetStrategy(8),
				newSetStrategy(9),
			),
			Chain(
				newSetStrategy(20),
				newSetStrategy(21),
			),
			Multi(
				newSetStrategy(10),
				newSetStrategy(11),
			),
			Chain(
				newSetStrategy(22),
				newSetStrategy(23),
			),
			Multi(
				newSetStrategy(12),
				newSetStrategy(13),
			),
			Chain(
				newSetStrategy(24),
				newSetStrategy(25),
			),
			Multi(
				newSetStrategy(14),
				newSetStrategy(15),
			),
			Chain(
				newSetStrategy(26),
				newSetStrategy(27),
			),
			Chain(
				Multi(
					newSetStrategy(28),
					newSetStrategy(29),
					newSetStrategy(30),
				),
				Chain(
					newSetStrategy(31),
					newSetStrategy(32),
					newSetStrategy(33),
				),
			),
		),
	}

	c := New(strategies)
	c.Do()
	defer c.Close()

	contextAssert(t, c, 33)
}
