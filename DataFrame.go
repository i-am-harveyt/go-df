package go_df

import (
	"fmt"
	"strings"
)

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

func (df *DataFrame) PrintHead(rows uint) {
	fmt.Printf("| %v |\n", strings.Join(df.Cols, "|"))
	for i := range rows {
		row := []string{}
		for _, name := range df.Cols {
			row = append(
				row,
				fmt.Sprint(df.Get(name, i)),
			)
		}
		fmt.Printf("| %v |\n", strings.Join(row, "|"))
	}
}
