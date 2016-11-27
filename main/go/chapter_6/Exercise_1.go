package chapter_6

func sum(x[] int) (int) {
	sum := 0
	for _, value := range x {
		sum += value
	}
	return sum
}