package __quit_signal

import (
	"math/rand"
	"time"
)

const doneMessage = "See you!"

func boring(msg string, done chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- msg:
				// do nothing
			case <-done:
				done <- doneMessage
				return
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
