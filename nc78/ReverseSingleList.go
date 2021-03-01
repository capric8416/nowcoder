package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 反转单链表，双指针，时间复杂度为O(n)，空间复杂度为O(1)
 * @param pHead ListNode类
 * @return ListNode类
 */
func ReverseList(pHead *ListNode) *ListNode {
	// write code here

	// 空链直接返回
	if pHead == nil {
		return pHead
	}

	// 初始化prev为空，curr为头节点
	var prev, curr *ListNode = nil, pHead

	// 遍历
	for curr != nil {
		// 临时保存原来的后驱指针(因为接下来要修改当前节点的后驱指针，以免丢失后面的节点)
		next := curr.Next
		// 当前节点后缀指针指向前节点
		curr.Next = prev
		// 当前节点下一次将变成前节点
		prev = curr
		// 移动到下一个节点(之前临时保存的后驱指针)
		curr = next
	}

	return prev
}

// 创建[start, stop)范围内，步长为step的int升序单链表
func CreateList(start int, stop int, step int) *ListNode {
	// 如果步长小于0或者stop小于start返回空链表
	if step < 0 || stop < start {
		return nil
	}

	// 创建头节点
	head := ListNode{
		Val:  start,
		Next: nil,
	}

	// 下一次前节点
	var prev *ListNode = &head

	// 创建后面的节点
	for i := start + 1; i < stop; i += step {
		curr := ListNode{
			Val:  i,
			Next: nil,
		}

		// 前节点后驱指向当前节点
		prev.Next = &curr

		// 下一次前节点
		prev = &curr
	}

	return &head
}

// 遍历并打印单链表
func PrintList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}

func main() {
	fmt.Println("==================")
	head := CreateList(1, 10, 1)
	PrintList(head)

	fmt.Println("==================")
	reversed := ReverseList(head)
	PrintList(reversed)
}
