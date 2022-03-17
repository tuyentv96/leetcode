# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def postorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        result, stack = [], []
        if root is None:
            return result
        
        cur = root
        prev = None
        while stack or cur:
            if cur:
                stack.append(cur)
                cur = cur.left
            else:
                cur = stack[-1]
                if cur.right is None or prev == cur.right:
                    stack.pop()
                    result.append(cur.val)
                    prev = cur
                    cur = None
                else:
                    cur = cur.right
        
        return result
            