package adapter

import (
	"encoding/csv"
	"os"
)

type csvAdapter struct{}

func GetCSVAdapter() *csvAdapter {
	return &csvAdapter{}
}

func (a *csvAdapter) Write(filename string, data [][]string) error {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return err
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	w.WriteAll(data)

	return nil
}

func (a *csvAdapter) Read(filename string) ([][]string, error) {
	var records [][]string

	file, err := os.Open(filename)
	if err != nil {
		return records, err
	}

	records, err = csv.NewReader(file).ReadAll()
	if err != nil {
		return records, err
	}

	return records, nil
}
