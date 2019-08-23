#
# @lc app=leetcode id=138 lang=python3
#
# [138] Copy List with Random Pointer
#
"""
# Definition for a Node.
class Node:
    def __init__(self, val, next, random):
        self.val = val
        self.next = next
        self.random = random
"""


# class Node:
#     def __init__(self, val, next, random):
#         self.val = val
#         self.next = next
#         self.random = random


class Solution:

    node_dict = {}

    def getOrCreateNode(self, old_node):
        if self.node_dict.get(old_node):
            return self.node_dict[old_node]
        else:
            if old_node is None:
                new_node = None
            else:
                new_node = Node(old_node.val, None, None)
            self.node_dict[old_node] = new_node
            return new_node

    def copyRandomList(self, head):
        new_head = None
        curr_node = head
        while curr_node is not None:
            print(curr_node, curr_node.next, curr_node.random)
            new_node = self.getOrCreateNode(curr_node)
            if new_head is None:
                new_head = new_node
            new_next_node = self.getOrCreateNode(curr_node.next)
            new_random_node = self.getOrCreateNode(curr_node.random)
            new_node.next = new_next_node
            new_node.random = new_random_node
            curr_node = curr_node.next
        # print(self.node_dict)
        return new_head


# if __name__ == "__main__":
#     node1 = Node(1, None, None)
#     node2 = Node(2, None, None)
#     node1.next = node2
#     node1.random = node2
#     node2.random = node2
#     node2.next = None
#     head = node1
#     Solution().copyRandomList(head)
