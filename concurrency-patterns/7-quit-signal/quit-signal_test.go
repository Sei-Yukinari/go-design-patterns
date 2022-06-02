package __quit_signal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_boring(t *testing.T) {
	done := make(chan string)
	c := boring("Joe", done)
	type args struct {
		msg  string
		quit chan string
	}
	tests := []struct {
		name string
		args args
		want chan string
	}{
		{
			name: "boring",
			args: args{msg: "quit", quit: make(chan string)},
			want: done,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 3; i >= 0; i-- {
				fmt.Println(<-c)
			}
			tt.want <- "Bye"
			assert.Equal(t, <-tt.want, doneMessage)
		})
	}
}
