package main

import (
	"prediction/pkg/arguments"
	"prediction/pkg/output"
	predictionmodels "prediction/pkg/predictionModels"
	"prediction/pkg/readers"
)

func main() {
	defer output.PrintErrors()

	args := arguments.ParseArguments(arguments.ArgumentsConfig[:])

	var argErrors []string = arguments.GetErrors(&args)

	if len(argErrors) > 0 {
		panic(argErrors)
	}

	aggregate := arguments.MapAggratateToFieldName(arguments.GetArgumentValue(arguments.AGGREGATE, &args))
	source := arguments.GetArgumentValue(arguments.SOURCE, &args)
	model := arguments.GetArgumentValue(arguments.MODEL, &args)

	aggregatedData := readers.Reader(source, aggregate)

	output.PrintPredictions(aggregatedData, predictionmodels.Selector(model))

}
