package rhombus

import (
	"testing"
)

type lazyIntsTask struct {
	Ids *LazyInts
}

func (t *lazyIntsTask) Do(c *Context) {
	t.Ids.Resolve(c)
}

func (t *lazyIntsTask) Value() interface{} {
	return t.Ids.Val
}

func newLazyIntsTask(name string) Task {
	return &lazyIntsTask{NewInts(name)}
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
		except := c.Get(key2).([]int)

		size := len(value)

		for i := 0; i < size; i++ {
			if value[i] != except[i] {
				t.Fatalf("Val (%g) %+v ; want (%+v)", value, value, except)
			}
		}
	}
}
