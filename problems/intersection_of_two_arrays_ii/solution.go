func intersect(nums1 []int, nums2 []int) []int {
	var result []int
	m := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]]++
	}

	for i := 0; i < len(nums2); i++ {
		if v, ok := m[nums2[i]]; ok && v > 0 {
			result = append(result, nums2[i])
			m[nums2[i]]--
		}
	}

	return result
}