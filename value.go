package rhombus

type (
	Value struct {
		Key string
		Job Task
	}

	ValueWith struct {
		Key string
		With func(c *Context) interface{}
		value interface{}
	}
)

func (v *Value) Do(c *Context) {
	v.Job.Do(c)
	c.Set(v.Key, v.Job.Value())
}

func (v *Value) Value() interface{} {
	return nil
}

func (v *ValueWith) Do(c *Context) {
	v.value = v.With(c)
}

func (v *ValueWith) Value() interface{} {
	return v.value
}
