package rhombus

func Chain(tasks ...Task) Task {
	return &batchTasks{
		size: 1,
		tasks: tasks,
	}
}
