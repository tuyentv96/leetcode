# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def countUnivalSubtrees(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0

        self.count = 0
        self.helper(root)
        return self.count

    def helper(self, root) -> bool:        
        if not root.left and not root.right:
            self.count += 1
            return True
        
        is_uni = True
        
        if root.left:
            is_uni = self.helper(root.left) and root.val == root.left.val 
        
        if root.right:
            is_uni = self.helper(root.right) and root.val == root.right.val and is_uni

        self.count += is_uni
        return is_uni
