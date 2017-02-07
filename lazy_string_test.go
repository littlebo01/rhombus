package rhombus

import (
	"testing"
)

func newLazyStringTask(name string) Task {
	return &lazyTask{NewString(name)}
}

func TestLazyString(t *testing.T) {
	tasks := []Task{
		Value("1", newLazyStringTask("a")),
		Value("2", newLazyStringTask("b")),
		Value("3", newLazyStringTask("c")),
		Value("4", newLazyStringTask("d")),
	}

	items := map[string]string{
		"a": "1010101001010101",
		"b": "202002002020200222",
		"c": "3030300303030303030",
		"d": "400404040400404040400",
	}

	pairs := map[string]string {
		"a": "1",
		"b": "2",
		"c": "3",
		"d": "4",
	}

	c := New(tasks)
	for key, value := range(items) {
		c.Set(key, value)
	}

	c.Do()

	for key1, key2 := range(pairs)  {
		value := items[key1]
		except := c.Get(key2).(*LazyString)

		if value != except.Val {
			t.Fatalf("Val (%g) %+v ; want (%+v)", value, value, except)
		}
	}
}
