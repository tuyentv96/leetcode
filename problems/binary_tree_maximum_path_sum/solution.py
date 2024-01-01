# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def maxPathSum(self, root: Optional[TreeNode]) -> int:
        
        self.res = float('-inf')
        self.dfs(root)
        return self.res

    def dfs(self, node):
        if not node:
            return 0

        left = self.dfs(node.left)
        right = self.dfs(node.right)

        max_split = max(node.val, max(left, right) + node.val)
        self.res = max(self.res, max_split, left + right + node.val)
        return max_split