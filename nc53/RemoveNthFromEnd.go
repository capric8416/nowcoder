package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param head ListNode类
 * @param n int整型
 * @return ListNode类
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// write code here

	if head == nil || (head.Next == nil && n == 1) {
		return nil
	}

	// fast := head
	// for ; fast != nil && n > 0; n-- {
	// 	fast = fast.Next
	// }
	// if fast == nil {
	// 	return head.Next
	// }

	// prev, slow := head, head
	// for ; fast != nil; slow, fast = slow.Next, fast.Next {
	// 	if slow != nil {
	// 		prev = slow
	// 	}
	// }
	// prev.Next = slow.Next

	dummy := &ListNode{Val: 0, Next: head}
	head = dummy

	fast := head
	for ; n > 0 && fast != nil; n-- {
		fast = fast.Next
	}

	slow := head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}

func output(head *ListNode) {
	if head == nil {
		fmt.Printf("nil\n")
		return
	}

	for curr := head; curr != nil; curr = curr.Next {
		fmt.Printf("%d", curr.Val)
		if curr.Next != nil {
			fmt.Printf(" -> ")
		} else {
			fmt.Printf("\n")
		}
	}
}

func main() {
	// zero node
	output(nil)
	result := removeNthFromEnd(nil, 1)
	output(result)

	fmt.Println("")

	// only one node
	zero := &ListNode{Val: 0}

	output(zero)
	result = removeNthFromEnd(zero, 1)
	output(result)

	fmt.Println("")

	// remove first node
	two := &ListNode{Val: 2}
	one := &ListNode{Val: 1, Next: two}
	zero = &ListNode{Val: 0, Next: one}

	output(zero)
	result = removeNthFromEnd(zero, 3)
	output(result)

	fmt.Println("")

	// remove last node
	two = &ListNode{Val: 2}
	one = &ListNode{Val: 1, Next: two}
	zero = &ListNode{Val: 0, Next: one}

	output(zero)
	result = removeNthFromEnd(zero, 1)
	output(result)

	fmt.Println("")

	// remove middle node
	ten := &ListNode{Val: 10}
	nine := &ListNode{Val: 9, Next: ten}
	eight := &ListNode{Val: 8, Next: nine}
	seven := &ListNode{Val: 7, Next: eight}
	six := &ListNode{Val: 6, Next: seven}
	five := &ListNode{Val: 5, Next: six}
	four := &ListNode{Val: 4, Next: five}
	three := &ListNode{Val: 3, Next: four}
	two = &ListNode{Val: 2, Next: three}
	one = &ListNode{Val: 1, Next: two}
	zero = &ListNode{Val: 0, Next: one}

	output(zero)
	result = removeNthFromEnd(zero, 5)
	output(result)
}
