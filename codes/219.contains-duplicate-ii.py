#
# @lc app=leetcode id=219 lang=python3
#
# [219] Contains Duplicate II
# 实际意思是找同数字最小间隔，若不超过 k 则满足条件


class Solution:
    def containsNearbyDuplicate(self, nums, k: int) -> bool:
        results = {}
        # 初始化值，保证这两个值肯定不符合条件
        diff = k+1
        for i, n in enumerate(nums):
            if results.get(n) is not None:
                if i - results[n] <= k:
                    return True
            results[n] = i
        return False


# if __name__ == "__main__":
#     print(Solution().containsNearbyDuplicate([1, 2, 3, 1], 3))
#     print(Solution().containsNearbyDuplicate([1, 0, 1, 1], 1))
#     print(Solution().containsNearbyDuplicate([1, 2, 3, 1, 2, 3], 2))
#     print(Solution().containsNearbyDuplicate([1, 2, 3, 1, 2, 3], 3))
