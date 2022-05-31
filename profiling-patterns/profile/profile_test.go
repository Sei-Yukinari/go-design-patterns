package profile

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	type args struct {
		invocation time.Time
		name       string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "profile",
			args: args{
				invocation: time.Now(),
				name:       "IntFactorial",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer Duration(tt.args.invocation, tt.args.name)
			time.Sleep(1 * time.Second)
		})
	}
}
