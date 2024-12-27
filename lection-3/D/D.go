package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInteger(s string) (int, bool) {
	v, err := strconv.Atoi(s)
	return v, err == nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	parts := strings.Split(line, " ")
	stack := make([]int, 0, len(parts))
	for _, v := range parts {
		if num, flag := isInteger(v); flag != false {
			stack = append(stack, num)
		} else {
			var val int
			n := len(stack)
			switch v {
			case "+":
				val = stack[n-2] + stack[n-1]
			case "-":
				val = stack[n-2] - stack[n-1]
			case "*":
				val = stack[n-2] * stack[n-1]
			}
			stack = stack[:n-2]
			stack = append(stack, val)
		}
	}

	fmt.Println(stack[0])
}
