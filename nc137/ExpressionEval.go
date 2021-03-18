package main

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type Stack struct {
	Top    *Node
	Length int
}

func New() *Stack {
	return &Stack{Top: &Node{Value: nil, Next: nil}, Length: 0}
}

func (this *Stack) Peek() interface{} {
	return this.Top.Next.Value
}

func (this *Stack) Push(Value interface{}) {
	this.Length++
	node := &Node{Value: Value, Next: this.Top.Next}
	this.Top.Next = node
}

func (this *Stack) Pop() interface{} {
	this.Length--
	node := this.Top.Next
	this.Top.Next = node.Next
	return node.Value
}

func (this *Stack) IsEmpty() bool {
	return this.Length == 0
}

type PriorityType byte

const (
	ER PriorityType = 0 // error
	LT PriorityType = 1 // less than
	EQ PriorityType = 2 // equal to
	GT PriorityType = 3 // greater than
)

func (t PriorityType) String() string {
	switch t {
	case ER:
		return "ER"
	case LT:
		return "LT"
	case EQ:
		return "EQ"
	case GT:
		return "GT"
	default:
		return "NA"
	}
}

const DUMMY byte = '\''

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 返回表达式的值
 * @param s string字符串 待计算的表达式
 * @return int整型
 */
func solve(s string) int {
	// write code here

	if len(s) == 0 {
		return 0
	}

	/**
	 * ASCII
	 * 39 = ' -> 0
	 * 40 = ( -> 1
	 * 41 = ) -> 2
	 * 42 = * -> 3
	 * 43 = + -> 4
	 * 44 = , -> 5
	 * 45 = - -> 6
	 * 46 = . -> 7
	 * 47 = / -> 8
	 */

	priority := [][]PriorityType{
		//least           invalid invalid
		// ' (   )   *   +   ,   -   .   /
		{EQ, LT, LT, LT, LT, ER, LT, ER, LT}, // ' - least
		{ER, LT, EQ, LT, LT, ER, LT, ER, LT}, // (
		{GT, ER, GT, GT, GT, ER, GT, ER, GT}, // )
		{GT, LT, GT, GT, GT, ER, GT, ER, GT}, // *
		{GT, LT, GT, LT, GT, ER, GT, ER, LT}, // +
		{ER, ER, ER, ER, ER, ER, ER, ER, ER}, // , - invalid
		{GT, LT, GT, LT, GT, ER, GT, ER, LT}, // -
		{ER, ER, ER, ER, ER, ER, ER, ER, ER}, // . - invalid
		{GT, LT, GT, GT, GT, ER, GT, ER, GT}, // /
	}

	numbers := New()

	operators := New()
	operators.Push(DUMMY)

	length := len(s)
	var prev, curr byte
	nagative, digits := false, make([]int, 0)
	for i := 0; i <= length; i++ {
		if i != length {
			curr = s[i]
			// if curr != ' ' {
			// 	fmt.Printf("%c", curr)
			// }
		} else {
			curr = DUMMY
		}

		if curr == ' ' { // space
			continue
		} else if '0' <= curr && curr <= '9' { // digit
			prev = curr
			digits = append(digits, int(curr-'0'))
		} else { // operator
			// convert digits to number
			if curr == '-' && prev < '0' && prev != ')' {
				nagative = true
				continue
			}
			if len(digits) > 0 {
				n, m := 0, 1
				for j := len(digits) - 1; j >= 0; j-- {
					n += m * digits[j]
					m *= 10
				}
				if nagative {
					n = -n
				}
				numbers.Push(n)
				nagative, digits = false, make([]int, 0)
			}

			prev = curr

			// compare op priority
			for !operators.IsEmpty() {
				op := operators.Peek().(byte)
				p := priority[PriorityType(op-DUMMY)][PriorityType(curr-DUMMY)]
				if p == LT { // top op priority < current op priority
					operators.Push(curr)
					break
				} else if p == EQ { // top op priority = current op priority
					operators.Pop()
					break
				} else { // top op priority > current op priority
					operators.Pop()
					a := numbers.Pop().(int)
					b := numbers.Pop().(int)
					if op == '+' {
						numbers.Push(b + a)
					} else if op == '-' {
						numbers.Push(b - a)
					} else if op == '*' {
						numbers.Push(b * a)
					} else {
						numbers.Push(b / a)
					}
				}
			}
		}
	}

	return numbers.Pop().(int)
}

func main() {
	fmt.Printf(" = %d\n", solve("-1+(-2)+(-3)-4-5-6*2+(-4)*3-8/2"))
	fmt.Printf(" = %d\n", solve("1+2"))
	fmt.Printf(" = %d\n", solve("-2*3+1+(-2)"))
	fmt.Printf(" = %d\n", solve("1 - (3 + 2) * 7 - 9 / 3 + (1 + 2) * 3"))
	fmt.Printf(" = %d\n", solve("1+22+333+4444+55555"))
	fmt.Printf(" = %d\n", solve("(2*(3-4))*5"))
	fmt.Printf(" = %d\n", solve("33 + 22 * 3 * 4 - 1111"))
}
