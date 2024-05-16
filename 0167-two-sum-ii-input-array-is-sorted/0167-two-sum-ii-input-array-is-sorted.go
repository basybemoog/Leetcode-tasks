func twoSum(nums []int, target int) []int {
    	start, end := 0, len(nums)-1
	for {
		if nums[start]+nums[end] > target {
			end--
		} else if target > nums[start]+nums[end] {
			start++
		} else {
			return []int{start + 1, end + 1}
		}
	}
}