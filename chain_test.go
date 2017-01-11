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

	v1 := c.Get("1")
	v2 := c.Get("2")
	v3 := c.Get("3")
	v4 := c.Get("4")

	if v1.(int) != 1 {
		t.Fatalf("v1(%g) = %d; want %d", v1, 1, v1)
	}
	if v2.(int) != 2 {
		t.Fatalf("v1(%g) = %d; want %d", v2, 2, v2)
	}
	if v3.(int) != 3 {
		t.Fatalf("v1(%g) = %d; want %d", v3, 3, v3)
	}
	if v4.(int) != 4 {
		t.Fatalf("v1(%g) = %d; want %d", v4, 4, v4)
	}
}
