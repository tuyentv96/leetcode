# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from queue import Queue
class Solution:
    def findLeaves(self, root: Optional[TreeNode]) -> List[List[int]]:
        self.result = []
        self.getHeight(root)
        return self.result
        
    def getHeight(self, root: TreeNode) -> int:
        if not root:
            return -1
        
        lh = self.getHeight(root.left)
        rh = self.getHeight(root.right)
        
        height = max(lh, rh) + 1
        
        if len(self.result) == height:
            self.result.append([])
        
        self.result[height].append(root.val)
        return height
        
        
        
        