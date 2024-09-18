package main

import (
	"fmt"
	dataframe "github.com/i-am-harveyt/go_df"
)

func main() {
	filename := "test.csv"
	_, err := dataframe.FromCSV(filename)
	if err != nil {
		panic(fmt.Sprintf("CSV not exist: %v\n", filename))
	}
}
