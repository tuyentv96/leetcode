# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        if not preorder:
            return None

        self.idx_map = {val:i for i,val in enumerate(inorder)}
        self.preorder = preorder
        self.preorder_idx = 0
        return self.helper(0, len(inorder) - 1)
        
    def helper(self, left: int, right: int) -> Optional[TreeNode]:
        if left > right:
            return None
        
        val = self.preorder[self.preorder_idx]
        self.preorder_idx += 1
        idx = self.idx_map[val]
        root = TreeNode(val)
        
        root.left = self.helper(left, idx - 1)
        root.right = self.helper(idx + 1, right)
        
        return root
        