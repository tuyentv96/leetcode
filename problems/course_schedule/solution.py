from collections import defaultdict

class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        graph = defaultdict(list)
        for p in prerequisites:
            graph[p[1]].append(p[0])
        
        visited = [False] * numCourses
        current_path = [False] * numCourses
        for i in range(numCourses):
            if not visited[i] and self.is_cycle(graph, visited, current_path, i):
                return False
        return True

    def is_cycle(self, graph, visited, current_path, i):
        if visited[i]:
            return current_path[i]

        visited[i] = True
        current_path[i] = True

        for j in graph[i]:
            if self.is_cycle(graph, visited, current_path, j):
                return True
        
        current_path[i] = False
        return False
