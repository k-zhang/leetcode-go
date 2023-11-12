package common

func ArrayToLinkedList[T any](array []T) *ListNode[T] {
	var head *ListNode[T]
	var tail *ListNode[T]

	var current *ListNode[T]

	for _, i := range array {
		current = &ListNode[T]{Val: i, Next: nil}
		if head == nil {
			head = current
		} else {
			tail.Next = current
		}

		tail = current
	}

	return head
}
