# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from queue import Queue
from collections import defaultdict
class Solution:
    def verticalOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        if not root:
            return []
        
        columns = defaultdict(list)
        queue = Queue()
        queue.put((root, 0))
        
        while not queue.empty():
            size = queue.qsize()
            for _ in range(size):
                cur, i = queue.get()
                columns[i].append(cur.val)

                if cur.left:
                    queue.put((cur.left, i - 1))
                if cur.right:
                    queue.put((cur.right, i + 1))

        return [columns[x] for x in sorted(columns.keys())]
        