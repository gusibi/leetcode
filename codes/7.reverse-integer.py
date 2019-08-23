#
# @lc app=leetcode id=7 lang=python3
#
# [7] Reverse Integer
#


class Solution:
    def reverse(self, x: int) -> int:
        flag = -1 if x < 0 else 1
        result = 0
        x = flag*x
        while True:
            x, m = divmod(x, 10)
            # print(x, m)
            if result == 0:
                result = m
            else:
                result = result*10+m
            # print(result)
            if x == 0:
                break
        result = result*flag
        # if result < -(2 ** 31) or result > 2**31 - 1:
        if result >> 31 < -1 or result >> 31 > 0:
            return 0
        return result
