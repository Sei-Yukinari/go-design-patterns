package __select_timeout

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_boring(t *testing.T) {
	type args struct {
		msg string
	}
	msg := "Hello, World!"
	c := boring(msg)
	tests := []struct {
		name string
		args args
		want <-chan string
	}{
		{
			name: "select-timeout",
			args: args{
				msg: msg,
			},
			want: c,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeout := time.After(1 * time.Second)
			for {
				select {
				case s := <-c:
					assert.Equal(t, msg, s)
				case time := <-timeout:
					assert.NotEmpty(t, time)
					return
				}
			}
		})
	}
}
