/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
    for node.Next !=nil{
        next:=node.Next
        node.Val=next.Val
        if next.Next==nil{
            node.Next=nil
            return
        }
        node=next
    }
}