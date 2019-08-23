#
# @lc app=leetcode id=69 lang=python3
#
# [69] Sqrt(x)
#


def sqrt(x, m) -> float:
    # x num
    # m round
    left, right = 0, x
    n = right / 2
    while (right - left) > m:
        if n * n == x:
            return n
        elif n * n > x:
            right = n
        else:
            left = n
        n = (right + left) / 2
    # print("N: ", n, left, right)
    return right


class Solution:
    def mySqrt(self, x: int) -> int:
        if x == 1:
            return 1
        elif x == 0:
            return 0
        return int(sqrt(x, 0.00000001))


# if __name__ == "__main__":
#     Solution().mySqrt(7)
#     Solution().mySqrt(36)
#     Solution().mySqrt(2147395599)
