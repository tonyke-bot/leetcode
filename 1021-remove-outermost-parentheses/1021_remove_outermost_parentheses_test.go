package problem1021

import (
	"fmt"
	"testing"
)

func Test_removeOuterParentheses(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		args args
		want string
	}{
		{args{"()"}, ""},
		{args{""}, ""},
		{args{"(()())(())"}, "()()()"},
		{args{"(()())(())(()(()))"}, "()()()()(())"},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprintf("case %d", index), func(t *testing.T) {
			if got := removeOuterParentheses(tt.args.S); got != tt.want {
				t.Errorf("removeOuterParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
