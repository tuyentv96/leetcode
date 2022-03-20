# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    result = None
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        self.helper(root, p ,q)
        return self.result

    def helper(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> bool:
        if not root:
            return False
        
        left = self.helper(root.left, p, q)
        right = self.helper(root.right, p, q)
        parent = p == root or q == root
        
        if left + right + parent == 2:
            self.result = root
        
        return left or right or parent