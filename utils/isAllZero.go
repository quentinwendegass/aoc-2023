package utils

func IsAllZero(values []int) bool {
	for _, val := range values {
		if val != 0 {
			return false
		}
	}

	return true
}
