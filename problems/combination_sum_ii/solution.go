func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	sort.Ints(candidates)
	findcombinationSum(candidates, target, 0, c, &res)
	return res
}

func findcombinationSum(nums []int, target, index int, c []int, res *[][]int) {
	if target <= 0 {
		if target == 0 {
			b := make([]int, len(c))
			copy(b, c)
			*res = append(*res, b)
		}
	}

	for i := index; i < len(nums); i++ {
		if nums[i] > target {
			break
		}

		if i > index && nums[i] == nums[i-1] {
			continue
		}

		c = append(c, nums[i])
		findcombinationSum(nums, target-nums[i], i+1, c, res)
		c = c[:len(c)-1]
	}
}