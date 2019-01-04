package leetcode

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	var r []string
	r = generateParenthesis(3)
	t1 := []string{
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()",
	}
	if !reflect.DeepEqual(r, t1) {
		t.Fatalf("\n%+v \n%+v", r, t1)
	}
}
