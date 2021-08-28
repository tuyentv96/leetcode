func maxArea(height []int) int {
    left,right:=0,len(height)-1
    var maxSoFar int
    
    for left<right{
        maxSoFar=max(maxSoFar,(right-left)*min(height[left],height[right]))
        
        if height[left]<height[right]{
            left++
        }else{
            right--
        }
    }
    
    return maxSoFar
}


func max(a,b int) int{
    if a>b{
        return a
    }
    
    return b
}

func min(a,b int) int{
    if a<b{
        return a
    }
    
    return b
}