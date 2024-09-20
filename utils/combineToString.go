package utils

import "strconv"

func CombineIntsToString(values []int) int {
	combinedValueString := ""

	for _, value := range values {
		combinedValueString += strconv.Itoa(value)
	}

	combinedValue, err := strconv.Atoi(combinedValueString)
	if err != nil {
		panic(err)
	}

	return combinedValue
}
