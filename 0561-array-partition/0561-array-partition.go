func arrayPairSum(nums []int) int {
    result:=0
    sort.Ints(nums)
    for i:=0;i<len(nums);i++{
      if i==0 || i % 2 == 0{
        result +=nums[i]
      }
    }
    return result
}