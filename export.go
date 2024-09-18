package go_df

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func (df *DataFrame) ToCSV(filename string) error {
	out := []string{}
	shape := df.Shape()
	if shape[0] == 0 {
		return errors.New("No data is exported")
	}

	// header
	headers := []string{}
	for _, headerCol := range df.Cols {
		headers = append(headers, headerCol)
	}
	out = append(out, strings.Join(headers, ","))

	// data
	for r := range shape[0] {
		row := []string{}
		for c := range shape[1] {
			col := df.Data[df.Cols[c]]
			switch col.Type {
			case String:
				row = append(row, fmt.Sprintf("\"%v\"", col.Data[r]))
			default:
				row = append(row, fmt.Sprintf("%v", col.Data[r]))
			}
		}
		out = append(out, strings.Join(row, ","))
	}

	layers := strings.Split(filename, "/")
	if len(layers) > 1 {
		os.MkdirAll(strings.Join(layers[:len(layers)-1], "/"), 0444)
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = f.WriteString(strings.Join(out, "\n"))
	return err
}
