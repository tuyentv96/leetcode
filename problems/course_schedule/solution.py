class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        graph = {}
        for p in prerequisites:
            pre, course = p[0], p[1]
            if course not in graph:
                graph[course] = []
            graph[course].append(pre)

        visited = set()
        courses = {}

        for course in range(numCourses):
            if not self.dfs(course, graph, visited, courses):
                return False

        return True
        
    def dfs(self, val, graph, visited, courses):
        if val in visited:
            return courses.get(val, False)

        visited.add(val)
        for v in graph.get(val, []):
            if not self.dfs(v, graph, visited, courses):
                courses[v] = False
                return False
        
        courses[val] = True
        return True