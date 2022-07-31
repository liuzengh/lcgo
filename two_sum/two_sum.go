package lcgo

func twoSum(nums []int, target int) []int {
	hasSeen := make(map[int]int)
	for index, num := range nums {
		if idx, ok := hasSeen[target-num]; ok {
			return []int{idx, index}
		}
		hasSeen[num] = index
	}
	return nil
}
