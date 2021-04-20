package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

/**
 *
 * @param n int整型
 * @param m int整型
 * @return int整型
 */

/*
反推过程：(当前index + m) % 上一轮剩余数字的个数
最终剩下一个人时的安全位置肯定为1，反推安全位置在人数为n时的编号
人数为1： 0+1
人数为2： (0+m) % 2 + 1
人数为3： ((0+m) % 2 + m) % 3 + 1
人数为4： (((0+m) % 2 + m) % 3 + m) % 4 + 1
*/
func ysf(n int, m int) int {
	i, p := 0, 0
	for i = 2; i <= n; i++ {
		p = (p + m) % i
	}
	return p + 1
}

func ysf2(n int, m int) int {
	if n < 1 || m < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return (ysf(n-1, m)+m-1)%n + 1
}

// func ysf(n int, m int) int {
// 	// write code here

// 	i := 1
// 	head := &Node{value: i}
// 	curr := head
// 	for i, curr = 2, head; i <= n; i++ {
// 		n := &Node{value: i}
// 		curr.next = n
// 		curr = n
// 	}
// 	curr.next = head

// 	prev := curr
// 	for i, curr = 1, head; curr.next != curr; i, curr = i+1, curr.next {
// 		if i == m {
// 			i = 0
// 			prev.next = curr.next
// 		} else {
// 			prev = prev.next
// 		}
// 	}

// 	return curr.value
// }

func main() {
	fmt.Println(ysf2(5, 2))
	fmt.Println(ysf2(43, 9001))
}
