# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from collections import deque
class Solution:
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        queue = deque()
        queue.append(root)
        
        while queue:
            cur = queue.popleft()
            if not cur:
                continue

            queue.append(cur.left)
            queue.append(cur.right)
            cur.left, cur.right = cur.right, cur.left
            
        return root
            