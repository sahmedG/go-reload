package proj

func DeleteNode(head *Node, index int) *Node {
	// If the head is nil or the index is negative, return the head
	if head == nil || index < 0 {
		return head
	}

	// If the index is 0, delete the head and return the new head
	if index == 0 {
		return head.Next
	}

	// Traverse the list until the node before the one to be deleted
	curr := head
	for i := 0; i < index-1; i++ {
		if curr.Next == nil {
			// If the index is greater than or equal to the length of the list,
			// return the original head
			return head
		}
		curr = curr.Next
	}

	// Delete the node at the specified index and return the original head
	curr.Next = curr.Next.Next
	return head
}
