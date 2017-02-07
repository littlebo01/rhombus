package rhombus

type Lazy interface {
	Resolve(c *Context)
}
