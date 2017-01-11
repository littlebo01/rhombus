package rhombus

type Context struct {
	store map[string]interface{}
	strategies []Strategy
}

func New(strategies []Strategy) *Context {
	return &Context{
		store: make(map[string]interface{}),
		strategies: strategies,
	}
}

func (c *Context) Do() {
	for _, strategy := range c.strategies {
		strategy.Do(c)
	}
}

func (c *Context) Get(key string) interface{} {
	return c.store[key]
}

func (c *Context) Set(key string, value interface{}) {
	c.store[key] = value
}
