package patterns

import (
	"reflect"
	"testing"
)

func TestPairWithTargetSum(t *testing.T) {
	target1 := 6
	nums1 := []int{1, 2, 3, 4, 6}
	target2 := 11
	nums2 := []int{2, 5, 9, 11}
	target3 := 6
	nums3 := []int{2, 3, 4}
	tests := []struct {
		name       string
		giveTarget int
		giveNums   []int
		want       []int
	}{
		{
			name:       "test_1",
			giveTarget: target1,
			giveNums:   nums1,
			want:       []int{1, 3},
		},
		{
			name:       "test_2",
			giveTarget: target2,
			giveNums:   nums2,
			want:       []int{0, 2},
		},
		{
			name:       "test_3",
			giveTarget: target3,
			giveNums:   nums3,
			want:       []int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PairWithTargetSum(tt.giveNums, tt.giveTarget)
			if !reflect.DeepEqual(got, tt.want) {
				const msg = `FindMinSubArray : wanted %v but got %v.`
				t.Fatalf(msg, tt.want, got)
			}
		})
	}
}
