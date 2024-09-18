package go_df

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func FromMap(map_ *map[string][]interface{}) *DataFrame {
	df := DataFrame{
		Cols: []string{},
		Data: map[string]*Col{},
	}
	for k, v := range *map_ {
		df.Cols = append(df.Cols, k)
		df.Data[k] = &Col{
			Name: k,
			Data: v,
		}
		switch v[0].(type) {
		case bool:
			df.Data[k].Type = Bool
		case float32, float64:
			df.Data[k].Type = Float64
		case int:
			df.Data[k].Type = Int
		case string:
			df.Data[k].Type = String
		}
	}
	return &df
}

func FromCSV(filepath string) (*DataFrame, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	df := Init()

	// init & header
	scanner.Scan()
	headers := strings.Split(scanner.Text(), ",")
	df.OverwriteCols(headers)
	for _, col := range df.Cols {
		df.Data[col] = &Col{
			Name: col,
		}
	}

	// data
	for scanner.Scan() {
		cells := strings.Split(scanner.Text(), ",")
		if len(cells) == 0 { // skip the empty row
			continue
		}
		for i, cell := range cells {
			col := df.Data[df.Cols[i]]
			if v, err := strconv.ParseInt(cell, 10, 32); err == nil {
				col.Type = Int
				col.Data = append(col.Data, v)
			} else if v, err := strconv.ParseFloat(cell, 64); err == nil {
				col.Type = Float64
				col.Data = append(col.Data, v)
			} else if v, err := strconv.ParseBool(cell); err == nil {
				col.Type = Bool
				col.Data = append(col.Data, v)
			} else {
				col.Type = String
				col.Data = append(col.Data, cell)
			}
		}
	}

	return df, nil
}
