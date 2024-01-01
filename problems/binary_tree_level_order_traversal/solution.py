# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
# Input: root = [3,9,20,null,null,15,7]
# queue = [3], arr=[3]
# queue = [9, 20], arr = [9,20]
# queue = [15, 7], arr = [15,7]
class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        if not root:
            return []
        queue = deque()
        queue.appendleft(root)
        res = []
        while len(queue) > 0:
            size = len(queue)
            arr = []
            for i in range(size):
                node = queue.pop()
                arr.append(node.val)
                if node.left:
                    queue.appendleft(node.left)
                if node.right:
                    queue.appendleft(node.right)    
            res.append(arr)             
        
        return res
