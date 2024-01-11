# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        if not root:
            return True
        
        return self.helper(float('-inf'), float('inf'), root)
        
    def helper(self, l, r, node):
        if not node:
            return True
        
        if node.val <= l or node.val >= r:
            return False
        
        return self.helper(l, node.val, node.left) and self.helper(node.val, r, node.right)