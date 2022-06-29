package patterns

// findMaxAverage finds the max average of a subarrays
// in a given array (see LC 643)
func findMaxAverage(nums []int, k int) float64 {
	if k > len(nums) {
		return 0.0
	}
	windowSum := 0
	for j := 0; j <= k-1; j++ {
		windowSum += nums[j]
	}
	sumMax := windowSum
	winStart := 0
	for j := k; j < len(nums); j++ {
		windowSum += nums[j] - nums[winStart]
		winStart++
		if windowSum > sumMax {
			sumMax = windowSum
		}
	}
	return float64(sumMax) / float64(k)
}

func findMinSubArray(target int, nums []int) int {
	/*
		Use 2 indexes to slide the original array (windowStart and windowEnd).
		Move windowEnd until windowSum >= target. If found, save windowSize.
		Then, move windowStart one step and repeat the operation. Each time
		you find a new windowSize, save it if it's smaller than the previous one.
	*/
	var (
		windowStart int
		windowSize  int
		windowSum   int
	)
	windowMin := len(nums)
	foundSubarray := false // if a subarray respecting the condition is found
	for windowEnd, v := range nums {
		windowSum += v
		for windowSum >= target {
			foundSubarray = true
			windowSize = windowEnd - windowStart + 1
			windowSum -= nums[windowStart]
			windowStart++
			if windowSize < windowMin {
				windowMin = windowSize
			}
		}

	}
	if !foundSubarray {
		windowMin = 0
	}
	return windowMin
}
