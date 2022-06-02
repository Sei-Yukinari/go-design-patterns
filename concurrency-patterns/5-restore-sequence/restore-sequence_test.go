package __restore_sequence

import (
	"fmt"
	"testing"
)

func Test_restoreSequence(t *testing.T) {
	type args struct {
		cs []<-chan Message
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
				cs: []<-chan Message{
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
				msg1 := <-c // wait to receive message
				fmt.Println(msg1.str)
				msg2 := <-c
				fmt.Println(msg2.str)

				// each go routine have to wait
				msg1.wait <- true // main goroutine allows the boring goroutine to send next value to message channel.
				msg2.wait <- true
			}
			fmt.Println("You're both boring. I'm leaving")
		})
	}
}
