package chapter_5

func lengthOfSlice() int {
	slice := make([]int, 3, 9)
	return len(slice)
}