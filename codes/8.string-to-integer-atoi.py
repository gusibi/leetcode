#
# @lc app=leetcode id=8 lang=python3
#
# [8] String to Integer (atoi)
#

nums = set(["1", "2", "3", "4", "5", "6", "7", "8", "9", "0"])


class Solution:
    def myAtoi(self, str: str) -> int:
        numstr = str.strip()
        flag = 1
        if numstr.startswith("+"):
            numstr = numstr[1:]
        elif numstr.startswith("-"):
            flag = -1
            numstr = numstr[1:]
        n = len(numstr)
        _numstr = ""
        for i in range(n):
            if numstr[i] in nums:
                _numstr += numstr[i]
            elif _numstr and numstr[i] not in nums:
                break
            elif not _numstr and numstr[i] not in nums:
                return 0
        # print(_numstr)
        if not _numstr:
            return 0
        n = int(_numstr)
        result = flag * n
        if result >= 0:
            return min(result, 2147483647)
        else:
            return max(result, -2147483648)


# if __name__ == "__main__":
#     print(Solution().myAtoi("4193 with words"))
#     print(Solution().myAtoi("+4193 with words"))
#     print(Solution().myAtoi("+4 193 with words"))
#     print(Solution().myAtoi("+0 193 with words"))
#     print(Solution().myAtoi("words and 987"))
