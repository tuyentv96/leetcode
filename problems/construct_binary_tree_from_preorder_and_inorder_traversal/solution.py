# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        inorder_map = {}
        self.count = 0
        for i, val in enumerate(inorder):
            inorder_map[val] = i
        
        return self.helper(preorder, inorder_map, 0, len(preorder) - 1)
    
    def helper(self, preorder, inorder_map, l, r):
        if l > r:
            return None

        val = preorder[self.count]
        idx = inorder_map[val]
        self.count += 1

        node = TreeNode(val)
        node.left = self.helper(preorder, inorder_map, l, idx - 1)
        node.right = self.helper(preorder, inorder_map, idx + 1, r)
        
        return node