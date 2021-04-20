package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param head ListNode类
 * @return ListNode类
 */
func oddEvenListByValue(head *ListNode) *ListNode {
	// write code here

	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	oddFirst := true
	var curr, odd, even *ListNode
	for curr = head; odd == nil || even == nil; curr = curr.Next {
		if curr.Val%2 != 0 {
			if odd == nil {
				odd = curr
				for odd.Next.Val%2 != 0 {
					odd = odd.Next
				}
				curr = odd
			}
		} else {
			if even == nil {
				even = curr
				for even.Next.Val%2 == 0 {
					even = even.Next
				}
				curr = even

				if odd == nil {
					oddFirst = false
				}
			}
		}
	}

	if oddFirst {
		for curr != nil {
			if curr.Val%2 == 0 {
				even = curr
				curr = curr.Next
			} else {
				next1, next2 := odd.Next, curr.Next
				odd.Next, curr.Next, even.Next = curr, next1, next2
				odd = curr
				curr = next2
			}
		}
	} else {
		for curr != nil {
			if curr.Val%2 != 0 {
				odd = curr
				curr = curr.Next
			} else {
				next1, next2 := even.Next, curr.Next
				even.Next, curr.Next, odd.Next = curr, next1, next2
				even = curr
				curr = next2
			}
		}
	}

	return head
}

func main123() {
	twentytwo := &ListNode{Val: 22}
	twentyone := &ListNode{Val: 21, Next: twentytwo}
	twenty := &ListNode{Val: 20, Next: twentyone}
	nineteen := &ListNode{Val: 19, Next: twenty}
	seventeen := &ListNode{Val: 17, Next: nineteen}
	fifthteen := &ListNode{Val: 15, Next: seventeen}
	fourteen := &ListNode{Val: 14, Next: fifthteen}
	twelve := &ListNode{Val: 12, Next: fourteen}
	ten := &ListNode{Val: 10, Next: twelve}
	nine := &ListNode{Val: 9, Next: ten}
	eight := &ListNode{Val: 8, Next: nine}
	seven := &ListNode{Val: 7, Next: eight}
	six := &ListNode{Val: 6, Next: seven}
	five := &ListNode{Val: 5, Next: six}
	four := &ListNode{Val: 4, Next: five}
	three := &ListNode{Val: 3, Next: four}
	two := &ListNode{Val: 2, Next: three}
	one := &ListNode{Val: 1, Next: two}
	zero4 := &ListNode{Val: 0, Next: one}
	zero3 := &ListNode{Val: 0, Next: zero4}
	zero2 := &ListNode{Val: 0, Next: zero3}
	zero1 := &ListNode{Val: 0, Next: zero2}

	for curr := zero1; curr != nil; curr = curr.Next {
		fmt.Printf(" -> %d", curr.Val)
	}
	fmt.Println()

	result := oddEvenListByValue(zero1)
	for curr := result; curr != nil; curr = curr.Next {
		fmt.Printf(" -> %d", curr.Val)
	}
	fmt.Println()

	seven = &ListNode{Val: 7}
	four = &ListNode{Val: 4, Next: seven}
	six = &ListNode{Val: 6, Next: four}
	five = &ListNode{Val: 5, Next: six}
	three = &ListNode{Val: 3, Next: five}
	one = &ListNode{Val: 1, Next: three}
	two = &ListNode{Val: 2, Next: one}

	for curr := two; curr != nil; curr = curr.Next {
		fmt.Printf(" -> %d", curr.Val)
	}
	fmt.Println()

	result = oddEvenListByValue(two)
	for curr := result; curr != nil; curr = curr.Next {
		fmt.Printf(" -> %d", curr.Val)
	}
	fmt.Println()
}
