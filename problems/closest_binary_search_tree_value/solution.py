# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def closestValue(self, root: Optional[TreeNode], target: float) -> int:
        stack = []
        prev = float('-inf')
        cur = root
        
        while stack or cur:
            if cur:
                stack.append(cur)
                cur = cur.left
            else:
                cur = stack.pop()
                if cur.val == target:
                    return cur.val
                
                if target > prev and target < cur.val:
                    return min(prev, cur.val, key = lambda x:abs(target - x))
                
                prev = cur.val
                cur = cur.right
        
        return prev