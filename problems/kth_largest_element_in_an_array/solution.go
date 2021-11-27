func findKthLargest(nums []int, k int) int {
	n := len(nums)

	if k > n {
		return 0
	}
    
    
    for i:=k-1;i>=0;i--{
        heapify(nums,i,k)
	}

	for i := k; i < n; i++ {
        if nums[i]>nums[0]{
            nums[0]=nums[i]
            heapify(nums,0,k)
        }
	}

	return nums[0]
}

// min heap
func heapify(nums []int, i int, n int) {
	smallest := i

	left, right := 2*i+1, 2*i+2

	if left < n && nums[left] < nums[smallest] {
		smallest = left
	}

	if right < n && nums[right] < nums[smallest] {
		smallest = right
	}

	if smallest != i {
		nums[smallest], nums[i] = nums[i], nums[smallest]
		heapify(nums, smallest, n)
	}
}
