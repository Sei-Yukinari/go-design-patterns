package __fanin

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- msg
		}
	}()
	return c
}

func fanIn(cs ...<-chan string) <-chan string {
	c := make(chan string)
	for _, ci := range cs {
		go func(cv <-chan string) {
			for {
				c <- <-cv
			}
		}(ci)
	}
	return c
}
