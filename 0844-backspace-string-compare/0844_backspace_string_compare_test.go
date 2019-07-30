package problem0844

import (
	"fmt"
	"testing"
)

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		S string
		T string
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{"nzp#o#g", "b#nzp#o#g"}, true},
		{args{"bbbextm", "bbb#extm"}, false},
		{args{"ab#c", "ad#c"}, true},
		{args{"ab##", "c#d#"}, true},
		{args{"a##c", "#a#c"}, true},
		{args{"a#c", "b"}, false},
		{args{"y#fo##f", "y#f#o##f"}, true},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprintf("case %d", index), func(t *testing.T) {
			if got := backspaceCompare(tt.args.S, tt.args.T); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
