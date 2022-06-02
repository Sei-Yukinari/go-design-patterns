package __generator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_boring(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "generator",
			args: args{msg: ""},
			want: []string{"Joe", "Ahn"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			joe := boring(tt.want[0])
			ahn := boring(tt.want[1])
			for i := 0; i < 10; i++ {
				assert.Equal(t, tt.want[0], <-joe)
				assert.Equal(t, tt.want[1], <-ahn)
			}
			fmt.Println("You're both boring. I'm leaving")
		})
	}
}
