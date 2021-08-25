func findKthLargest(nums []int, k int) int {
    return quickSelect(nums,len(nums)-k)
}

func quickSelect(nums []int, i int) int{
   left,right:=0,len(nums)-1
    for left<right{
        pivot:=partition(nums,left,right)
        if pivot==i{
            return nums[pivot]
        }
        
        if pivot>i{
            right=pivot-1
        }else{
            left=pivot+1
        }
    }
    
    return nums[left]
}


func partition(nums []int,left,right int) int{
    pivot:=right
    j:=left

    for i:=left;i<right;i++{
        if nums[i]<nums[pivot]{
            nums[i],nums[j]=nums[j],nums[i]
            j++
        }
    }
    
    nums[pivot],nums[j]=nums[j],nums[pivot]
    return j
}