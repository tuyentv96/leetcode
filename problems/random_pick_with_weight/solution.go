type Solution struct {
	cum []int
}

func Constructor(w []int) Solution {
	sum := 0
	cum := make([]int, len(w))
	for i, x := range w {
		sum += x
		cum[i] = sum
	}

	return Solution{cum: cum}
}

func (this *Solution) PickIndex() int {
	randValue := rand.Intn(this.cum[len(this.cum)-1])+1
	left := 0
	right := len(this.cum) - 1
	for left < right {
		mid := left + (right-left)/2
		if this.cum[mid] == randValue {
			return mid
		} else if randValue < this.cum[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
