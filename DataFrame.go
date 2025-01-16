package go_df

import (
	"fmt"
	"strings"
)

type DataFrame struct {
	Cols []string
	Data map[string]*Series
}

func Init() *DataFrame {
	return &DataFrame{
		Cols: []string{},
		Data: map[string]*Series{},
	}
}

func (df *DataFrame) PrintHead(rows uint) {
	fmt.Printf("| %v |\n", strings.Join(df.Cols, " | "))
	for i := range rows {
		row := []string{}
		for _, name := range df.Cols {
			row = append(
				row,
				fmt.Sprint(df.Get(name, i)),
			)
		}
		fmt.Printf("| %v |\n", strings.Join(row, " | "))
	}
}

func (df *DataFrame)Copy() DataFrame {
	return *df
}
