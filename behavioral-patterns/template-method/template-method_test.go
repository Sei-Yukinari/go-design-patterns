package template_method

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTemplateMethod(t *testing.T) {
	tests := []struct {
		name string
		want iOtp
	}{
		{
			name: "sms",
			want: &sms{},
		},
		{
			name: "email",
			want: &email{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := otp{
				iOtp: tt.want,
			}
			err := o.genAndSendOTP(4)
			assert.NoError(t, err)
			assert.IsType(t, tt.want, o.iOtp)
		})
	}
}
