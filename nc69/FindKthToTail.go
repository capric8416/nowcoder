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
 * @param pHead ListNode类
 * @param k int整型
 * @return ListNode类
 */
func FindKthToTail(pHead *ListNode, k int) *ListNode {
	// write code here

	slow, fast := pHead, pHead
	for i := k; i > 0; i-- {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}

	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}

	return slow
}

func main() {
	seven := &ListNode{Val: 7}
	six := &ListNode{Val: 6, Next: seven}
	five := &ListNode{Val: 5, Next: six}
	four := &ListNode{Val: 4, Next: five}
	three := &ListNode{Val: 3, Next: four}
	two := &ListNode{Val: 2, Next: three}
	one := &ListNode{Val: 1, Next: two}

	for i := 1; i < 10; i++ {
		result := FindKthToTail(one, i)
		if result != nil {
			fmt.Println(result.Val)
		} else {
			fmt.Println("nil")
		}
	}
}
