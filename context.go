package scenes

type Context struct {
	isNext bool
	isLeft bool
	data   map[string]any
}

func NewContext() *Context {
	return &Context{
		isNext: false,
		isLeft: false,
	}
}

func (c *Context) Next() {
	c.isNext = true
}

func (c *Context) Leave() {
	c.isLeft = true
}

func (c *Context) Cancel() {
	c.isNext = false
	c.isLeft = false
}

func (c *Context) AddValue(key string, value any) {
	c.data[key] = value
}

func (c *Context) GetValue(key string) (any, bool) {
	value, ok := c.data[key]
	return value, ok
}
