package __chan

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_boring(t *testing.T) {
	type args struct {
		msg string
		c   chan string
	}
	c := make(chan string)
	tests := []struct {
		name string
		args args
	}{
		{
			name: "chan",
			args: args{
				msg: "boring!",
				c:   c,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go boring("boring!", c)

			for i := 0; i < 5; i++ {
				fmt.Printf("You say: %q\n", <-c)
				assert.Equal(t, tt.args.msg, <-c)
			}
			fmt.Println("You're boring. I'm leaving")
		})
	}
}
