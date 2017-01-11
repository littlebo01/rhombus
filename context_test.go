package rhombus

import (
	"testing"
)

type echoStrategyParams struct {
	Key string
	Content string
}
type echoStrategy struct {
	params *echoStrategyParams
}

func newEchoStrategy(params *echoStrategyParams) *echoStrategy {
	return &echoStrategy{
		params: params,
	}
}

func (s *echoStrategy) Do(c *Context) {
	c.Set(s.params.Key, s.params.Content)
}

func TestContextDo(t *testing.T) {
	strategies := []Strategy{
		newEchoStrategy(&echoStrategyParams{
			Key: "w1",
			Content: "1",
		}),
		newEchoStrategy(&echoStrategyParams{
			Key: "w2",
			Content: "2",
		}),
		newEchoStrategy(&echoStrategyParams{
			Key: "w3",
			Content: "3",
		}),
	}

	c := New(strategies)
	c.Do()

	v1 := c.Get("w1")
	v2 := c.Get("w2")
	v3 := c.Get("w3")

	if v1.(string) != "1" {
		t.Fatalf("Content(%q) = %s; want %d", v1, v1, 1)
	}

	if v2.(string) != "2" {
		t.Fatalf("Content(%q) = %s; want %d", v2, v2, 2)
	}

	if v3.(string) != "3" {
		t.Fatalf("Content(%q) = %s; want %d", v3, v3, 3)
	}
}
