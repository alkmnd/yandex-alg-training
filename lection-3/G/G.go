package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, b int
	fmt.Scan(&n, &b)

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	parts := strings.Split(line, " ")

	arr := make([]int, len(parts))
	for i := 0; i < len(parts); i++ {
		arr[i], _ = strconv.Atoi(parts[i])
	}

	queue := 0

	var res int

	for i := 0; i < len(arr); i++ {
		queue += arr[i]
		res += queue
		if queue < b {
			queue = 0
		} else {
			queue = queue - b
		}
	}

	res += queue

	fmt.Println(res)
}
