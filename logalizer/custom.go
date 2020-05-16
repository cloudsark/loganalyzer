package logalizer

func ExecuteCustomRegex(regex *string, filename *string) {
	var list []string
	F := ReadFile(*filename)
	matches := ParseCustom(*regex, F)
	for _, match := range matches {
		list = append(list, match)
	}
	result := TopOccurr(list)
	TablePrint("Match", "Count", result)
}
