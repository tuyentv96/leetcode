from queue import PriorityQueue

class Solution:
    ListNode.__eq__ = lambda self, other: self.val == other.val
    ListNode.__lt__ = lambda self, other: self.val < other.val
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        head = tail = ListNode(0)
        queue = PriorityQueue()
        
        for l in lists:
            if l:
                queue.put((l.val, l))
                
        while not queue.empty():
            val,node = queue.get()
            tail.next = node
            tail = tail.next
            node = node.next
            if node:
                queue.put((node.val, node))

        return head.next
                
