package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	stack := make([]int32, 0)
	for _, v := range str {
		if v == '\n' {
			break
		}
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		} else {
			switch stack[len(stack)-1] {
			case '{':
				if v == '}' {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, v)
				}
			case '(':
				if v == ')' {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, v)
				}

			case '[':
				if v == ']' {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, v)
				}
			default:
				stack = append(stack, v)
			}
		}
	}

	if len(stack) == 0 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
