package problem0682

import (
	"fmt"
	"testing"
)

func Test_calPoints(t *testing.T) {
	tests := []struct {
		args []string
		want int
	}{
		{[]string{"5", "2", "C", "D", "+"}, 30},
		{[]string{"5", "-2", "4", "C", "D", "9", "+", "+"}, 27},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprintf("case %d", index), func(t *testing.T) {
			if got := calPoints(tt.args); got != tt.want {
				t.Errorf("calPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
