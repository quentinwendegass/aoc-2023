package utils

import "strings"

func ValuesFromLabelString(str string, labelSeperator string, valueSeperator string) (label string, values []string) {
	labelAndValues := strings.Split(str, labelSeperator)
	label = labelAndValues[0]

	values = ValuesFromString(labelAndValues[1], valueSeperator)
	return
}

func ValuesFromLabelStringNoLabel(str string, labelSeperator string, valueSeperator string) []string {
	_, values := ValuesFromLabelString(str, labelSeperator, valueSeperator)
	return values
}

func ValuesFromString(str string, valueSeperator string) []string {
	values := make([]string, 0)
	rawValues := strings.Split(str, valueSeperator)

	for _, value := range rawValues {
		if value == "" {
			continue
		}

		values = append(values, value)
	}

	return values
}
