"""
# Definition for a Node.
class Node:
    def __init__(self, val = 0, neighbors = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []
"""

from typing import Optional
class Solution:
    def cloneGraph(self, node: Optional['Node']) -> Optional['Node']:
        visited = dict()

        return self.dfs(node, visited)

    def dfs(self, node, visited):
        if not node:
            return None

        if node.val in visited:
            return visited[node.val]

        clone = Node(node.val)
        visited[node.val] = clone

        for neigh in node.neighbors:
            n = self.dfs(neigh, visited)
            if n:
                clone.neighbors.append(n)
        
        return clone
        