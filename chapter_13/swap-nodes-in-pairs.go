package main

func main() {
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := head.Next
	cur := head
	for cur != nil && cur.Next != nil {
		node1 := cur
		node2 := cur.Next
		if node2.Next != nil {
			if node2.Next.Next != nil {
				node1.Next = node2.Next.Next
			} else {
				node1.Next = node2.Next
			}
		} else {
			node1.Next = nil
		}

		cur = node2.Next
		node2.Next = node1
	}
	return res
}
