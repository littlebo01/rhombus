package rhombus

import (
	"testing"
)

func TestBatchDo(t *testing.T) {
	s1 := &setStrategy{&setStrategyParams{
				key: "1",
				value: 1,
	}}

	s2 := &setStrategy{&setStrategyParams{
				key: "2",
				value: 2,
	}}

	s3 := &setStrategy{&setStrategyParams{
				key: "3",
				value: 3,
	}}

	s4 := &setStrategy{&setStrategyParams{
				key: "4",
				value: 4,
	}}


	strategies := []Strategy{
		Batch(3, s1, s2, s3, s4),
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
