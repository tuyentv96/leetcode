func spiralOrder(matrix [][]int) []int {
    
    if len(matrix)==0{
        return []int{}
    }
    
    m,n:=len(matrix),len(matrix[0])
    dir:=1
    top,bot,left,right:=0,m-1,0,n-1
    result:=make([]int,0,m*n)

    for top<=bot && left<=right{
        switch dir{
            //left->right
            case 1:
            for i:=left;i<=right;i++{
                result=append(result,matrix[top][i])
            }
            
            top++
            dir=2
            // top->bot
            case 2:
            for i:=top;i<=bot;i++{
                result=append(result,matrix[i][right])
            }
            
            right--
            dir=3
            // right->left
            case 3:
            for i:=right;i>=left;i--{
                result=append(result,matrix[bot][i])
            }
            
            bot--
            dir=4
            // bot->top
            case 4:
            for i:=bot;i>=top;i--{
                result=append(result,matrix[i][left])
            }
            
            left++
            dir=1
        }
    }
    
    return result
}