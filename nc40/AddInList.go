package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversList(head *ListNode) *ListNode {
	var prev *ListNode
	for curr := head; curr != nil; {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

/**
 *
 * @param head1 ListNode类
 * @param head2 ListNode类
 * @return ListNode类
 */
func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
	// write code here

	temp, dummy := 0, &ListNode{Val: 0}
	for node1, node2 := reversList(head1), reversList(head2); ; {
		finished := true
		if node1 != nil {
			temp += node1.Val
			node1 = node1.Next
			finished = false
		}
		if node2 != nil {
			temp += node2.Val
			node2 = node2.Next
			finished = false
		}
		if finished {
			break
		}

		dummy.Next = &ListNode{Val: temp % 10, Next: dummy.Next}
		temp /= 10
	}

	if temp > 0 {
		dummy.Next = &ListNode{Val: temp, Next: dummy.Next}
	}

	return dummy.Next
}

func main() {
	a1 := &ListNode{Val: 7}
	a2 := &ListNode{Val: 3, Next: a1}
	a := &ListNode{Val: 9, Next: a2}

	b1 := &ListNode{Val: 3}
	b := &ListNode{Val: 6, Next: b1}

	c := addInList(a, b)
	for curr := c; curr != nil; curr = curr.Next {
		fmt.Printf("%d", curr.Val)
	}
	fmt.Printf("\n")
}
