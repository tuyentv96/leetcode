# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def buildTree(self, inorder: List[int], postorder: List[int]) -> Optional[TreeNode]:
        if not postorder:
            return None

        self.postorder = postorder
        self.index_map = {val:i for i,val in enumerate(inorder)}
        return self.helper(0, len(inorder)-1)
        
    def helper(self, left: int, right: int) -> Optional[TreeNode]:
        if left > right:
            return None
        
        val = self.postorder.pop()
        root = TreeNode(val)
        idx = self.index_map[val]

        root.right = self.helper(idx+1, right)
        root.left = self.helper(left, idx-1)
        
        return root
        
        