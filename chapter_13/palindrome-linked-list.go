package main

/**
先找到链表中心，反转后半段，然后逐个比对
*/
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow.Next = reverseList(slow.Next)
	slow = slow.Next
	h := head
	for slow != nil && slow.Val == h.Val {
		slow = slow.Next
		h = h.Next
	}
	return slow == nil
}

func reverseList(node *ListNode) *ListNode {
	var pre, next *ListNode
	cur := node
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		if next == nil {
			return cur
		} else {
			cur = next
		}
	}
	return nil
}
