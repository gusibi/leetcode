#
# @lc app=leetcode id=164 lang=python3
#
# [164] Maximum Gap
#


class Solution:

    def maximumGap(self, nums: List[int]) -> int:
        if len(nums) < 2:
            return 0
        nums = sorted(nums)
        maximum = 0
        for i in range(len(nums)):
            maximum = max(maximum, nums[i]-nums[i-1])
        # print(start, end)
        return maximum
