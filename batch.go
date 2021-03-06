package rhombus

import (
	"sync"
)

func Batch(size int, tasks ...Task) Task {
	return &batchTasks{
		size:  size,
		tasks: tasks,
	}
}

type batchTasks struct {
	size  int
	tasks []Task
}

func (t *batchTasks) Do(c *Context) {
	var wg sync.WaitGroup

	for i, task := range t.tasks {
		if c.err != nil {
			break
		}

		wg.Add(1)

		go func(task Task) {
			defer wg.Done()
			defer c.catcher()

			task.Do(c)
		}(task)

		if i%t.size == 0 {
			wg.Wait()
		}
	}

	wg.Wait()
}

func (t *batchTasks) Value() interface{} {
	return nil
}
