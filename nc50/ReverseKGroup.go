package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
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
 * @param k int整型
 * @return ListNode类
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	// write code here

	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Val: 0, Next: head}
	pre, end := dummy, dummy

	for end.Next != nil {
		// 寻找待反转区间的尾节点
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}

		if end == nil {
			break
		}

		// 保存链表断开处的指针
		next := end.Next
		// 断开链表
		end.Next = nil

		// 待反转区间头节点
		start := pre.Next
		// 反转
		pre.Next = reverse(start)
		// 反转后头节点变成尾节点，续上断开处的链表
		start.Next = next

		// 下次待反转区间的头节点的前节点
		pre, end = start, start
	}

	return dummy.Next
}

func main() {
	for k := 1; k <= 10; k++ {
		zero := &ListNode{Val: 0}
		one := &ListNode{Val: 1}
		two := &ListNode{Val: 2}
		three := &ListNode{Val: 3}
		four := &ListNode{Val: 4}
		five := &ListNode{Val: 5}
		six := &ListNode{Val: 6}
		seven := &ListNode{Val: 7}
		eight := &ListNode{Val: 8}
		nine := &ListNode{Val: 9}

		zero.Next = one
		one.Next = two
		two.Next = three
		three.Next = four
		four.Next = five
		five.Next = six
		six.Next = seven
		seven.Next = eight
		eight.Next = nine

		for curr := zero; curr != nil; curr = curr.Next {
			fmt.Printf("%d ", curr.Val)
		}

		fmt.Printf(" --%d->  ", k)

		reversed := reverseKGroup(zero, k)
		for curr := reversed; curr != nil; curr = curr.Next {
			fmt.Printf("%d ", curr.Val)
		}

		fmt.Println("")
	}
}
