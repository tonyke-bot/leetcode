package problem0739

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{[]int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}, []int{8, 1, 5, 4, 3, 2, 1, 1, 0, 0}},
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{1, 2}, []int{1, 0}},
		{[]int{1}, []int{0}},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprintf("test %d", index), func(t *testing.T) {
			if got := dailyTemperatures(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dailyTemperatures() = %v, want %v", got, tt.want)
			}
		})
	}
}
