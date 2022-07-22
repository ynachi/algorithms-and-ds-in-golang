/*
Sliding windows pattern
*/
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

// Given a string, find the length of the longest substring in it
// with no more than K distinct characters.
func LongestSubstringKDistinct(str string, k int) int {
	var (
		windowStart int
		windowSize  int
	)
	charTrack := map[rune]int{}
	strToRune := []rune(str)
	windowMax := -1
	for windowEnd, v := range strToRune {
		charTrack[v]++
		for len(charTrack) > k {
			windowSize = windowEnd - windowStart
			if charTrack[strToRune[windowStart]] > 1 {
				charTrack[v]--
			} else {
				delete(charTrack, strToRune[windowStart])
			}
			windowStart++
			if windowSize > windowMax {
				windowMax = windowSize
			}
		}

	}
	// if we did not find relevant substring, then
	// result is the length of the original string
	if windowMax == -1 {
		windowMax = len(strToRune)
	}
	return windowMax
}

// Given a string s, find the length of the longest substring without repeating characters.
func lengthOfLongestSubstring(s string) int {
	var (
		windowStart int
		windowSize  int
		windowEnd   int
	)
	charTrack := map[rune]bool{} // simulates set in golang
	strToRune := []rune(s)       // will be esier to loop over vs bytes
	windowMax := -1
	for windowEnd < len(strToRune) {
		//if found a dupplicate rune, it means we are at
		// the end of a substring so we don't increment
		// the end of the window. We process the substring
		// and shrink the windows from the beginning
		if _, ok := charTrack[strToRune[windowEnd]]; ok {
			windowSize = windowEnd - windowStart
			if windowSize > windowMax {
				windowMax = windowSize
			}
			delete(charTrack, strToRune[windowStart])
			windowStart++
		} else {
			charTrack[strToRune[windowEnd]] = true
			windowEnd++
		}

	}
	// Process the last windowSize
	windowSize = windowEnd - windowStart
	if windowSize > windowMax {
		windowMax = windowSize
	}
	return windowMax
}

// Given a string and a pattern, find out if the string contains any permutation of the pattern.
// LC 567
func FindPermutation(str string, pattern string) bool {
	patternToRune := []rune(pattern)
	var windowStart int
	strToRune := []rune(str)
	patternSize := len(patternToRune)
	for windowEnd := patternSize; windowEnd <= len(strToRune); windowEnd++ {
		if isPermutation(patternToRune, strToRune[windowStart:windowEnd]) {
			return true
		}
		windowStart++
	}
	return false
}

func isPermutation(str []rune, pattern []rune) bool {
	strMap := map[rune]int{}
	for _, v := range str {
		strMap[v]++
	}
	for _, v := range pattern {
		if strMap[v] > 1 {
			strMap[v]--
		} else {
			delete(strMap, v)
		}
	}
	return len(strMap) == 0
}
