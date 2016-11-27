package chapter_6

import "fmt"

func recoverAfterPanic() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
