func topKFrequent(nums []int, k int) []int {
    frequentMap:=buildFrequentMap(nums)
    frequentArray:=buildFrequentArray(frequentMap)
    
    for i:=k/2-1;i>=0;i--{
        heapify(frequentArray,frequentMap,i,k)
    }
    
    for i:=k;i<len(frequentArray);i++{
        element:=frequentArray[i]
        beginElement:=frequentArray[0]
        
        if frequentMap[element]>frequentMap[beginElement]{
            frequentArray[0]=element
            heapify(frequentArray,frequentMap,0,k)
        }
    }
    
    return frequentArray[:k]
}

func buildFrequentMap(nums []int) map[int]int{
    result:=make(map[int]int)
    for i:=0;i<len(nums);i++{
        result[nums[i]]++
    }
    
    return result
}

func buildFrequentArray(frequentMap map[int]int) []int{
    result:=make([]int,len(frequentMap))
    i:=0
    for k:=range frequentMap{
        result[i]=k
        i++
    }
    
    return result
}

func heapify(nums []int,frequentMap map[int]int,i int,n int){
    smallest:=i
    left,right:=2*i+1,2*i+2
    
    for left<n && frequentMap[nums[left]]<frequentMap[nums[smallest]]{
        smallest=left
    }
    
    for right<n && frequentMap[nums[right]]<frequentMap[nums[smallest]]{
        smallest=right
    }
    
    if smallest!=i{
        nums[i],nums[smallest]=nums[smallest],nums[i]
        heapify(nums,frequentMap,smallest,n)
    }
}