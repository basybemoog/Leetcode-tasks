class Solution(object):
    def findNumbers(self, nums):
        i = 0
        counter = 0
        for i in range(len(nums)):
            if len(str(nums[i])) % 2 == 0: counter +=1
        return(counter)

        