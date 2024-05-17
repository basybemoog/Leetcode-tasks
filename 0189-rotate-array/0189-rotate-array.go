func rotate(nums []int, key int)  {
    if len(nums)==0{
        return
    }
	var output =make([]int,len(nums))
	for i := 0; i < len(nums); i++ {
		newIdx:= (i+key) % len(nums)
    output[newIdx] =nums[i]
	}
	copy(nums, output)
}