package output

import (
	"fmt"
	"prediction/pkg/common"
	"strconv"
)

func PrintPredictions(
	aggregatedData map[string]common.AverageData,
	predictionModel func(value common.AverageData,
	) float64) {
	for key, value := range aggregatedData {
		prediction := predictionModel(value)
		fmt.Println(key + ": " + strconv.FormatFloat(prediction, 'f', 2, 64))
	}
}
