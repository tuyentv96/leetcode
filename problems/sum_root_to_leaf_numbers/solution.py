# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sumNumbers(self, root: Optional[TreeNode]) -> int:
        self.sum = 0
        self.helper(root, 0)
        
        return self.sum
        
    def helper(self, root, current_number):
        if not root:
            return
        
        current_number = current_number * 10 + root.val
        
        if not root.left and not root.right:
            self.sum += current_number
            return

        self.helper(root.left, current_number)
        self.helper(root.right, current_number)
        