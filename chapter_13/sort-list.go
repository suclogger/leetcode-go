package main

// 148 sort-list

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
	}
	sortList(head)
	t := head
	for t != nil {
		print(t.Val, "->")
		t = t.Next
	}
	println()
}

func sortList(head *ListNode) *ListNode {

	length, c := 0, head
	for c != nil {
		c = c.Next
		length++
	}
	dummy := new(ListNode)
	dummy.Next = head
	for l := 1; l < length; l = l << 1 {
		cur, pre := dummy.Next, dummy
		for cur != nil {
			head1 := cur
			for i := 1; i < l && cur.Next != nil; i++ {
				cur = cur.Next
			}
			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < l && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			var next *ListNode = nil
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			pre.Next = merge(head1, head2)
			for pre.Next != nil {
				pre = pre.Next
			}
			cur = next

		}
	}
	return dummy.Next
}

func merge(h *ListNode, cur *ListNode) *ListNode {
	dummy := new(ListNode)
	c, head1, head2 := dummy, h, cur
	for head1 != nil && head2 != nil {
		if head1.Val <= head2.Val {
			c.Next = head1
			head1 = head1.Next
		} else {
			c.Next = head2
			head2 = head2.Next
		}
		c = c.Next
	}
	if head1 == nil {
		c.Next = head2
	} else {
		c.Next = head1
	}
	return dummy.Next
}
