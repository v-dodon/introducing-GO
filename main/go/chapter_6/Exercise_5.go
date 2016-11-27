package chapter_6

func fib(number uint) uint {
	switch number {
	case 0 : return number
	case 1 : return number
	default : return fib(number - 1) + fib(number - 2)
	}
}