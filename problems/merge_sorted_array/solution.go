func merge(nums1 []int, m int, nums2 []int, n int)  {
    var first,second = m-1,n-1
    for i:=m+n-1;i>=0;i--{
        if second < 0{
            break
        }
        
        if first>=0 && nums1[first]>nums2[second]{
            nums1[i]=nums1[first]
            first--
        }else{
            nums1[i]=nums2[second]
            second--   
        }
    }
}