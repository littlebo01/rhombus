package rhombus

func Multi(tasks ...Task) Task {
	return &batchTasks{
		size:  len(tasks),
		tasks: tasks,
	}
}
