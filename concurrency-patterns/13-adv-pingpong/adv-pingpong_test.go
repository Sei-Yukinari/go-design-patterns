package _3_adv_pingpong

import (
	"fmt"
	"testing"
	"time"
)

func Test_player(t *testing.T) {
	type args struct {
		name  []string
		table chan *Ball
	}
	table := make(chan *Ball)
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ping pong",
			args: args{
				name:  []string{"ping", "pong"},
				table: table,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go player(tt.args.name[0], table)
			go player(tt.args.name[1], table)
			table <- new(Ball)
			time.Sleep(10 * time.Millisecond)
			b := <-table
			//assert.Equal(t, 10/1, b.hits)
			fmt.Println(b.hits)
		})
	}
}
