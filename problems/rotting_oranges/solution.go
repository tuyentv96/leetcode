var directions = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	return &Queue{
		list: list.New(),
	}
}

func (q *Queue) Enqueue(e []int) {
	q.list.PushBack(e)
}

func (q *Queue) Dequeue() []int {
	e := q.list.Front()
	q.list.Remove(e)
	return e.Value.([]int)
}

func (q *Queue) Size() int {
	return q.list.Len()
}

func (q *Queue) IsEmpty() bool {
	return q.list.Len() == 0
}

func orangesRotting(grid [][]int) int {
	countFresh := 0
	queue := NewQueue()

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				countFresh++
			} else if grid[i][j] == 2 {
				queue.Enqueue([]int{i, j})
			}
		}
	}

	if countFresh == 0 {
		return 0
	}

	count := 0
	for !queue.IsEmpty() {
		count++
		size := queue.Size()

		for j := 0; j < size; j++ {
			e := queue.Dequeue()

			nbs := getNeighboors(grid, e[0], e[1])
			for i := 0; i < len(nbs); i++ {
				r, c := nbs[i][0], nbs[i][1]
				if grid[r][c] == 0 || grid[r][c] == 2 {
					continue
				}

				countFresh--
				grid[r][c] = 2
				queue.Enqueue([]int{r, c})
			}
		}

	}

	if countFresh > 0 {
		return -1
	}

	return count - 1
}

func getNeighboors(grid [][]int, row, col int) [][]int {
	m, n := len(grid), len(grid[0])
	var result [][]int

	for i := 0; i < len(directions); i++ {
		r, c := row+directions[i][0], col+directions[i][1]

		if r < 0 || r >= m || c < 0 || c >= n {
			continue
		}

		result = append(result, []int{r, c})
	}

	return result
}