class Solution:
    def validPath(self, n: int, edges: List[List[int]], source: int, destination: int) -> bool:
        adj_list = [[] for _ in range(n)]
        for edge in edges:
            adj_list[edge[0]].append(edge[1])
            adj_list[edge[1]].append(edge[0])

        queue = collections.deque([source])
        seen = set([source])
        
        while queue:
            cur = queue.popleft()
            if cur == destination:
                return True
            
            for nb in adj_list[cur]:
                if nb not in seen:
                    seen.add(nb)
                    queue.append(nb)
        
        return False
        
        
        
        
        