/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    newNode:=&TreeNode{Val: val}
    if root==nil{
        return newNode
    }
    
    var prev *TreeNode
    cur:=root
    
    for cur!=nil{
        prev=cur
        
        if val<cur.Val{
            cur=cur.Left
        }else{
            cur=cur.Right
        }
    }
    
    
    if val<prev.Val{
        prev.Left=newNode
    }else{
        prev.Right=newNode
    }
    
    return root
}