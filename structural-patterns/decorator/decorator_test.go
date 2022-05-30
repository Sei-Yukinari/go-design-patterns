package decorator

import "testing"

func Test_vegetablePizza_getPrice(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "vegetable pizza",
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &vegetablePizza{}
			if got := v.getPrice(); got != tt.want {
				t.Errorf("getPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cheeseTopping_getPrice(t *testing.T) {
	type fields struct {
		pizza pizza
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "add cheese",
			fields: fields{pizza: &vegetablePizza{}},
			want:   25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cheeseTopping{
				pizza: tt.fields.pizza,
			}
			if got := c.getPrice(); got != tt.want {
				t.Errorf("getPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tomatoTopping_getPrice(t1 *testing.T) {
	type fields struct {
		pizza pizza
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "add tomato",
			fields: fields{pizza: &vegetablePizza{}},
			want:   22,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &tomatoTopping{
				pizza: tt.fields.pizza,
			}
			if got := t.getPrice(); got != tt.want {
				t1.Errorf("getPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
