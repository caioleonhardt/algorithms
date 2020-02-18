package algorithms

import "testing"

func TestGCD(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"small", 1424, 3084, 4},
		{"medium", 4278, 8602, 46},
		{"return_1", 406, 555, 1},
		{"really big", 3_918_848, 1_653_264, 61_232},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCD(tt.a, tt.b); got != tt.want {
				t.Errorf("GCD() = %v, want %v", got, tt.want)
			}
		})
	}
}
