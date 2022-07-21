package queue

import (
	"reflect"
	"testing"
)

func TestInitQueue(t *testing.T) {
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
			var stack = InitQueue[int](tt.capacity)
			if len(stack.data) != 0 || cap(stack.data) != tt.capacity {
				const msg = `Stack.InitQueue : wanted %v but got %v.`
				t.Fatalf(msg, tt.wantCap, cap(stack.data))
			}
		})
	}
}

func TestEnqueue(t *testing.T) {
	data1 := []int{0, 1, 2, 3, 4}
	data2 := []int{1000}
	tests := []struct {
		name string
		give []int
		want Queue[int]
	}{
		{
			name: "test_1",
			give: data1,
			want: Queue[int]{data: []int{0, 1, 2, 3, 4}},
		},
		{
			name: "test_2",
			give: data2,
			want: Queue[int]{data: []int{1000}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s = InitQueue[int](10)
			for _, v := range tt.give {
				s.Enqueue(v)
			}
			if !reflect.DeepEqual(s, tt.want) {
				const msg = `Queue.Enqueue : wanted %v but got %v.`
				t.Fatalf(msg, tt.give, tt.want)
			}
		})
	}
}

func TestDequeue(t *testing.T) {
	data1 := Queue[int]{data: []int{0, 1, 2, 3, 4}}
	data2 := Queue[int]{data: []int{1000}}
	tests := []struct {
		name        string
		give        Queue[int]
		wantQueue   Queue[int]
		wantExtract int
	}{
		{
			name:        "test_1",
			give:        data1,
			wantQueue:   Queue[int]{data: []int{1, 2, 3, 4}},
			wantExtract: 0,
		},
		{
			name:        "test_2",
			give:        data2,
			wantQueue:   Queue[int]{data: []int{}},
			wantExtract: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			extract, err := tt.give.Dequeue()
			// we have to check 3 things
			// 1. No exception
			// 2. extracted data is the right one
			// 3. The receiver stack is still what it should be
			if err != nil || !reflect.DeepEqual(tt.give.data, tt.wantQueue.data) || tt.wantExtract != extract {
				const msg = `Queue.Dequeue : wanted queue %v but got queue %v. Wanted data %d but got %d`
				t.Fatalf(msg, tt.give, tt.wantQueue, tt.wantExtract, extract)
			}
		})
	}
}

func TestPrint(t *testing.T) {
	data1 := Queue[int]{data: []int{0, 1, 2, 3, 4}}
	data2 := Queue[int]{data: []int{}}
	tests := []struct {
		name string
		give Queue[int]
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
				const msg = `Queue.String : wanted %v but got %v.`
				t.Fatalf(msg, tt.want, got)
			}
		})
	}
}

//@TODO also test Dequeue on empty Queue
