/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    inorderMap:=make(map[int]int)
    for i:=0;i<len(inorder);i++{
        inorderMap[inorder[i]]=i
    }
    
    var preorderIndex int
    return build(inorderMap,preorder,&preorderIndex,0,len(preorder)-1)
}

func build(inorderMap map[int]int,preorder []int,preorderIndex *int,left,right int) *TreeNode{
    if left>right{
        return nil
    }
    
    rootValue:=preorder[*preorderIndex]
    *preorderIndex++
    
    var root TreeNode
    root.Val=rootValue
    
    rootInorderIndex:=inorderMap[rootValue]
    
    root.Left=build(inorderMap,preorder,preorderIndex,left,rootInorderIndex-1)
    root.Right=build(inorderMap,preorder,preorderIndex,rootInorderIndex+1,right)
    
    return &root
}