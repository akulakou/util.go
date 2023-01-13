package predictionmodels

import "prediction/pkg/common"

func Pessimistic(
	aggregatedData common.AverageData,
) float64 {
	if aggregatedData.Ltv1 > 0 {
		return aggregatedData.Ltv1
	}
	return 0.01
}
