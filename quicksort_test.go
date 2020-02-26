package algorithms

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{"first",
			[]int{2, 1, 4, 3, 7, 6},
			[]int{1, 2, 3, 4, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.arr)
			Quicksort(tt.arr)
			fmt.Println(tt.arr)
			if !reflect.DeepEqual(tt.arr, tt.want) {
				t.Fatalf("quisort want %v and got %v", tt.want, tt.arr)
			}
		})
	}
}

func Test_quicksort(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name  string
		arr   []int
		left  int
		right int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quicksort(tt.arr, tt.left, tt.right)
		})
	}
}
