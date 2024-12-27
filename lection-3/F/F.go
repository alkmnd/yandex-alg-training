package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	reader := bufio.NewReader(os.Stdin)
	w, _ := reader.ReadString('\n')
	w = strings.TrimSpace(w)

	expr, _ := reader.ReadString('\n')
	expr = strings.TrimSpace(expr)

	closing := map[rune]rune{')': '(', ']': '['}
	stack := make([]rune, 0)

	for _, v := range expr {
		if v == ')' || v == ']' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}

	res := make([]rune, 0)

	for i := 0; i < n-len(expr); i++ {
		for _, c := range w {
			if (c == ')' || c == ']') && len(stack) > 0 &&
				stack[len(stack)-1] == closing[c] {
				res = append(res, c)
				stack = stack[:len(stack)-1]
				break
			} else if !(c == ')' || c == ']') && n-len(expr)-i > len(stack) {
				res = append(res, c)
				stack = append(stack, c)
				break
			}
		}
	}

	fmt.Print(expr)
	for i := range res {
		fmt.Print(string(res[i]))
	}
}
