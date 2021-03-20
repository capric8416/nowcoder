package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param head ListNode类
 * @param m int整型
 * @param n int整型
 * @return ListNode类
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// write code here

	if head == nil || head.Next == nil || m == n {
		return head
	}

	curr := head
	var i int
	var prevM, currM *ListNode = nil, head
	for i = 0; i < m-1; i, curr = i+1, curr.Next {
		if i == m-2 {
			prevM = curr
			currM = curr.Next
		}
	}

	var prev *ListNode
	for curr = currM; i < n; i++ {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	if currM != nil {
		currM.Next = curr
	}
	if prevM != nil {
		prevM.Next = prev
		return head
	} else {
		return prev
	}
}

func main() {
	five := &ListNode{Val: 5}
	four := &ListNode{Val: 4, Next: five}
	three := &ListNode{Val: 3, Next: four}
	two := &ListNode{Val: 2, Next: three}
	one := &ListNode{Val: 1, Next: two}
	result := reverseBetween(one, 2, 4)
	for curr := result; curr != nil; curr = curr.Next {
		fmt.Printf("-> %d ", curr.Val)
	}
	fmt.Println("-> nil")

	five = &ListNode{Val: 5}
	three = &ListNode{Val: 3, Next: five}
	result = reverseBetween(three, 1, 2)
	for curr := result; curr != nil; curr = curr.Next {
		fmt.Printf("-> %d ", curr.Val)
	}
	fmt.Println("-> nil")

	three = &ListNode{Val: 3}
	two = &ListNode{Val: 2, Next: three}
	one = &ListNode{Val: 1, Next: two}
	result = reverseBetween(one, 1, 2)
	for curr := result; curr != nil; curr = curr.Next {
		fmt.Printf("-> %d ", curr.Val)
	}
	fmt.Println("-> nil")
}
