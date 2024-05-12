func findDisappearedNumbers(nums []int) []int {
	var values []int
	for i := 1; i <= len(nums); i++ {
		if !check(i, nums) {
			values = append(values, i)
		}
	}
	return values
}
func check(i int, nums []int) bool {
	for j := 0; j < len(nums); j++ {
		if nums[j] == i {
			return true
		}
	}
	return false
}