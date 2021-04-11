package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param head ListNode类 the head node
 * @return ListNode类
 */
func sortInList(head *ListNode) *ListNode {
	// write code here

	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	part := slow.Next
	slow.Next = nil

	left := sortInList(head)
	right := sortInList(part)

	current := &ListNode{Val: 0}
	result := current
	for ; left != nil && right != nil; current = current.Next {
		if left.Val <= right.Val {
			current.Next = left
			left = left.Next
		} else {
			current.Next = right
			right = right.Next
		}
	}

	if left != nil {
		current.Next = left
	} else {
		current.Next = right
	}

	return result.Next
}
