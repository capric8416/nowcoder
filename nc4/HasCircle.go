package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param head ListNode类
 * @return bool布尔型
 */
func hasCycle(head *ListNode) bool {
	// write code here

	// 方法1 - 遍历，放到set里面，如果当前节点已存在set里面，则说明有环，空间复杂度O(N)

	// 方法2 - 遍历，每次删除头结点，使用head.next = head
	// 如果head或者head.next为空，表示无环
	// 如果head = head.next说明有环
	// 递归调用
	// 破坏了原来的链表

	// 方法3 - 快慢指针
	// 快指针每次前进2步
	// 慢指针每次前进1步
	// 如果相遇，说明有环

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}
