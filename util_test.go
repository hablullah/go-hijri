package hijri_test

import (
	"encoding/csv"
	"io"
	"os"
)

type TestData struct {
	Gregorian string
	Hijri     string
}

func generateTestData(csvFilePath string) ([]TestData, error) {
	// Open test file
	f, err := os.Open(csvFilePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Parse test file
	dataList := []TestData{}
	csvReader := csv.NewReader(f)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		dataList = append(dataList, TestData{
			Gregorian: record[0],
			Hijri:     record[1],
		})
	}

	return dataList, nil
}
