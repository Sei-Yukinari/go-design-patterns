package __chan

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s", msg)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
