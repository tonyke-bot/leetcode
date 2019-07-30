package problem0496

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_nextGreaterElement(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		args args
		want []int
	}{
		{args{[]int{4, 1, 2}, []int{1, 3, 4, 2}}, []int{-1, 3, -1}},
		{args{[]int{2, 4}, []int{1, 2, 3, 4}}, []int{3, -1}},
		{args{[]int{1, 1}, []int{1, 2, 3, 4}}, []int{2, 2}},
		{args{[]int{1}, []int{1, 2, 3, 4}}, []int{2}},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprintf("test %d", index), func(t *testing.T) {
			if got := nextGreaterElement(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
