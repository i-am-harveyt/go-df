package go_df

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
			df.Data[col] = &Series{
				Name: col,
				Data: []Element{},
			}
		}
	}
}
