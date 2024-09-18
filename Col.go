package go_df

type Type int

const (
	Bool = iota
	Int
	Float64
	String
)

type Col struct {
	Name string
	Type Type
	Data []interface{}
}

func (c *Col) Len() int {
	return len(c.Data)
}

func (c *Col) Get(i int) interface{} {
	return c.Data[i]
}
