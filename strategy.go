package rhombus

type Strategy interface {
	Do(c *Context)
}
