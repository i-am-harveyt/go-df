package go_df

type DataFrame struct {
	Cols []string
	Data map[string]*Col
}

func Init() *DataFrame {
	return &DataFrame{
		Cols: []string{},
		Data: map[string]*Col{},
	}
}

func (df *DataFrame) Shape() [2]int {
	if len(df.Cols) == 0 {
		return [2]int{0, 0}
	}
	return [2]int{
		df.Data[df.Cols[0]].Len(),
		len(df.Cols),
	}
}

func (df *DataFrame) OverwriteCols(cols []string) {
	df.Cols = cols
	for _, col := range cols {
		if _, ok := df.Data[col]; !ok {
			df.Data[col] = &Col{
				Name: col,
				Data: []interface{}{},
			}
		}
	}
}

func (df *DataFrame) GetBool(col string, idx int) bool {
	data := df.Data[col].Get(idx)
	return data.(bool)
}

func (df *DataFrame) GetInt(col string, idx int) int {
	data := df.Data[col].Get(idx)
	return data.(int)
}

func (df *DataFrame) GetFloat64(col string, idx int) float64 {
	data := df.Data[col].Get(idx)
	return data.(float64)
}

func (df *DataFrame) GetString(col string, idx int) string {
	data := df.Data[col].Get(idx)
	return data.(string)
}
