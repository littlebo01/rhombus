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
	if t.job != nil {
		t.job.Do(c)
		c.Set(t.key, t.job.Value())
	} else {
		c.Set(t.key, t.with(c))
	}
}

func (t *valueTask) Value() interface{} {
	return nil
}
