package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	arr := make([]int, n)
	d := make([]int, 0, k+1)
	ans := make([]int, 0, len(arr)-k+1)

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(parts[i])
		arr[i] = num
	}
	for i := 0; i < len(arr); i++ {
		if i > k-1 && d[0] <= i-k {
			d = d[1:]
		}
		for len(d) > 0 && arr[i] < arr[d[len(d)-1]] {
			d = d[:len(d)-1]
		}
		d = append(d, i)
		if i >= k-1 {
			ans = append(ans, arr[d[0]])
		}
	}
	for _, v := range ans {
		fmt.Println(v)
	}
}
