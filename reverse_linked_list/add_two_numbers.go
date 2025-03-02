package reverse_linked_list

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersHelper(l1, l2, 0)
}

func addTwoNumbersHelper(l1 *ListNode, l2 *ListNode, sum int) *ListNode {
	if l1 == nil && l2 == nil && sum == 0 {
		return nil
	}
	if l1 != nil {
		sum += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		sum += l2.Val
		l2 = l2.Next
	}
	return &ListNode{
		Val:  sum % 10,
		Next: addTwoNumbersHelper(l1, l2, sum/10),
	}
}
