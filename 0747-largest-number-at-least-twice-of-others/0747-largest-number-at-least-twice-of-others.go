func dominantIndex(nums []int) int {
        max := 0
	index := 0
	for i := 0; i < len(nums); i++ {
		if max < nums[i]{
			max = nums[i]
			index = i
		}
	}
	for i:=0; i < len(nums); i++ {
		if nums[i]*2 > max && nums[i] != max{
			return -1
		}
	}
	return index
}