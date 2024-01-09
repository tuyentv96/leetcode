# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        idx_map = {}
        for i in range(len(inorder)):
            idx_map[inorder[i]] = i
        
        self.i = 0
        return self.build(0, len(preorder) - 1, preorder, idx_map)

    def build(self, l, r, preorder, idx_map):
        if l > r:
            return
        
        val = preorder[self.i]
        idx = idx_map[val]
        node = TreeNode(val)
        self.i += 1
        node.left = self.build(l, idx - 1, preorder, idx_map)
        node.right = self.build(idx + 1, r, preorder, idx_map)
        return node