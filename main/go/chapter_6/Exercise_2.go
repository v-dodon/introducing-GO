package chapter_6

func evenOrOdd(number int) (int, bool) {
	return number / 2, number % 2 == 0
}