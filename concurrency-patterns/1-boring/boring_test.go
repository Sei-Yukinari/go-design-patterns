package __boring

import (
	"fmt"
	"testing"
	"time"
)

func Test_boring(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "boring",
			args: args{msg: "boring!"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go boring(tt.args.msg)
			fmt.Println("I'm listening")
			time.Sleep(2 * time.Second)
			fmt.Println("You're boring. I'm leaving")
		})
	}
}
