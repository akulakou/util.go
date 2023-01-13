package readers

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

func rowToMap(row []string, header []string) ReadResult {
	result := ReadResult{}

	for index, value := range row {
		if value == "UserId" {
			continue
		}
		if columnValue, err := strconv.ParseFloat(value, 32); err == nil {
			result[header[index]] = columnValue
		} else {
			result[header[index]] = value
		}
	}

	return result
}

func CsvStrategy(
	file *os.File,
	aggregated map[string]AverageData,
	reducer func(
		aggregated map[string]AverageData,
		data ReadResult,
	),
) *[]ReadResult {
	reader := csv.NewReader(file)

	result := []ReadResult{}

	header, err := reader.Read()

	if err != nil {
		panic([]string{"Error read csv header"})
	}

	for {
		row, readError := reader.Read()
		if readError == io.EOF {
			break
		}

		if err != nil {
			panic([]string{"Error read csv row"})
		}

		rowMap := rowToMap(row, header)

		result = append(result, rowMap)

		reducer(aggregated, rowMap)

	}

	return &result
}
