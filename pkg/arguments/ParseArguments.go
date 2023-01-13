package arguments

import (
	"flag"
	"regexp"
)

type CmdArgument struct {
	name        string
	description string
	ptr         *string
	regexp      string
}

func initArgumentItem(arg CmdArgument) CmdArgument {
	ptr := flag.String(arg.name, "", arg.description)

	var result = CmdArgument{
		name:        arg.name,
		description: arg.description,
		ptr:         ptr,
		regexp:      arg.regexp,
	}

	return result
}

func ParseArguments(config []CmdArgument) []CmdArgument {
	var result []CmdArgument
	for _, configItem := range config {
		result = append(result, initArgumentItem(configItem))
	}

	flag.Parse()

	return result
}

func getErrorMsg(name string, errorString string) string {
	return name + errorString
}

func ResetParams() {
	for _, arg := range ArgumentsConfig {
		flag.Set(arg.name, "")
	}
}

func PrintUsage() {
	flag.PrintDefaults()
}

func GetErrors(arguments *[]CmdArgument) []string {
	var result []string

	for _, arg := range *arguments {
		matched, _ := regexp.MatchString(arg.regexp, *(arg.ptr))
		if !matched {
			result = append(result, getErrorMsg(arg.name, WRONG_ARGUMENT))
		}
	}

	return result
}

func GetArgumentValue(argumentName string, arguments *[]CmdArgument) string {
	var result string

	for _, arg := range *arguments {
		if argumentName == arg.name {
			result = *arg.ptr
		}
	}

	return result
}

func MapAggratateToFieldName(aggregate string) string {
	switch aggregate {
	case "campaign":
		return "CampaignId"
	case "country":
		return "Country"
	}

	panic([]string{"Aggregate parameter is unknonw"})
}
