package __select_timeout

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- msg
		}
	}()
	return c
}
