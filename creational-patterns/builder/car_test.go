package car

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCar(t *testing.T) {
	basicColor := RedColor
	assembly := New().Paint(basicColor)
	t.Run("paint", func(t *testing.T) {
		assert.Equal(t, basicColor, assembly.color)
	})
	tests := []struct {
		name string
		want *car
	}{
		{
			name: "family car",
			want: &car{
				color:    basicColor,
				wheel:    SteelWheels,
				topSpeed: 100 * MPH,
			},
		},
		{
			name: "sports car",
			want: &car{
				color:    basicColor,
				wheel:    SportsWheels,
				topSpeed: 180 * KPH,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			car := assembly.Wheels(tt.want.wheel).
				TopSpeed(tt.want.topSpeed).
				Build()
			assert.Equal(t, tt.want, car)
		})
	}
}
