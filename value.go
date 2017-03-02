package rhombus

type valueTask struct {
	key  string
	job  Task
	with func(c *Context) interface{}
}

func Value(key string, job Task) Task {
	return &valueTask{
		key:  key,
		job:  job,
		with: nil,
	}
}

func ValueWith(key string, with func(c *Context) interface{}) Task {
	return &valueTask{
		key:  key,
		job:  nil,
		with: with,
	}
}

func (t *valueTask) Do(c *Context) {
	var val interface{}

	if t.job != nil {
		t.job.Do(c)
		val = t.job.Value()
	} else {
		val = t.with(c)

	}

	if t.key != "_" && val != nil {
		c.Set(t.key, val)
	}
}

func (t *valueTask) Value() interface{} {
	return nil
}
