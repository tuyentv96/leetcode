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

        l = self.dfs(node.left)
        r = self.dfs(node.right)

        max_single = max(max(l, r) + node.val, node.val)
        self.res = max(self.res, max_single, l + r + node.val)
        return max_single