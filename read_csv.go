package main

import (
	"encoding/csv"
	"os"
)

func NewCSVReader() (*csv.Reader, error) {
	csvFile, err := os.Open("./data.csv")
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(csvFile)
	return reader, nil
}
