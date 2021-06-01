func intersection(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]bool)
	nums2Map := make(map[int]bool)
	for _, num := range nums1 {
		nums1Map[num] = true
	}

	for _, num := range nums2 {
		nums2Map[num] = true
	}

	var res []int
	for num, _ := range nums1Map {
		if _, ok := nums2Map[num]; ok {
			res = append(res, num)
		}
	}

	return res
}
