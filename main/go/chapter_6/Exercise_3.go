package chapter_6

func findGreatest(x...int) int {
	greatest := 0
	for _, value := range x {
		if value > greatest {
			greatest = value
		}
	}
	return greatest
}
