package __daisy_chan

func f(left chan<- int, right <-chan int) {
	left <- 1 + <-right
}
