class Solution:
    def allPathsSourceTarget(self, graph: List[List[int]]) -> List[List[int]]:        
        self.seen = set()
        self.result = []
        self.tmp = []
        self.dfs(0, len(graph) - 1, graph)
        return self.result
    
    def dfs(self, node: int, n: int, graph: List[List[int]]):
        if node in self.seen:
            return
        
        self.seen.add(node)
        self.tmp.append(node)
        
        if node == n:
            self.result.append(self.tmp.copy())
        
        for nb in graph[node]:
            self.dfs(nb, n, graph)
        
        self.seen.remove(node)
        self.tmp.pop()
        
        
                