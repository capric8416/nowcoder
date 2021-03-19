package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var length int
	var prev *ListNode
	for curr := head; curr != nil; length++ {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

/**
 *
 * @param head ListNode类 the head
 * @return bool布尔型
 */
func isPail(head *ListNode) bool {
	// write code here

	if head == nil {
		return false
	}

	if head.Next == nil {
		return true
	}

	if head.Next.Next == nil {
		return head.Val == head.Next.Val
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
	}

	tail := reverseList(slow.Next)
	for h, t := head, tail; t != nil; h, t = h.Next, t.Next {
		if h.Val != t.Val {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPail(nil))

	fmt.Println(isPail(&ListNode{Val: 1}))

	fmt.Println(isPail(&ListNode{Val: 1, Next: &ListNode{Val: 1}}))

	d := &ListNode{Val: 1}
	c := &ListNode{Val: 2, Next: d}
	b := &ListNode{Val: 2, Next: c}
	a := &ListNode{Val: 1, Next: b}
	fmt.Println(isPail(a))

	d = &ListNode{Val: 1}
	c = &ListNode{Val: 2, Next: d}
	e := &ListNode{Val: 3, Next: c}
	b = &ListNode{Val: 2, Next: e}
	a = &ListNode{Val: 1, Next: b}
	fmt.Println(isPail(a))
}
