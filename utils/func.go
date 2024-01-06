package utils

func FilterZerosFromArray(input []int) []int {
	var result []int
	for _, value := range input {
		if value != 0 {
			result = append(result, value)
		}
	}
	return result
}
