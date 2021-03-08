package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param l1 ListNode类
 * @param l2 ListNode类
 * @return ListNode类
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// write code here

	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	dummy := ListNode{Val: 0, Next: nil}
	p1, p2, p3 := l1, l2, &dummy
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			p3.Next = p1
			p1 = p1.Next
		} else {
			p3.Next = p2
			p2 = p2.Next
		}
		p3 = p3.Next
	}

	if p1 != nil {
		p3.Next = p1
	}

	if p2 != nil {
		p3.Next = p2
	}

	return dummy.Next
}

func main() {
	zero := &ListNode{Val: 0}
	one := &ListNode{Val: 1}
	two := &ListNode{Val: 2}
	three := &ListNode{Val: 3}
	four := &ListNode{Val: 4}
	dummy := &ListNode{Val: 4}
	five := &ListNode{Val: 5}
	six := &ListNode{Val: 6}

	zero.Next = two
	two.Next = four
	four.Next = five

	one.Next = three
	three.Next = dummy
	dummy.Next = six

	result := mergeTwoLists(zero, one)
	for p := result; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}
}
