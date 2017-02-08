package rhombus

import (
	"testing"
)

func TestMutliNestTasks(t *testing.T) {
	tasks := []Task{
		Multi(
			newSetTask(1),
			newSetTask(2),
			Batch(
				2,
				newSetTask(3),
				newSetTask(4),
				Multi(
					newSetTask(5),
					newSetTask(6),
				),
			),
		),
	}

	c := New(tasks)
	c.Do()

	contextAssert(t, c, 6)
}

func TestChainNestTasks(t *testing.T) {
	tasks := []Task{
		Chain(
			newSetTask(1),
			newSetTask(2),
			Multi(
				newSetTask(3),
				newSetTask(4),
				Chain(
					newSetTask(5),
					newSetTask(6),
				),
				Batch(
					2,
					newSetTask(7),
					newSetTask(8),
					newSetTask(9),
					newSetTask(10),
					newSetTask(11),
					newSetTask(12),
				),
			),
		),
	}

	c := New(tasks)
	c.Do()

	contextAssert(t, c, 12)
}

func TestNestTasks(t *testing.T) {
	tasks := []Task{
		newSetTask(1),
		newSetTask(2),
		newSetTask(3),
		Batch(
			2,
			Multi(
				newSetTask(4),
				newSetTask(5),
			),
			Chain(
				newSetTask(16),
				newSetTask(17),
			),
			Multi(
				newSetTask(6),
				newSetTask(7),
			),
			Chain(
				newSetTask(18),
				newSetTask(19),
			),
			Multi(
				newSetTask(8),
				newSetTask(9),
			),
			Chain(
				newSetTask(20),
				newSetTask(21),
			),
			Multi(
				newSetTask(10),
				newSetTask(11),
			),
			Chain(
				newSetTask(22),
				newSetTask(23),
			),
			Multi(
				newSetTask(12),
				newSetTask(13),
			),
			Chain(
				newSetTask(24),
				newSetTask(25),
			),
			Multi(
				newSetTask(14),
				newSetTask(15),
			),
			Chain(
				newSetTask(26),
				newSetTask(27),
			),
			Chain(
				Multi(
					newSetTask(28),
					newSetTask(29),
					newSetTask(30),
				),
				Chain(
					newSetTask(31),
					newSetTask(32),
					newSetTask(33),
				),
			),
		),
	}

	c := New(tasks)
	c.Do()

	contextAssert(t, c, 33)
}
