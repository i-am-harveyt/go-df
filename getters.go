package go_df

func (df *DataFrame) Get(col string, idx uint) interface{} {
	return df.Data[col].Get(idx)
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
