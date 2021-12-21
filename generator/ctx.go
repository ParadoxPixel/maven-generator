package generator

type Ctx struct {
	variables map[string]string
}

func NewContext() *Ctx {
	return &Ctx{
		make(map[string]string),
	}
}

func (c *Ctx) Set(key, value string) {
	c.variables[key] = value
}

func (c *Ctx) Get(key string) string {
	return c.variables[key]
}

func (c *Ctx) For(f func(key, value string)) {
	for key, val := range c.variables {
		f(key, val)
	}
}
