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
