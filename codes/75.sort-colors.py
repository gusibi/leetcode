#
# @lc app=leetcode id=75 lang=python3
#
# [75] Sort Colors
#


class Solution:
    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        head, tail = 0, len(nums) - 1
        i = 0
        while i <= tail:
            # print(i, nums[i], nums[head], nums[tail])
            if nums[i] == 0:
                nums[head], nums[i] = nums[i], nums[head]
                head += 1
            elif nums[i] == 2:
                nums[tail], nums[i] = nums[i], nums[tail]
                tail -= 1
                i -= 1
            i += 1
