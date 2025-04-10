package main

import "fmt"

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(value int) {
	*s = append(*s, value)
}

func (s *Stack) pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		ele := (*s)[index]
		*s = (*s)[:index]
		return ele, true
	}
}


func main() {
	var stack Stack

	stack.push(1)
	stack.push(2)
	stack.push(4)
	stack.push(5)

	for len(stack) > 0 {
		x, y := stack.pop()
		if y == true {
			fmt.Println(x)
		}
	}

}