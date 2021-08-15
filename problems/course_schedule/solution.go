func canFinish(numCourses int, prerequisites [][]int) bool {
	visited := make([]bool, numCourses)
	graph := make([][]int, numCourses)
	dp := make([]bool, numCourses)

	for i := 0; i < len(prerequisites); i++ {
		graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(graph, visited, dp, i) {
			return false
		}
	}

	return true
}

func dfs(graph [][]int, visited []bool, dp []bool, course int) bool {
	if visited[course] {
		return dp[course]
	}

	visited[course] = true
	for i := 0; i < len(graph[course]); i++ {
		if !dfs(graph, visited, dp, graph[course][i]) {
			dp[course] = false
			return false
		}
	}

	dp[course] = true
	return true
}
