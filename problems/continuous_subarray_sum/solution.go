func checkSubarraySum(nums []int, k int) bool {
	sum := 0
	m := make(map[int]int)
    m[0]=-1
	for i := range nums {
		sum += nums[i]
		sum = sum % k
		if _, ok := m[sum]; ok {
			if i-m[sum]>=2 {
				return true
			}
		}else {
			m[sum]=i
		}
	}
	
	return false
}