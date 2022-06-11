"""
# Definition for a Node.
class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
"""

class Solution:
    def treeToDoublyList(self, root: 'Optional[Node]') -> 'Optional[Node]':
        if not root:
            return None
        head = None
        stack = []
        cur = root
        prev = None
        
        while stack or cur:
            if cur:
                stack.append(cur)
                cur=cur.left
            else:
                cur = stack.pop()
                cur.left = prev
                if prev:
                    prev.right = cur
                prev = cur
                if not head:
                    head = cur
                cur=cur.right
        
        head.left=prev
        prev.right=head
        
        return head
        