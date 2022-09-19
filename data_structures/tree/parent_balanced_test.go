package tree

import (
	"fmt"
	"testing"
)

func TestParent_Balanced(t *testing.T) {
	tests := []struct {
		parens string
		want   bool
	}{
		{``, true},
		{`()`, true},
		{`((()))`, true},
		{`((()())()())()`, true},

		{`)`, false},
		{`(`, false},
		{`)()(`, false},
		{`(()`, false},
		{`((()`, false},
		{`())`, false},
		{`()))`, false},
	}

	for _, test := range tests {
		got := areParensBalanced(test.parens)
		if got != test.want {
			fmt.Printf("[FAIL] areParensBalanced[%q] = %t; want %t\n", test.parens, got, test.want)
		}
	}
}

type Stack []byte

func (s *Stack) Push(w byte) {
	*s = append(*s, w)
}
func (s *Stack) Pop() byte {
	if len(*s)-1 < 0 {
		return 0
	}

	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return last
}
func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func areParensBalanced(parens string) bool {
	res := len(parens) % 2
	if res != 0 {
		return false
	}
	if len(parens) == 0 {
		return true
	}

	s := Stack{}
	for _, ch := range parens {
		if string(ch) == "(" {
			s.Push(byte(ch))
		}
		if string(ch) == ")" {
			if v := s.Pop(); v == 0 {
				return false
			}
		}
	}

	return s.Empty()
}
