package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
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
 * @param head ListNode类
 * @return void
 */
func reorderList(head *ListNode) {
	// write code here

	if head == nil {
		return
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	center := slow.Next
	slow.Next = nil

	center = reverseList(center)

	var first, second *ListNode
	for first, second = head, center; first != nil && second != nil; {
		a, b := first.Next, second.Next
		first.Next, second.Next = second, a
		first, second = a, b
	}
}

func main() {
	reorderList(nil)

	a := &ListNode{Val: 1}
	reorderList(a)
	for curr := a; curr != nil; curr = curr.Next {
		fmt.Printf("-> %d ", curr.Val)
	}
	fmt.Println("")

	d := &ListNode{Val: 2}
	c := &ListNode{Val: 1, Next: d}
	reorderList(c)
	for curr := c; curr != nil; curr = curr.Next {
		fmt.Printf("-> %d ", curr.Val)
	}
	fmt.Println("")

	four := &ListNode{Val: 4}
	three := &ListNode{Val: 3, Next: four}
	two := &ListNode{Val: 2, Next: three}
	one := &ListNode{Val: 1, Next: two}
	reorderList(one)
	for curr := one; curr != nil; curr = curr.Next {
		fmt.Printf("-> %d ", curr.Val)
	}
	fmt.Println("")

	nine := &ListNode{Val: 9}
	eight := &ListNode{Val: 8, Next: nine}
	seven := &ListNode{Val: 7, Next: eight}
	six := &ListNode{Val: 6, Next: seven}
	five := &ListNode{Val: 5, Next: six}
	reorderList(five)
	for curr := five; curr != nil; curr = curr.Next {
		fmt.Printf("-> %d", curr.Val)
	}
	fmt.Println("")

	one.Next = two
	two.Next = three
	three.Next = four
	four.Next = five
	five.Next = six
	six.Next = seven
	seven.Next = eight
	eight.Next = nil
	reorderList(one)
	for curr := one; curr != nil; curr = curr.Next {
		fmt.Printf("-> %d", curr.Val)
	}
	fmt.Println("")

	one.Next = two
	two.Next = three
	three.Next = four
	four.Next = five
	five.Next = six
	six.Next = seven
	seven.Next = eight
	eight.Next = nine
	nine.Next = nil
	reorderList(one)
	for curr := one; curr != nil; curr = curr.Next {
		fmt.Printf("-> %d", curr.Val)
	}
	fmt.Println("")
}
