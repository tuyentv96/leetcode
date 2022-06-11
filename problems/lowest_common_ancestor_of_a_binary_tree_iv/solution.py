# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', nodes: 'List[TreeNode]') -> 'TreeNode':
        self.result = None
        self.helper(root, nodes)
        return self.result
    
    def helper(self, root: 'TreeNode', nodes: 'List[TreeNode]') -> int:
        if not root:
            return 0
        
        l = self.helper(root.left, nodes)
        r = self.helper(root.right, nodes)
        
        cur = 1 if root in nodes else 0
        
        if l + r + cur == len(nodes) and not self.result:
            self.result = root
        return l + r + cur
        