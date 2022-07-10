package lcgo

//func twoSum(nums []int, target int) []int {
//	for i, num1 := range nums {
//		for j, num2 := range nums {
//			if i != j && num1+num2 == target {
//				return []int{i, j}
//			}
//		}
//	}
//	return nil
//}

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
