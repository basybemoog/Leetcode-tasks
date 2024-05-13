func twoSum(nums []int, target int) []int {
	var list []int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[i]+nums[j] == target && j != i {
				list = append(list, i)
				list = append(list, j)
				return list
			}
		}
	}
	return list
}