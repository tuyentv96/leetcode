from heapq import *
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    ListNode.__eq__ = lambda self, other: self.val == other.val
    ListNode.__lt__ = lambda self, other: self.val < other.val
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        heap = []
        root = ListNode(None)
        cur = root
        
        for node in lists:
            if node:
                heappush(heap, (node.val, node))
        
        while len(heap) > 0:
            cur.next = heappop(heap)[1]
            cur = cur.next
            if cur.next:
                heappush(heap, (cur.next.val, cur.next))
        
        return root.next