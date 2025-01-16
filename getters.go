package go_df

func (df *DataFrame) Get(col string, idx uint) interface{} {
	return df.Data[col].Get(idx)
}

func (df *DataFrame) GetRowSlice(idx uint) *RowSlice {
	ret := make(RowSlice, df.Shape()[1])

	for i, col := range df.Cols {
		ret[i] = df.Get(col, uint(i))
	}

	return &ret
}

func (df *DataFrame) GetRowDict(idx uint) *RowDict {
	ret := make(RowDict)

	for i, col := range df.Cols {
		ret[col] = df.Get(col, uint(i))
	}

	return &ret
}

func (df *DataFrame) GetBool(col string, idx uint) bool {
	data := df.Data[col].Get(idx)
	return data.(bool)
}

func (df *DataFrame) GetInt(col string, idx uint) int {
	data := df.Data[col].Get(idx)
	return data.(int)
}

func (df *DataFrame) GetFloat64(col string, idx uint) float64 {
	data := df.Data[col].Get(idx)
	return data.(float64)
}

func (df *DataFrame) GetString(col string, idx uint) string {
	data := df.Data[col].Get(idx)
	return data.(string)
}
