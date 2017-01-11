package rhombus

import (
	"testing"
	"strconv"
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

func contextAssert(t *testing.T, c *Context, top int) {
	for i := 1; i <= top; i++ {
		key := strconv.Itoa(i)
		val := c.Get(key)

		if val.(int) != i {
			t.Fatalf("Val(%g) = %d; want %d", val, val, i)
		}
	}
}

func TestContextSet(t *testing.T) {
	strategies := []Strategy{}
	c := New(strategies)
	c.Set("1", 1)
	c.Set("2", 2)
	c.Set("3", 3)
	c.Set("4", 4)

	contextAssert(t, c, 4)
}

func TestContextDoSet(t *testing.T) {
	strategies := []Strategy{
		newSetStrategy("1", 1),
		newSetStrategy("2", 2),
		newSetStrategy("3", 3),
		newSetStrategy("4", 4),
	}

	c := New(strategies)
	c.Do()
	defer c.Close()

	contextAssert(t, c, 4)
}
