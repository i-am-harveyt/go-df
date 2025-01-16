package go_df

type Series struct {
	Name string
	Type DataType
	Data []Element
}

func (c *Series) Len() int {
	return len(c.Data)
}

func (c *Series) Get(i uint) interface{} {
	return c.Data[i]
}
