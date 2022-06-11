# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        self.result = None
        self.helper(root, p ,q)
        return self.result
        
    def helper(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> bool:
        if not root:
            return False
        
        l = self.helper(root.left, p, q)
        r = self.helper(root.right, p, q)
        cur = root.val == p.val or root.val == q.val
        
        if l + r + cur == 2:
            self.result = root
        
        return l or r or cur
        