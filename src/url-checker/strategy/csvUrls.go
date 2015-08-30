package strategy

import (
	"encoding/csv"
	"os"
)

type CsvUrls struct {
	Path string
}

func (s *CsvUrls) Get() ([]string, error) {
	csvfile, err := os.Open(s.Path)
	if err != nil {
		return []string{}, err
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = 1

	CSVdata, err := reader.ReadAll()
	if err != nil {
		return []string{}, err
	}

	//Unroll to remove the extra array for ease of use later
	var urls []string
	for _, val := range CSVdata {
		urls = append(urls, val[0])
	}

	return urls, nil
}
