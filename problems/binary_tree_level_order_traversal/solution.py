# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from queue import Queue
class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        result = []
    
        if root is None:
            return result

        q = Queue()
        q.put(root)
        
        while not q.empty():
            level = []
            for _ in range(q.qsize()):
                node = q.get()
                level.append(node.val)
                if node.left:
                    q.put(node.left)
                if node.right:
                    q.put(node.right)
            
            result.append(level)
        
        return result
                
        