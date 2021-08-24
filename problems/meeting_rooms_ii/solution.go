func minMeetingRooms(intervals [][]int) int {
    if len(intervals)==0{
        return 0
    }
    
    starts,ends:=make([]int,len(intervals)),make([]int,len(intervals))
    
    for i:=0;i<len(intervals);i++{
        starts[i]=intervals[i][0]
        ends[i]=intervals[i][1]
    }
    
    sort.Slice(starts,func(i,j int) bool{
        return starts[i]<starts[j]
    })
    
    sort.Slice(ends,func(i,j int) bool{
        return ends[i]<ends[j]
    })
    
    min:=0
    startIdx:=0
    endIdx:=0
    
    for startIdx<len(intervals) {
        if ends[endIdx]<=starts[startIdx] {
            min--
            endIdx++
        }
        
        min++
        startIdx++
    }
    
    return min
}
