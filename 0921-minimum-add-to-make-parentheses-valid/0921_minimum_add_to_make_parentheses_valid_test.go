package problem0921

import (
	"fmt"
	"testing"
)

func Test_minAddToMakeValid(t *testing.T) {
	tests := []struct {
		args string
		want int
	}{
		{"())", 1},
		{"(((", 3},
		{"()", 0},
		{"()))((", 4},
		{"()())", 1},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprintf("case %d", index), func(t *testing.T) {
			if got := minAddToMakeValid(tt.args); got != tt.want {
				t.Errorf("minAddToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
