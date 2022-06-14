# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def findDuplicateSubtrees(self, root: Optional[TreeNode]) -> List[Optional[TreeNode]]:
        def trv(root):
            if root:
                struct = frozenset([(root.val, trv(root.right), trv(root.left))])
                nodes[struct].append(root)
                return struct
        
        nodes = collections.defaultdict(list)
        trv(root)
        return [node[0] for node in nodes.values() if len(node) > 1]