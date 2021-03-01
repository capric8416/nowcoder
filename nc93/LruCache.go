package main

import "fmt"

type DoubleListNode struct {
	Key   int
	Value int
	Next  *DoubleListNode
	Prev  *DoubleListNode
}

/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */
func LRU(operators [][]int, k int) []int {
	// write code here

	// 双链表大小
	listLength := 0
	// 双链表头节点
	var listHead *DoubleListNode
	var listTail *DoubleListNode

	// 存储key - value(list node pointer)
	kv := make(map[int]*DoubleListNode)

	// 存储get结果
	results := make([]int, 0)

	for _, op := range operators {
		key := op[1]
		curr := kv[key]

		if op[0] == 2 { // get
			result := -1

			// 如果value存在，更新cache顺序
			if curr != nil {
				result = curr.Value

				// 如果不是表头
				if curr != listHead {
					// 当前节点的前节点的尾指针指向当前节点的后节点
					curr.Prev.Next = curr.Next
					if curr == listTail {
						// 如果当前节点是尾节点, 则前节点成为新的尾节点
						listTail = curr.Prev
					} else {
						// 当前节点的后节点的头指针指向当前节点的前节点
						curr.Next.Prev = curr.Prev
					}
					// 放到头部
					listHead.Prev = curr
					curr.Next = listHead
					listHead = curr
				}
			}

			results = append(results, result)

		} else { // set
			value := op[2]

			if curr != nil { // 如果value存在
				// 更新值
				if curr.Value != value {
					curr.Value = value
				}
				// 如果不是表头
				if curr != listHead {
					// 当前节点的前节点的尾指针指向当前节点的后节点
					curr.Prev.Next = curr.Next
					if curr == listTail {
						// 如果当前节点是尾节点, 则前节点成为新的尾节点
						listTail = curr.Prev
					} else {
						// 当前节点的后节点的头指针指向当前节点的前节点
						curr.Next.Prev = curr.Prev
					}
					// 当前节点成为新的头节点
					curr.Next = listHead
					listHead.Prev = curr
					listHead = curr
				}
			} else { // 如果value不存在
				// 长度加1
				listLength++

				// 创建节点
				curr := &DoubleListNode{
					Key:   key,
					Value: value,
					Next:  nil,
					Prev:  nil,
				}

				if listHead == nil {
					// 如果链表为空, 当前节点既是头节点又是尾节点
					listHead, listTail = curr, curr
				} else {
					// 当前节点成为新的头节点
					curr.Next = listHead
					listHead.Prev = curr
					listHead = curr
				}

				// 保存指针
				kv[key] = curr

				// 检查长度
				if listLength > k {
					// 删除尾节点
					kv[listTail.Key] = nil
					listTail.Prev.Next = nil
					listTail = listTail.Prev
					// 长度减1
					listLength--
				}
			}
		}
	}

	return results
}

func main() {
	operators := [][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2}}
	results := LRU(operators, 3)
	fmt.Println(results)
}
