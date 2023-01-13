package predictionmodels

import "prediction/pkg/common"

type PredictionModelFn func(aggregatedRow common.AverageData) float64

var Models map[string]PredictionModelFn = map[string]PredictionModelFn{
	"Optimistic":  Optimistic,
	"Pessimistic": Pessimistic,
}

func Selector(modelName string) PredictionModelFn {
	if Models[modelName] != nil {
		return Models[modelName]
	}

	panic("Unknown prediction model")
}
