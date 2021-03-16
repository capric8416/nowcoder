package main

import "fmt"

/**
 *
 * @param s string字符串
 * @return bool布尔型
 */
func isValid(s string) bool {
	// write code here

	// stack := make([]rune, 0)
	// for _, c := range s {
	// 	if c == '(' || c == '[' || c == '{' {
	// 		stack = append(stack, c)
	// 	} else {
	// 		if len(stack) == 0 {
	// 			return false
	// 		}

	// 		top := stack[len(stack)-1]
	// 		if (top == '(' && c == ')') || (top == '[' && c == ']') || (top == '{' && c == '}') {
	// 			stack = stack[0 : len(stack)-1]
	// 		} else {
	// 			return false
	// 		}
	// 	}
	// }

	// return len(stack) == 0

	stack := make([]rune, 0)
	for _, c := range s {
		switch c {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		default:
			if len(stack) == 0 || c != stack[len(stack)-1] {
				return false
			}

			stack = stack[0 : len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	for _, s := range []string{"(", "[]", "(]", "()[]{}", "([)]", "[()]", "{[()]}", "{[([{([{()}])}])]}"} {
		fmt.Println(s, "->", isValid(s))
	}
}
