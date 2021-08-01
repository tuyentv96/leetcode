/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
    return buildBST(head,nil)
}

func buildBST(head *ListNode,tail *ListNode) *TreeNode{
    if head==tail{
        return nil
    }
    
    slow,fast:=head,head
    for fast!=tail && fast.Next!=tail{
        fast=fast.Next.Next
        slow=slow.Next
    }
    
    var root TreeNode
    root.Val=slow.Val
    root.Left=buildBST(head,slow)
    root.Right=buildBST(slow.Next,tail)
    return &root
}


