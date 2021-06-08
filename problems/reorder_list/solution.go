func reorderList(head *ListNode) {
	slow, fast := head, head
	l1 := head

	// prev tail of first
	// fast tail of second
	// l1 head of first
	// slow head of second
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// break connection with second
	if prev != nil {
		prev.Next = nil
	}

	l2 := reverse(slow)
	merge(l1, l2)
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func merge(l1, l2 *ListNode) {
	for l1 != nil {
		l1Next := l1.Next
		l2Next := l2.Next

		l1.Next = l2
		l2.Next = l1Next
		if l1Next == nil {
			l2.Next = l2Next
			break
		}

		l1 = l1Next
		l2 = l2Next
	}
}
