package rhombus

import (
	"testing"
)

func newLazyStringsTask(name string) Task {
	return &lazyTask{NewStrings(name)}
}

func TestLazyStrings(t *testing.T) {
	tasks := []Task{
		Value("1", newLazyStringsTask("a")),
		Value("2", newLazyStringsTask("b")),
		Value("3", newLazyStringsTask("c")),
		Value("4", newLazyStringsTask("d")),
	}

	items := map[string][]string{
		"a": []string{"1", "10", "100", "1000"},
		"b": []string{"2", "20", "200", "2000"},
		"c": []string{"3", "30", "300", "3000"},
		"d": []string{"4", "40", "400", "4000"},
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
		except := c.Get(key2).(*LazyStrings)

		size := len(value)

		for i := 0; i < size; i++ {
			if value[i] != except.Val[i] {
				t.Fatalf("Val (%g) %+v ; want (%+v)", value, value, except)
			}
		}
	}
}
