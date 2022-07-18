package stack

import (
	"reflect"
	"testing"
)

func TestInitStack(t *testing.T) {
	initCap1 := 10
	initCap2 := 0
	tests := []struct {
		name     string
		capacity int
		wantCap  int
	}{
		{
			name:     "test_1",
			capacity: initCap1,
			wantCap:  initCap1,
		},
		{
			name:     "test_2",
			capacity: initCap2,
			wantCap:  initCap2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stack = InitStack[int](tt.capacity)
			if len(stack.value) != 0 || cap(stack.value) != tt.capacity {
				const msg = `FindMinSubArray : wanted %v but got %v.`
				t.Fatalf(msg, tt.wantCap, cap(stack.value))
			}
		})
	}
}

func TestPush(t *testing.T) {
	data1 := []int{0, 1, 2, 3, 4}
	data2 := []int{1000}
	tests := []struct {
		name string
		give []int
		want Stack[int]
	}{
		{
			name: "test_1",
			give: data1,
			want: Stack[int]{value: []int{0, 1, 2, 3, 4}},
		},
		{
			name: "test_2",
			give: data2,
			want: Stack[int]{value: []int{1000}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s = InitStack[int](10)
			for _, v := range tt.give {
				s.Push(v)
			}
			if !reflect.DeepEqual(s, tt.want) {
				const msg = `FindMinSubArray : wanted %v but got %v.`
				t.Fatalf(msg, tt.give, tt.want)
			}
		})
	}
}

func TestPeek(t *testing.T) {
	data1 := Stack[int]{value: []int{0, 1, 2, 3, 4}}
	data2 := Stack[int]{value: []int{1000}}
	tests := []struct {
		name        string
		give        Stack[int]
		wantStack   Stack[int]
		wantExtract int
	}{
		{
			name:        "test_1",
			give:        data1,
			wantStack:   Stack[int]{value: []int{0, 1, 2, 3, 4}},
			wantExtract: 4,
		},
		{
			name:        "test_2",
			give:        data2,
			wantStack:   Stack[int]{value: []int{1000}},
			wantExtract: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			extract, err := tt.give.Peek()
			// we have to check 3 things
			// 1. No exception
			// 2. extracted data is the right one
			// 3. The receiver stack is still what it should be
			if err != nil || !reflect.DeepEqual(tt.give, tt.wantStack) || tt.wantExtract != extract {
				const msg = `FindMinSubArray : wanted %v but got %v.`
				t.Fatalf(msg, tt.wantExtract, extract)
			}
		})
	}
}

func TestPop(t *testing.T) {
	data1 := Stack[int]{value: []int{0, 1, 2, 3, 4}}
	data2 := Stack[int]{value: []int{1000}}
	tests := []struct {
		name        string
		give        Stack[int]
		wantStack   Stack[int]
		wantExtract int
	}{
		{
			name:        "test_1",
			give:        data1,
			wantStack:   Stack[int]{value: []int{0, 1, 2, 3}},
			wantExtract: 4,
		},
		{
			name:        "test_2",
			give:        data2,
			wantStack:   Stack[int]{value: []int{}},
			wantExtract: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			extract, err := tt.give.Pop()
			// we have to check 3 things
			// 1. No exception
			// 2. extracted data is the right one
			// 3. The receiver stack is still what it should be
			if err != nil || !reflect.DeepEqual(tt.give, tt.wantStack) || tt.wantExtract != extract {
				const msg = `FindMinSubArray : wanted %v but got %v.`
				t.Fatalf(msg, tt.wantExtract, extract)
			}
		})
	}
}

func TestPrint(t *testing.T) {
	data1 := Stack[int]{value: []int{0, 1, 2, 3, 4}}
	data2 := Stack[int]{value: []int{}}
	tests := []struct {
		name string
		give Stack[int]
		want string
	}{
		{
			name: "test_1",
			give: data1,
			want: "0 1 2 3 4",
		},
		{
			name: "test_2",
			give: data2,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.give.String()
			if got != tt.want {
				const msg = `FindMinSubArray : wanted %v but got %v.`
				t.Fatalf(msg, tt.want, got)
			}
		})
	}
}

//@TODO also test Pop/Peek on empty stacks
