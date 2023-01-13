package readers

import (
	"encoding/json"
	"os"
)

func JsonStrategy(
	file *os.File,
	aggregated map[string]AverageData,
	reducer func(
		aggregated map[string]AverageData,
		data ReadResult,
	),
) *[]ReadResult {
	decoder := json.NewDecoder(file)
	decoder.Token()
	data := ReadResult{}
	result := []ReadResult{}

	for decoder.More() {
		decoder.Decode(&data)

		result = append(result, data)

		reducer(aggregated, data)

	}
	return &result
}
