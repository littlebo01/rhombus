package rhombus

type Strategy interface {
	Prepare(c *Context)
	Do(c *Context)
	Done(c *Context)
}
