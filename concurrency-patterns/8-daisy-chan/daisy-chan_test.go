package __daisy_chan

import (
	"fmt"
	"testing"
)

func Test_f(t *testing.T) {
	leftmost := make(chan int)
	left := leftmost
	right := make(chan int)
	type args struct {
		left  chan int
		right chan int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "daisy-chan",
			args: args{
				left:  left,
				right: right,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 1000; i++ {
				right = make(chan int)
				go f(left, right)
				left = right
			}

			go func(c chan int) {
				c <- 1
			}(right)
			fmt.Println(<-leftmost)
		})
	}
}
