package main

import (
	"fmt"

	"github.com/i-am-harveyt/go_df"
)

func main() {
	filename := "test.csv"
	_, err := go_df.FromCSV(filename)
	if err != nil {
		panic(fmt.Sprintf("CSV not exist: %v\n", filename))
	}
}
