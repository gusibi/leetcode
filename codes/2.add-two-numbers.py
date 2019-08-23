#
# @lc app=leetcode id=2 lang=python3
#
# [2] Add Two Numbers
#
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        if l1 is None:
            return l2
        elif l2 is None:
            return l1
        list1, list2 = [], []
        while l1 != None:
            list1.append(l1.val)
            l1 = l1.next
        while l2 != None:
            list2.append(l2.val)
            l2 = l2.next

        # print(list1, list2)
        len1, len2 = len(list1), len(list2)
        m = max(len1, len2)
        d = 0
        head, tail = None, None
        for i in range(m):
            # print("index: %d" % i)
            if i < len1 and i < len2:
                v = list1[i] + list2[i] + d
            elif i >= len1:
                v = list2[i] + d
            elif i >= len2:
                v = list1[i] + d
            d, r = divmod(v, 10)
            # print(v, d, r)
            node = ListNode(r)
            if tail:
                tail.next = node
                tail = node
            if head is None:
                head = node
                tail = node
        if d > 0:
            node = ListNode(d)
            tail.next = node
        return head


# if __name__ == "__main__":
#     l1 = ListNode(1)
#     l1.next = ListNode(8)
#     l2 = ListNode(0)
#     # l2.next = ListNode(4)
#     num = Solution().addTwoNumbers(l1, l2)
#     print(num.val, num.next.val)

