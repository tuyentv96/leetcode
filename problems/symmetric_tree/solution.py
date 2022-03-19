# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from queue import Queue
class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        queue = Queue()
        queue.put(root)
        queue.put(root)
        
        while not queue.empty():
            left = queue.get()
            right = queue.get()
            if left is None and right is None:
                continue
            
            if left is None or right is None:
                return False
            
            if left.val != right.val:
                return False
            
            queue.put(left.left)
            queue.put(right.right)
            queue.put(left.right)
            queue.put(right.left)
        
        return True