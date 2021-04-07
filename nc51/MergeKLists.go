package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MinHeap struct {
	length int
	data   []*ListNode
}

func (this *MinHeap) Init(a []*ListNode) {
	this.data = a
	this.length = len(a)
	// 从中间位置开始下沉
	for i := (this.length - 1) / 2; i >= 0; i-- {
		this.Sink(i)
	}
}

func (this *MinHeap) Push(n *ListNode) {
	// 添加到末尾并上浮, 长度加1
	this.data = append(this.data, n)
	this.length++
	this.Swin(this.length - 1)
}

func (this *MinHeap) Pop() *ListNode {
	n := this.data[0]

	// 长度减1, 删除首个元素, 将最后一个元素放在首位并下沉
	this.length--
	this.data[0] = this.data[this.length]
	this.data = this.data[:this.length]
	this.Sink(0)

	return n
}

func (this *MinHeap) Sink(parent int) {
	for child := 2*parent + 1; child < this.length; child = 2*child + 1 {
		// 选取较小的子节点
		if child+1 < this.length && this.data[child].Val > this.data[child+1].Val {
			child++
		}

		// 如果父节点小于等于子节点
		if this.data[parent].Val <= this.data[child].Val {
			break
		}

		this.data[parent], this.data[child] = this.data[child], this.data[parent]

		parent = child
	}
}

func (this *MinHeap) Swin(child int) {
	for child > 0 {
		// 如果父节点大于子节点
		parent := (child - 1) / 2
		if this.data[parent].Val > this.data[child].Val {
			this.data[parent], this.data[child] = this.data[child], this.data[parent]
		}
		child = parent
	}
}

/**
 *
 * @param lists ListNode类一维数组
 * @return ListNode类
 */
func mergeKLists(lists []*ListNode) *ListNode {
	// write code here

	temp := make([]*ListNode, 0)
	for _, node := range lists {
		if node != nil {
			temp = append(temp, node)
		}
	}

	length := len(temp)
	if length == 0 {
		return nil
	}

	var heap MinHeap
	heap.Init(temp)

	result := &ListNode{}
	current := result
	for heap.length > 0 {
		node := heap.Pop()
		if node.Next != nil {
			heap.Push(node.Next)
		}
		current, current.Next = node, node
	}

	return result.Next
}

func CreateList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	curr := head
	for _, n := range nums[1:] {
		curr.Next = &ListNode{Val: n}
		curr = curr.Next
	}

	return head
}

func main() {
	first := CreateList([]int{1, 3, 6, 8, 9})
	second := CreateList([]int{0, 2, 4, 5, 10})
	third := CreateList([]int{3, 5, 7, 11})
	fourth := CreateList([]int{4, 6, 12, 13})

	result := mergeKLists([]*ListNode{nil, first, second, third, fourth})
	for curr := result; curr != nil; curr = curr.Next {
		fmt.Printf("-> %d ", curr.Val)
	}
	fmt.Println("")

	first = CreateList([]int{-10, -9, -9, -3, -1, -1, 0})
	second = CreateList([]int{-5})
	third = CreateList([]int{4})
	fourth = CreateList([]int{-8})
	fifth := CreateList([]int{-9, -6, -5, -4, -2, 2, 3})
	sixth := CreateList([]int{-3, -3, -2, -1, 0})

	result = mergeKLists([]*ListNode{first, second, third, fourth, nil, fifth, sixth})
	for curr := result; curr != nil; curr = curr.Next {
		fmt.Printf("-> %d ", curr.Val)
	}
	fmt.Println("")

}
