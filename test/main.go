package main

import (
	"fmt"

	"github.com/i-am-harveyt/go-df"
)

func main() {
	filename := "test.csv"
	f, err := go_df.FromCSV(filename)
	if err != nil {
		panic(fmt.Sprintf("CSV not exist: %v\n", filename))
	}
	f.PrintHead(5)
	f.ToCSV("output.csv")
}
