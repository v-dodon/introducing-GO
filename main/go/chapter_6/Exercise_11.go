package chapter_6

func swap(x *int, y *int) {
	newX := *x
	*x = *y
	*y = newX
}
