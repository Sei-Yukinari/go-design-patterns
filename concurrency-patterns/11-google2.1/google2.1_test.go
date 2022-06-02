package _0_google2_1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGoogle(t *testing.T) {
	type args struct {
		query string
	}
	q := "golang"
	tests := []struct {
		name string
		args args
		want []Result
	}{
		{
			name: "google2.1",
			args: args{
				query: q,
			},
			want: []Result{"web", "image", "video"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			results := Google("golang")
			elapsed := time.Since(start)
			fmt.Println(results)
			fmt.Println(elapsed)
			assert.Len(t, results, len(tt.want))
		})
	}
}
