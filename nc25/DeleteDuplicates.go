package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param head ListNode类
 * @return ListNode类
 */
func deleteDuplicates(head *ListNode) *ListNode {
	// write code here

	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Val: head.Val - 1, Next: head}
	for prev, curr := dummy, head; curr != nil; curr = curr.Next {
		if prev.Val == curr.Val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
	}

	return dummy.Next
}

func main() {
	nine3 := &ListNode{Val: 9}
	nine2 := &ListNode{Val: 9, Next: nine3}
	nine1 := &ListNode{Val: 9, Next: nine2}
	eight4 := &ListNode{Val: 8, Next: nine1}
	eight3 := &ListNode{Val: 8, Next: eight4}
	eight2 := &ListNode{Val: 8, Next: eight3}
	eight1 := &ListNode{Val: 8, Next: eight2}
	seven := &ListNode{Val: 7, Next: eight1}
	six2 := &ListNode{Val: 6, Next: seven}
	six1 := &ListNode{Val: 6, Next: six2}
	five3 := &ListNode{Val: 5, Next: six1}
	five2 := &ListNode{Val: 5, Next: five3}
	five1 := &ListNode{Val: 5, Next: five2}
	four2 := &ListNode{Val: 4, Next: five1}
	four1 := &ListNode{Val: 4, Next: four2}
	three := &ListNode{Val: 3, Next: four1}
	two := &ListNode{Val: 2, Next: three}
	one2 := &ListNode{Val: 1, Next: two}
	one1 := &ListNode{Val: 1, Next: one2}
	zero5 := &ListNode{Val: 0, Next: one1}
	zero4 := &ListNode{Val: 0, Next: zero5}
	zero3 := &ListNode{Val: 0, Next: zero4}
	zero2 := &ListNode{Val: 0, Next: zero3}
	zero1 := &ListNode{Val: 0, Next: zero2}

	for curr := zero1; curr != nil; curr = curr.Next {
		fmt.Printf(" -> %d", curr.Val)
	}
	fmt.Println()

	result := deleteDuplicates(zero1)
	for curr := result; curr != nil; curr = curr.Next {
		fmt.Printf(" -> %d", curr.Val)
	}
	fmt.Println()
}
