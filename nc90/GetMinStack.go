package main

import "fmt"

type StackNode struct {
	value int
	next  *StackNode
}

type Stack struct {
	top *StackNode
}

func New() *Stack {
	return &Stack{top: nil}
}

func (this *Stack) Push(n int) {
	current := &StackNode{value: n, next: this.top}
	this.top = current
}

func (this *Stack) Pop() {
	this.top = this.top.next
}

func (this *Stack) Top() int {
	return this.top.value
}

func (this *Stack) Empty() bool {
	return this.top == nil
}

/**
 * return a array which include all ans for op3
 * @param op int整型二维数组 operator
 * @return int整型一维数组
 */
func getMinStack(op [][]int) []int {
	// write code here

	dataStack := New()
	minStack := New()

	result := make([]int, 0)

	for _, item := range op {
		switch item[0] {
		case 1:
			{
				v := item[1]
				dataStack.Push(v)
				if minStack.Empty() || minStack.Top() > v {
					minStack.Push(v)
				}
			}
		case 2:
			{
				v := dataStack.Top()
				dataStack.Pop()
				if v == minStack.Top() {
					minStack.Pop()
				}
			}
		case 3:
			{
				result = append(result, minStack.Top())
			}
		}
	}

	return result
}

func main() {
	fmt.Println(getMinStack([][]int{{1, 3}, {1, 2}, {1, 1}, {3}, {2}, {3}}))
}
