package main

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h1, h2, h3 := head, head.Next, head.Next
	for h2 != nil && h2.Next != nil {
		h1.Next = h2.Next
		h1 = h1.Next
		h2.Next = h1.Next
		h2 = h2.Next
	}
	h1.Next = h3

	return head
}
