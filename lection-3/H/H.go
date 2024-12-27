package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	var stackSum []int
	reader := bufio.NewReader(os.Stdin)
	var op string
	var num int
	for i := 0; i < n; i++ {
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		if str == "-" {
			op = "-"
		} else {
			op = string(str[0])
			numStr := str[1:]
			num, _ = strconv.Atoi(numStr)
		}
		if op == "+" {
			if len(stackSum) == 0 {
				stackSum = append(stackSum, num)
			} else {
				stackSum = append(stackSum, num+stackSum[len(stackSum)-1])
			}
		}
		if op == "-" {
			if len(stackSum) == 1 {
				fmt.Println(stackSum[len(stackSum)-1])
			} else {
				fmt.Println(stackSum[len(stackSum)-1] - stackSum[len(stackSum)-2])
			}
			stackSum = stackSum[:len(stackSum)-1]
		}
		if op == "?" {
			if len(stackSum)-num == 0 {
				fmt.Println(stackSum[len(stackSum)-1])
			} else {
				fmt.Println(stackSum[len(stackSum)-1] - stackSum[len(stackSum)-num-1])
			}
		}

	}

}
