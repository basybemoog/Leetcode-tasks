func pivotIndex(nums []int) int {
    	for i := 0; i < len(nums); i++ {
		left := 0
		right := 0
		for j:=0; j < i; j++{
			left += nums[j]
		}
		for k:= i+1; k < len(nums);k++ {
			right+= nums[k]
		}
		if left == right {
			return i
		}
	}
    return -1
}