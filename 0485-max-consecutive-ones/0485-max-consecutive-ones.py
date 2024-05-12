class Solution(object):
    def findMaxConsecutiveOnes(self, nums):
            counter, answer = 0, 0
            for i in range(len(nums)):
                if nums[i] == 1:
                    counter += 1
                else:
                    answer = max(answer,counter)
                    counter = 0

            return max(answer, counter)