# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def levelOrderBottom(self, root: Optional[TreeNode]) -> List[List[int]]:
        result = []
        if not root:
            return result

        stack = []
        stack.append((root, 0))
        
        while stack:
            (cur, i) = stack.pop()
            if len(result) < i+1:
                result.append([])
            
            result[i].append(cur.val)

            if cur.right:
                stack.append((cur.right, i+1))
            if cur.left:
                stack.append((cur.left, i+1))
        
        return reversed(result)
            