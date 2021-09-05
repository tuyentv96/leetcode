/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    var preorderIndex int
    inorderMap:=make(map[int]int)
    
    for i:=0;i<len(inorder);i++{
        inorderMap[inorder[i]]=i
    }
    
    return build(preorder,&preorderIndex,inorderMap,0,len(preorder)-1)
}


func build(preorder []int,preorderIndex *int,inorderMap map[int]int,l,r int) *TreeNode {
    if l>r{
        return nil
    }
    
    root:=&TreeNode{}
    root.Val = preorder[*preorderIndex]
    index:=inorderMap[root.Val]
    *preorderIndex++
    
    root.Left=build(preorder,preorderIndex,inorderMap,l,index-1)
    root.Right=build(preorder,preorderIndex,inorderMap,index+1,r)
    
    return root
}