package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return findLatest(head)
}

func findLatest(parent *ListNode) *ListNode {
	if parent.Next == nil {
		return parent
	}
	result := findLatest(parent.Next)

	// 4->5-> =5->4
	// 4->nil

	// 3->4 = 4->3
	// 3->nil

	// 2->3 = 3->2
	// 2->nil

	// 1->2 = 2->1
	// 1->nil
	parent.Next.Next = parent
	parent.Next = nil

	return result
}
