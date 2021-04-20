package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param head ListNode类
 * @param x int整型
 * @return ListNode类
 */
func partition(head *ListNode, x int) *ListNode {
	// write code here

	if head == nil || head.Next == nil {
		return head
	}

	low, high := &ListNode{Val: 0}, &ListNode{Val: 0}

	p, q := low, high
	for curr := head; curr != nil; curr = curr.Next {
		if curr.Val < x {
			p.Next = curr
			p = curr
		} else {
			q.Next = curr
			q = curr
		}
	}

	p.Next = high.Next
	q.Next = nil

	return low.Next
}

// func partition(head *ListNode, x int) *ListNode {
// 	// write code here

// 	if head == nil || head.Next == nil {
// 		return head
// 	}

// 	h, t := &ListNode{Val: 0, Next: head}, &ListNode{Val: 0, Next: head}

// 	var low, high *ListNode
// 	for low, high = h, head; high.Next != nil; high = high.Next {
// 	}

// 	high.Next = t
// 	high = t

// 	for curr := h.Next; curr != t; {
// 		if curr.Val < x {
// 			low = curr
// 			curr = curr.Next
// 		} else {
// 			low.Next = curr.Next
// 			high.Next = curr
// 			high = curr
// 			high.Next = nil
// 			curr = low.Next
// 		}
// 	}
// 	low.Next = t.Next

// 	return h.Next
// }

func main() {
	two2 := &ListNode{Val: 2}
	five := &ListNode{Val: 5, Next: two2}
	two1 := &ListNode{Val: 2, Next: five}
	three := &ListNode{Val: 3, Next: two1}
	four := &ListNode{Val: 4, Next: three}
	one := &ListNode{Val: 1, Next: four}
	six := &ListNode{Val: 6, Next: one}

	for curr := six; curr != nil; curr = curr.Next {
		fmt.Printf(" -> %d", curr.Val)
	}
	fmt.Println()

	result := partition(six, 3)
	for curr := result; curr != nil; curr = curr.Next {
		fmt.Printf(" -> %d", curr.Val)
	}
	fmt.Println()

	four3 := &ListNode{Val: 4}
	three3 := &ListNode{Val: 3, Next: four3}
	two3 := &ListNode{Val: 2, Next: three3}
	one3 := &ListNode{Val: 1, Next: two3}

	for curr := one3; curr != nil; curr = curr.Next {
		fmt.Printf(" -> %d", curr.Val)
	}
	fmt.Println()

	result = partition(one3, 3)
	for curr := result; curr != nil; curr = curr.Next {
		fmt.Printf(" -> %d", curr.Val)
	}
	fmt.Println()
}
