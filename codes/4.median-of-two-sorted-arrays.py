#
# @lc app=leetcode id=4 lang=python3
#
# [4] Median of Two Sorted Arrays
#


class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        len1, len2 = len(nums1), len(nums2)
        length = len1 + len2
        mid = int(length/2)
        pre, curr = 0, 0
        if len(nums1) == 0:
            if length == 1:
                return nums2[0]
            pre, curr = nums2[mid-1], nums2[mid]
        elif len(nums2) == 0:
            if length == 1:
                return nums1[0]
            pre, curr = nums1[mid-1], nums1[mid]
        else:
            m, n = 0, 0
            for i in range(length):
                pre = curr
                if m < len1 and n < len2:
                    if nums1[m] <= nums2[n]:
                        curr = nums1[m]
                        m += 1
                    else:
                        curr = nums2[n]
                        n += 1
                elif m == len1:
                    curr = nums2[n]
                    n += 1
                elif n == len2:
                    curr = nums1[m]
                    m += 1
                if i == mid:
                    break
        if length % 2 == 0:
            return (pre+curr)/2
        else:
            return curr
