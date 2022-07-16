/*
Two pointers pattern
*/
package patterns

func PairWithTargetSum(nums []int, target int) []int {
	leftPtr := 0
	rightPtr := len(nums) - 1
	for leftPtr < rightPtr {
		currentSum := nums[leftPtr] + nums[rightPtr]
		switch {
		case currentSum == target:
			return []int{leftPtr, rightPtr}
		case currentSum < target:
			leftPtr++
		case currentSum > target:
			rightPtr--
		}

	}
	return []int{}
}

// Alternate solution
// Works with non sorted array too
func PairWithTargetSumV2(nums []int, target int) []int {
	seen := map[int]int{}
	for numIndex, numValue := range nums {
		if mapIndex, ok := seen[target-numValue]; ok {
			return []int{numIndex, mapIndex}
		}
		seen[numValue] = numIndex
	}
	return []int{}
}
