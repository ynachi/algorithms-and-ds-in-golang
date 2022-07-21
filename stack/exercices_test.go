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

func TestDecToBin(t *testing.T) {
	number0 := 0
	number1 := 1
	number2 := 42
	number3 := 87459211
	tests := []struct {
		name string
		give int
		want string
	}{
		{
			name: "test_0",
			give: number0,
			want: "0",
		},
		{
			name: "test_1",
			give: number1,
			want: "1",
		},
		{
			name: "test_2",
			give: number2,
			want: "101010",
		},
		{
			name: "test_3",
			give: number3,
			want: "101001101101000010110001011",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DecToBin(tt.give)
			if tt.want != got {
				const msg = `DecToBin : wanted %v but got %v.`
				t.Fatalf(msg, tt.want, got)
			}
		})
	}
}
