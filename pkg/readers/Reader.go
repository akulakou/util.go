package readers

import (
	"os"
	"path/filepath"
	"prediction/pkg/aggregators"
)

func getReaderStrategy(fileName string) func(
	file *os.File,
	aggregated map[string]AverageData,
	reducer func(
		aggregated map[string]AverageData,
		data ReadResult,
	),
) *[]ReadResult {
	fileExtension := filepath.Ext(fileName)

	switch fileExtension {
	case ".json":
		return JsonStrategy
	case ".csv":
		return CsvStrategy
	}

	panic([]string{"Data file type unknown"})
}

func Reader(fileName string, aggregate string) map[string]AverageData {
	readerStrategy := getReaderStrategy(fileName)

	file, err := os.Open(fileName)

	if err != nil {
		panic([]string{"Wrong file name, or file " + fileName + " don't exists "})
	}

	defer file.Close()

	Aggregator := aggregators.GetAggregator(aggregate)

	var result map[string]AverageData = make(map[string]AverageData)
	readerStrategy(
		file,
		result,
		Aggregator,
	)

	return result
}
