package rhombus

import (
	"testing"
)

func newLazyIntsTask(name string) Task {
	return &lazyTask{NewInts(name)}
}

func TestLazyInts(t *testing.T) {
	tasks := []Task{
		Value("1", newLazyIntsTask("a")),
		Value("2", newLazyIntsTask("b")),
		Value("3", newLazyIntsTask("c")),
		Value("4", newLazyIntsTask("d")),
	}

	items := map[string][]int{
		"a": []int{1, 10, 100, 1000},
		"b": []int{2, 20, 200, 2000},
		"c": []int{3, 30, 300, 3000},
		"d": []int{4, 40, 400, 4000},
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
		except := c.Get(key2).(*LazyInts)

		size := len(value)

		for i := 0; i < size; i++ {
			if value[i] != except.Val[i] {
				t.Fatalf("Val (%g) %+v ; want (%+v)", value, value, except)
			}
		}
	}
}
