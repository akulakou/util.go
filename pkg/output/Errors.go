package output

import (
	"fmt"
	"prediction/pkg/arguments"
)

func PrintErrors() {
	if errorMsg := recover(); errorMsg != nil {
		fmt.Println(errorMsg)
		arguments.ResetParams()
		arguments.PrintUsage()

	}
}
