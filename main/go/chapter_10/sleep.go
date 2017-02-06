package chapter_10

import "time"

func sleep (duration time.Duration) {
	<-time.After(duration)
}

func createBufferedChannel() chan string {
	bufferedChannel := make(chan string, 20)
	return bufferedChannel
}