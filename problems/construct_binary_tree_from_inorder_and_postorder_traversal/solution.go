/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    inorderMap:=make(map[int]int)
    
    for i:=0;i<len(inorder);i++{
        inorderMap[inorder[i]]=i
    }
    
    postorderIndex:=len(postorder)-1
    return helper(postorder,&postorderIndex,inorderMap,0,len(postorder)-1)
}

func helper(postorder []int,postorderIndex *int,inorderMap map[int]int,l,r int) *TreeNode{
    if l>r{
        return nil
    }
    
    root:=&TreeNode{}
    root.Val=postorder[*postorderIndex]
    *postorderIndex--
    index:=inorderMap[root.Val]
    
    root.Right=helper(postorder,postorderIndex,inorderMap,index+1,r)
    root.Left=helper(postorder,postorderIndex,inorderMap,l,index-1)

    return root
}