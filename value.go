package rhombus

type Value struct {
	Key string
	Job Task
}

func (v *Value) Do(c *Context) {
	v.Job.Do(c)
	c.Set(v.Key, v.Job.Value())
}

func (v *Value) Value() interface{} {
	return nil
}
