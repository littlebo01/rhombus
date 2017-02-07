package rhombus

type (
	Lazy interface {
		Resolve(c *Context)
	}

	LazyString struct {
		Key string
		Val string
	}

	LazyStrings struct {
		Key string
		Val []string
	}

	LazyInt struct {
		Key string
		Val int
	}

	LazyInts struct {
		Key string
		Val []int
	}
)

func (l *LazyString) Resolve(c *Context) {
	l.Val = c.Get(l.Key).(string)
}

func (l *LazyStrings) Resolve(c *Context) {
	l.Val = c.Get(l.Key).([]string)
}

func (l *LazyInt) Resolve(c *Context) {
	l.Val = c.Get(l.Key).(int)
}

func (l *LazyInts) Resolve(c *Context) {
	l.Val = c.Get(l.Key).([]int)
}
