#
# @lc app=leetcode id=6 lang=python3
#
# [6] ZigZag Conversion
#


class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1:
            return s
        res_dict = {}
        i, incr = 0, True
        for c in s:
            if incr:
                i += 1
            elif not incr:
                i -= 1
            if i == numRows:
                incr = False
            elif i == 1:
                incr = True
            # print(i, c)
            res_dict.setdefault(i, []).append(c)
        # print(res_dict)
        result = ""
        for i in range(1, numRows+1, 1):
            result += "".join(res_dict.get(i, []))
        return result
