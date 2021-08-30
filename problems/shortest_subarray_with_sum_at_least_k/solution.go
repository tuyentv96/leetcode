type Dequeue struct {
	arr []int
}

func NewDequeue() Dequeue {
	return Dequeue{}
}

func (d *Dequeue) PushFront(i int) {
	d.arr = append([]int{i}, d.arr...)
}

func (d *Dequeue) PushBack(i int) {
	d.arr = append(d.arr, i)
}

func (d *Dequeue) PopFront() int {
	i := d.arr[0]
	d.arr = d.arr[1:]
	return i
}

func (d *Dequeue) PopBack() int {
	i := d.arr[len(d.arr)-1]
	d.arr = d.arr[:len(d.arr)-1]
	return i
}

func (d *Dequeue) IsEmpty() bool {
	return len(d.arr) == 0
}

func (d *Dequeue) Front() int {
	return d.arr[0]
}

func (d *Dequeue) Back() int {
	return d.arr[len(d.arr)-1]
}

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	res := n + 1
	dequeue := NewDequeue()

	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	for i := 0; i <= n; i++ {
		for !dequeue.IsEmpty() && preSum[i]-preSum[dequeue.Front()] >= k {
			res = min(res, i-dequeue.PopFront())
		}

		for !dequeue.IsEmpty() && preSum[i] <= preSum[dequeue.Back()] {
			dequeue.PopBack()
		}

		dequeue.PushBack(i)
	}

	if res == n+1 {
		return -1
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}