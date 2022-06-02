package _3_adv_pingpong

import (
	"fmt"
	"time"
)

type Ball struct{ hits int }

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(1 * time.Millisecond)
		table <- ball
	}
}
