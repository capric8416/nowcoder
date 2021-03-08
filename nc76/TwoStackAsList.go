package main

import "fmt"

var stack1 []int
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	length := len(stack2)
	if length != 0 {
		node := stack2[length-1]
		stack2 = stack2[0 : length-1]
		return node
	}

	for i := len(stack1) - 1; i > 0; i-- {
		stack2 = append(stack2, stack1[i])
	}
	node := stack1[0]
	stack1 = make([]int, 0)
	return node
}

func main() {
	Push(0)
	Push(1)
	Push(2)
	Push(3)
	Push(4)
	fmt.Println(Pop())
	fmt.Println(Pop())
	fmt.Println(Pop())
	fmt.Println(Pop())
	fmt.Println(Pop())
	Push(5)
	Push(6)
	fmt.Println(Pop())
	Push(7)
	fmt.Println(Pop())
	fmt.Println(Pop())
}
