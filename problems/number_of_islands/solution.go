func numIslands(grid [][]byte) int {
    if len(grid)==0 || len(grid[0])==0{
        return 0
    }
    
    var result int
    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[i]);j++{
            if grid[i][j]=='1'{
                result+=dfs(grid,i,j)
            }
        }
    }
    
    return result
}


func dfs(grid [][]byte, i,j int) int{
    if i<0 || i>=len(grid) || j<0 || j>=len(grid[i]) || grid[i][j]=='0'{
        return 0
    }
    
    grid[i][j]='0'
    dfs(grid,i+1,j)
    dfs(grid,i-1,j)
    dfs(grid,i,j+1)
    dfs(grid,i,j-1)
    
    return 1
}