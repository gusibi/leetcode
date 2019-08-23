#
# @lc app=leetcode id=56 lang=python3
#
# [56] Merge Intervals
#


class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        if not intervals:
            return []
        intervals = sorted(intervals, key=lambda s: s[0])
        # print(intervals)
        results = []
        left, right = intervals[0]
        n = len(intervals)
        for i in range(n):
            if right <= intervals[i][1]:
                right = intervals[i][1]
            if i == n - 1:
                results.append([left, right])
            else:
                if intervals[i + 1][0] > right:
                    results.append([left, right])
                    left = intervals[i+1][0]
        return results
