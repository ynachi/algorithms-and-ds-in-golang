package stack

import (
	"testing"
)

func TestBalancedParentheses(t *testing.T) {
	in0 := ""
	in1 := "(()()()())"
	in2 := "(((())))"
	in3 := "(()((())()))"
	in4 := "((((((())"
	in5 := "()))"
	in6 := "(()()(()"
	tests := []struct {
		name string
		give string
		want bool
	}{
		{
			name: "balanced_0",
			give: in0,
			want: true,
		},
		{
			name: "balanced_1",
			give: in1,
			want: true,
		},
		{
			name: "balanced_2",
			give: in2,
			want: true,
		},
		{
			name: "balanced_3",
			give: in3,
			want: true,
		},
		{
			name: "imbalanced_4",
			give: in4,
			want: false,
		},
		{
			name: "imbalanced_5",
			give: in5,
			want: false,
		},
		{
			name: "imbalanced_6",
			give: in5,
			want: false,
		},
		{
			name: "imbalanced_7",
			give: in6,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BalancedParentheses(tt.give)
			if tt.want != got {
				const msg = `BalancedParentheses : wanted %v but got %v.`
				t.Fatalf(msg, tt.want, got)
			}
		})
	}
}

func TestBalancedSymbols(t *testing.T) {
	in0 := ""
	in1 := "{{([][])}()}"
	in2 := "[[{{(())}}]]"
	in3 := "[][][](){}"
	in4 := "([)]"
	in5 := "((()]))"
	in6 := "[{()]"
	tests := []struct {
		name string
		give string
		want bool
	}{
		{
			name: "balanced_0",
			give: in0,
			want: true,
		},
		{
			name: "balanced_1",
			give: in1,
			want: true,
		},
		{
			name: "balanced_2",
			give: in2,
			want: true,
		},
		{
			name: "balanced_3",
			give: in3,
			want: true,
		},
		{
			name: "imbalanced_4",
			give: in4,
			want: false,
		},
		{
			name: "imbalanced_5",
			give: in5,
			want: false,
		},
		{
			name: "imbalanced_6",
			give: in5,
			want: false,
		},
		{
			name: "imbalanced_7",
			give: in6,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BalancedSymbols(tt.give)
			if tt.want != got {
				const msg = `BalancedSymbols : wanted %v but got %v.`
				t.Fatalf(msg, tt.want, got)
			}
		})
	}
}
