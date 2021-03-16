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
func detectCycle(head *ListNode) *ListNode {
	// write code here

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}

	return nil
}

func main() {
	ten := &ListNode{Val: 10}
	nine := &ListNode{Val: 9, Next: ten}
	eight := &ListNode{Val: 8, Next: nine}
	seven := &ListNode{Val: 7, Next: eight}
	six := &ListNode{Val: 6, Next: seven}
	five := &ListNode{Val: 5, Next: six}
	four := &ListNode{Val: 4, Next: five}
	three := &ListNode{Val: 3, Next: four}
	two := &ListNode{Val: 2, Next: three}
	one := &ListNode{Val: 1, Next: two}
	zero := &ListNode{Val: 0, Next: one}

	for _, n := range []*ListNode{one, two, three, four, five, six, seven, eight, nine, nil} {
		ten.Next = n
		cycle := detectCycle(zero)
		if cycle != nil {
			fmt.Println(cycle.Val)
		} else {
			fmt.Println(cycle)
		}
	}
}
