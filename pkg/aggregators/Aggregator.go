package aggregators

func GetAggregator(aggregationField string) func(
	aggregated map[string]AverageData,
	data ReadResult,
) {
	return func(
		aggregated map[string]AverageData,
		data ReadResult,
	) {
		if _, ok := data[aggregationField]; !ok {
			panic([]string{"Unknown aggregation field name"})
		}

		fieldName := data[aggregationField].(string)

		if _, ok := aggregated[fieldName]; ok {
			aggregated[fieldName] = AverageData{
				CampaignId: data["CampaignId"].(string),
				Country:    data["Country"].(string),
				Ltv1:       aggregated[fieldName].Ltv1 + data["Ltv1"].(float64),
				Ltv2:       aggregated[fieldName].Ltv1 + data["Ltv2"].(float64),
				Ltv3:       aggregated[fieldName].Ltv1 + data["Ltv3"].(float64),
				Ltv4:       aggregated[fieldName].Ltv1 + data["Ltv4"].(float64),
				Ltv5:       aggregated[fieldName].Ltv1 + data["Ltv5"].(float64),
				Ltv6:       aggregated[fieldName].Ltv1 + data["Ltv6"].(float64),
				Ltv7:       aggregated[fieldName].Ltv1 + data["Ltv7"].(float64),
			}
		} else {
			aggregated[fieldName] = AverageData{
				CampaignId: data["CampaignId"].(string),
				Country:    data["Country"].(string),
				Ltv1:       data["Ltv1"].(float64),
				Ltv2:       data["Ltv2"].(float64),
				Ltv3:       data["Ltv3"].(float64),
				Ltv4:       data["Ltv4"].(float64),
				Ltv5:       data["Ltv5"].(float64),
				Ltv6:       data["Ltv6"].(float64),
				Ltv7:       data["Ltv7"].(float64),
			}
		}
	}
}
