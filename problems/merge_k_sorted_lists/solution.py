from heapq import *
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        heap = []
        queue = deque()
        root = None
        cur = None
        prev = None

        for l in lists:
            queue.appendleft(l)

        while len(queue) > 0 or len(heap) > 0:
            prev = cur
            for i in range(len(queue)):
                l = queue.pop()
                if l:
                    heappush(heap, l.val)
                    if l.next:
                        queue.appendleft(l.next)
            
            if heap:
                val = heappop(heap)

                cur = ListNode(val)
                if not root:
                    root = cur
                if prev:
                    prev.next = cur
            
        return root