func leastInterval(tasks []byte, n int) int {
	charCount := make([]int, 26)
	for _, c := range tasks {
		charCount[c-'A']++
	}

	sort.Ints(charCount)
	maxValue := charCount[25] - 1
	idleSlot := maxValue * n
	for i := 0; i < 25; i++ {
		idleSlot -= int(math.Min(float64(charCount[i]), float64(maxValue)))
	}

	if idleSlot > 0 {
		return len(tasks) + idleSlot
	}

	return len(tasks)
}
