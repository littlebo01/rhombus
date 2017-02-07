package rhombus

type (
	LazyInt struct {
		key string
		Val int
	}

	LazyInts struct {
		key string
		Val []int
	}

	LazyString struct {
		key string
		Val string
	}

	LazyStrings struct {
		key string
		Val []string
	}
)

func NewInt(key string) *LazyInt {
	return &LazyInt{key: key}
}

func NewInts(key string) *LazyInts {
	return &LazyInts{key: key}
}

func NewString(key string) *LazyString {
	return &LazyString{key: key}
}

func NewStrings(key string) *LazyStrings {
	return &LazyStrings{key: key}
}

func (l *LazyString) Resolve(c *Context) {
	l.Val = c.Get(l.key).(string)
}

func (l *LazyStrings) Resolve(c *Context) {
	l.Val = c.Get(l.key).([]string)
}

func (l *LazyInt) Resolve(c *Context) {
	l.Val = c.Get(l.key).(int)
}

func (l *LazyInts) Resolve(c *Context) {
	l.Val = c.Get(l.key).([]int)
}
