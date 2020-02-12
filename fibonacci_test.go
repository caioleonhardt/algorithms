package algorithms

import "testing"

var (
	tests = []struct {
		name string
		n    int
		want int
	}{
		{"zero", 0, 0},
		{"one", 1, 1},
		{"two", 2, 1},
		{"nine", 9, 34},
		{"twelve", 12, 144},
		{"forteen", 14, 377},
	}
)

func TestMemoization(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := memoization(tt.n); got != tt.want {
				t.Errorf("memoization() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMemoization(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		memoization(10)
	}
}

func TestRecursion(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recursion(tt.n); got != tt.want {
				t.Errorf("recursion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRecursion(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		recursion(10)
	}
}
