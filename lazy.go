package rhombus

type Lazy string

func (l Lazy) Resolve(c *Context, dest interface{}) {
	src := l.Get(c)

	switch d := dest.(type) {
	case *int:
		*d = src.(int)
	case *string:
		*d = src.(string)
	case *[]int:
		*d = src.([]int)
	case *[]string:
		*d = src.([]string)
	case *[]byte:
		*d = src.([]byte)
	case *byte:
		*d = src.(byte)
	case *map[string]interface{}:
		*d = src.(map[string]interface{})
	default:
		panic("Unsupport type, use Get do your self.")
	}
}

func (l Lazy) Get(c *Context) interface{} {
	return c.Get(string(l))
}
