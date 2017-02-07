package rhombus

type Task interface {
	Do(c *Context)
	Value() interface{}
}
