/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var iterator *ListNode
	var root *ListNode
	var complement int
	for l1 != nil || l2 != nil {
		newLn := &ListNode{
			Val:  0,
			Next: nil,
		}

		if l1 != nil {
			newLn.Val += l1.Val
		}
		if l2 != nil {
			newLn.Val += l2.Val
		}

		if complement > 0 {
			newLn.Val += complement
			complement = 0
		}

		if root == nil {
			root = newLn
		}

		if iterator == nil {
			iterator = newLn
		} else {
			iterator.Next = newLn
			iterator = iterator.Next
		}

		if newLn.Val > 9 {
			complement = newLn.Val / 10
			newLn.Val = newLn.Val % 10
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if complement > 0 && iterator != nil {
		iterator.Next = &ListNode{
			Val:  complement,
			Next: nil,
		}
		iterator = iterator.Next
	}

	return root
}