package rhombus

import (
	"testing"
)

type setStrategyParams struct {
	key string
	value interface{}
}

type setStrategy struct {
	params *setStrategyParams
}

func newSetStrategy(key string, value interface{}) *setStrategy{
	return &setStrategy{&setStrategyParams{
		key: key,
		value: value,
	}}
}

func (s *setStrategy) Do(c *Context) {
	c.DoSet(s.params.key, s.params.value)
}

func TestContextSet(t *testing.T) {
	strategies := []Strategy{}
	c := New(strategies)
	c.Set("1", 1)
	c.Set("2", 2)
	c.Set("3", 3)
	c.Set("4", 4)

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

func TestContextDoSet(t *testing.T) {
	strategies := []Strategy{
		&setStrategy{&setStrategyParams{
			key: "1",
			value: 1,
		}},
		&setStrategy{&setStrategyParams{
			key: "2",
			value: 2,
		}},
		&setStrategy{&setStrategyParams{
			key: "3",
			value: 3,
		}},
		&setStrategy{&setStrategyParams{
			key: "4",
			value: 4,
		}},
	}

	c := New(strategies)
	c.Do()

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
