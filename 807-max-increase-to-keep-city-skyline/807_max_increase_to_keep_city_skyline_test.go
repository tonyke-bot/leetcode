package leetcode

import "testing"

func Test_maxIncreaseKeepingSkyline(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[][]int{
			[]int{3, 0, 8, 4},
			[]int{2, 4, 5, 7},
			[]int{9, 2, 6, 3},
			[]int{0, 3, 1, 0},
		}}, 35},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxIncreaseKeepingSkyline(tt.args.grid); got != tt.want {
				t.Errorf("maxIncreaseKeepingSkyline() = %v, want %v", got, tt.want)
			}
		})
	}
}
