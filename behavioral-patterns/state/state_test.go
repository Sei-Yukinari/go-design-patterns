package state

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestState(t *testing.T) {
	type args struct {
		c *Context
	}
	tests := []struct {
		name string
		want args
	}{
		{
			name: "start",
			want: args{
				c: &Context{name: string(Start), currentState: &StartState{}},
			},
		},
		{
			name: "inprogress",
			want: args{
				c: &Context{name: string(Inprogress), currentState: &InprogressState{}},
			},
		},
		{
			name: "stop",
			want: args{
				c: &Context{name: string(End), currentState: &StopState{}},
			},
		},
	}
	for _, tt := range tests {
		var state state
		c := &Context{}
		t.Run(tt.name, func(t *testing.T) {
			state = tt.want.c.currentState
			state.handle(c)
			assert.Equal(t, tt.want.c, c)
		})
	}
}
