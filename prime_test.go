package algorithms

import "testing"

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name string
		n    int64
		want bool
	}{
		{"non negative small", 1, false},
		{"negative", -1, false},
		{"first prime", 2, true},
		{"even non prime", 10, false},
		{"odd non prime", 45, false},
		{"small prime", 17, true},
		{"medium prime", 1217, true},
		{"big prime", 100003679, true},
		{"big big prime", 19000069423, true},
		{"big big big prime", 8916047033476472063, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.n); got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
