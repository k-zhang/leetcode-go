package easy

import "github.com/k-zhang/leetcode-go/common"

func mergeTwoLists(l1 *common.ListNode[int], l2 *common.ListNode[int]) *common.ListNode[int] {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	preHead := &common.ListNode[int]{Val: -1, Next: nil}
	p := preHead

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}

		p = p.Next
	}

	if l1 == nil {
		p.Next = l2
	} else {
		p.Next = l1
	}

	return preHead.Next
}
