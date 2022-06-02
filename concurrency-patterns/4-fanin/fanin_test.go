package __fanin

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fanIn(t *testing.T) {
	type args struct {
		cs []<-chan string
	}
	msgs := []string{"Joe", "Ahn"}
	j := boring(msgs[0])
	a := boring(msgs[1])
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "fanIn",
			args: args{
				cs: []<-chan string{
					j,
					a,
				},
			},
			want: msgs,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := fanIn(tt.args.cs...)
			for i := 0; i < 5; i++ {
				fmt.Println(<-c)
				assert.Contains(t, msgs, <-c)
			}
			fmt.Println("You're both boring. I'm leaving")
		})
	}
}
