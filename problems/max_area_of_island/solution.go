func maxAreaOfIsland(grid [][]int) int {
    res:=0
    
    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[i]);j++{
            res=max(res,dfs(grid,i,j))
        }
    }
    
    return res
}

func dfs(grid [][]int,i,j int) int{
    if (i<0 || j<0 || i>=len(grid) || j>=len(grid[i]) || grid[i][j]==0) {
        return 0
    }
    
    grid[i][j]=0
    return 1 + dfs(grid,i+1,j)+dfs(grid,i-1,j)+dfs(grid,i,j+1)+dfs(grid,i,j-1)
}

func max(a,b int) int{
    if a>b{
        return a
    }
    
    return b
}