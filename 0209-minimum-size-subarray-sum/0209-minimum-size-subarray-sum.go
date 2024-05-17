func minSubArrayLen(target int, nums []int) int {
	output:=0
	sum:=0
	counter:=0
	for i:=0; i<len(nums);i++{
    sum+= nums[i]
		for sum>=target{
			if output==0 || (i-counter) <output{
				output=i-counter+1
			}
			sum-=nums[counter]
			counter++
		}
	}
    return output
}