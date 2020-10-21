func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var min = 9999999
	var res int
	for i := 0; i < len(nums); i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if min >= abs(sum, target) {
				res = sum
				min = abs(sum, target)
			}

			if sum < target {
				l++
			} else {
				r--
			}
		}
	}

	return res
}


func abs(a,b int) int{
    if a>b{
        return a-b
    }
    
    return b-a
}