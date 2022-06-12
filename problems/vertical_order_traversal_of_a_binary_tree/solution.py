# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from collections import defaultdict
class Solution:
    def verticalTraversal(self, root: Optional[TreeNode]) -> List[List[int]]:
        stack = [(root, 0, 0)]
        result = []
        nodes = defaultdict(list)
        
        while stack:
            cur, row, col = stack.pop()
            nodes[col].append((row, cur.val))
            
            if cur.right:
                stack.append((cur.right, row+1, col+1))
            
            if cur.left:
                stack.append((cur.left, row+1, col-1))
        result = []
        for col in sorted(nodes.keys()):
            result.append([val for row, val in sorted(nodes[col])])
        return result