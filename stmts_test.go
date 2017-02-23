package rhombus

import (
	"testing"
)

type yesCond struct{}
type noCond struct{}

func (y *yesCond) Value(c *Context) bool { return true }
func (n *noCond) Value(c *Context) bool  { return false }

func TestStmtsIf(t *testing.T) {
	tasks := []Task{
		If(&yesCond{},
			Multi(
				newSetTask(1),
				newSetTask(2),
				newSetTask(3),
				newSetTask(4),
			),
		),
	}

	c := New(tasks)
	c.Do()

	contextAssert(t, c, 4)
}

func TestStmtsIfNot(t *testing.T) {
	tasks := []Task{
		Multi(
			IfNot(&noCond{},
				Multi(
					newSetTask(1),
					newSetTask(2),
					newSetTask(3),
					newSetTask(4),
				),
			),
		),
	}

	c := New(tasks)
	c.Do()

	contextAssert(t, c, 4)
}

func TestStmts(t *testing.T) {
	tasks := []Task{
		Multi(
			If(&yesCond{},
				Multi(
					newSetTask(1),
					newSetTask(2),
					newSetTask(3),
					newSetTask(4),
					Chain(
						If(&noCond{},
							newSetTask(5),
						),
					),
				),
			),
		),
	}

	c := New(tasks)
	c.Do()

	contextAssert(t, c, 4)

	size := len(c.store)

	if size != 4 {
		t.Fatalf("Val (size) is %d; want %d", size, 4)
	}
}
