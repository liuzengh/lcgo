package lcgo

import "fmt"

func pivotIndex(nums []int) int {
	var rightSum int
	for _, num := range nums {
		rightSum += num
	}

	var leftSum int
	for i := 0; i < len(nums); i += 1 {
		if i != 0 {
			leftSum += nums[i-1]
		}
		rightSum -= nums[i]
		fmt.Println(leftSum)
		if leftSum == rightSum {
			return i
		}

	}
	return -1
}
