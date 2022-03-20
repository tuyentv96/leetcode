"""
# Definition for a Node.
class Node:
    def __init__(self, val: int = 0, left: 'Node' = None, right: 'Node' = None, next: 'Node' = None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next
"""
class Solution:
    def connect(self, root: 'Node') -> 'Node':
        if not root:
            return None

        leftmost = root
        while leftmost:
            cur = leftmost
            head = None
            prev = None
            while cur:
                if cur.left:
                    if not head:
                        head = cur.left
                    if prev:
                        prev.next = cur.left
                    prev = cur.left
                if cur.right:
                    if not head:
                        head = cur.right
                    if prev:
                        prev.next = cur.right
                    prev = cur.right

                cur = cur.next
            leftmost = head
        
        return root
            