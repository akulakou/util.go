package predictionmodels

import (
	"prediction/pkg/common"
	"reflect"
)

func Optimistic(
	aggregatedRow common.AverageData,
) float64 {
	values := reflect.ValueOf(aggregatedRow)
	typesOf := values.Type()
	count := 0
	var min float64 = 0
	var max float64 = 0
	var delta float64 = 0
	var lastIndex = 0

	for i := 0; i < values.NumField(); i++ {
		if typesOf.Field(i).Name == "CampaignId" || typesOf.Field(i).Name == "Country" {
			continue
		}

		currentValue := values.Field(i).Float()

		if currentValue > 0 {
			count++
			lastIndex = i
		}

		if i == 2 {
			min = currentValue
		}
		if i > 1 {
			if min >= currentValue {
				min = currentValue
			}

			if max <= currentValue {
				max = currentValue
			}
		}
	}

	if count > 0 {
		delta = max - min
	}

	if delta < 0 {
		return max
	}

	return float64(60-lastIndex)*delta + max
}
