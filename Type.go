package go_df

type (
	Element  interface{}
	RowDict  map[string]interface{}
	RowSlice []interface{}
)

// For DataType
type DataType int

const (
	Bool = iota
	Int
	Float64
	String
)
