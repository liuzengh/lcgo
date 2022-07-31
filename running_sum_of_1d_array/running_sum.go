package lcgo

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i += 1 {
		nums[i] += nums[i-1]
	}
	return nums
}
