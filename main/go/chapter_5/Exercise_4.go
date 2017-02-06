package chapter_5

func smallestNumber(slice []int) int {
	min := 0
	for _, value := range slice {
		if (value < min) {
			min = value
		}
	}
	return min
}
