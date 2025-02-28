package solution

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		graph[prerequisite[0]] = append(graph[prerequisite[0]], prerequisite[1])
	}
	visited := make([]bool, numCourses)
	stack := make([]bool, numCourses)
	for i := range numCourses - 1 {
		if dfs(i, &graph, &visited, &stack) {
			return false
		}
	}

	return true
}

func dfs(course int, graph *[][]int, visited, stack *[]bool) bool {
	if (*stack)[course] {
		return true
	}

	if (*visited)[course] {
		return false
	}

	(*stack)[course] = true

	for _, prereq := range (*graph)[course] {
		if dfs(prereq, graph, visited, stack) {
			return true
		}

	}
	(*stack)[course] = false

	return false
}
