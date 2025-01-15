package go_df

type DataType int

const (
	Bool = iota
	Int
	Float64
	String
)

type Col struct {
	Name string
	Type DataType
	Data []interface{}
}

func (c *Col) Len() int {
	return len(c.Data)
}

func (c *Col) Get(i uint) interface{} {
	return c.Data[i]
}
