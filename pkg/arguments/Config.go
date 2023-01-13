package arguments

var ArgumentsConfig [3]CmdArgument = [3]CmdArgument{
	{
		name:        MODEL,
		description: "Prediction model: Pessimistic|Optimistic",
		regexp:      `^(Pessimistic|Optimistic)$`,
	},
	{
		name:        SOURCE,
		description: "Data source: *.json|*.csv ",
		regexp:      `^(.*)(json|csv)$`,
	},
	{
		name:        AGGREGATE,
		description: "Aggregation level: country|campaign",
		regexp:      `^(country|campaign)$`,
	},
}
