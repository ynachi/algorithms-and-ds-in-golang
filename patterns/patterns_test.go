package patterns

import "testing"

func TestFindMinSubArray(t *testing.T) {
	target1 := 7
	nums1 := []int{2, 3, 1, 2, 4, 3}
	target2 := 4
	nums2 := []int{1, 4, 4}
	target3 := 11
	nums3 := []int{1, 1, 1, 1, 1, 1, 1, 1}
	target4 := 15
	nums4 := []int{1, 2, 3, 4, 5}
	tests := []struct {
		name       string
		giveTarget int
		giveNums   []int
		want       int
	}{
		{
			name:       "test_1",
			giveTarget: target1,
			giveNums:   nums1,
			want:       2,
		},
		{
			name:       "test_2",
			giveTarget: target2,
			giveNums:   nums2,
			want:       1,
		},
		{
			name:       "test_3",
			giveTarget: target3,
			giveNums:   nums3,
			want:       0,
		},
		{
			name:       "test_4",
			giveTarget: target4,
			giveNums:   nums4,
			want:       5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMinSubArray(tt.giveTarget, tt.giveNums)
			if got != tt.want {
				const msg = `FindMinSubArray : wanted %v but got %v.`
				t.Fatalf(msg, tt.want, got)
			}
		})
	}
}
