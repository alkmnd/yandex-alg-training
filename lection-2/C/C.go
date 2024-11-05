package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	nr := strings.Fields(line)
	n, _ := strconv.Atoi(nr[0])
	k, _ := strconv.Atoi(nr[1])
	line, _ = reader.ReadString('\n')
	nums := strings.Fields(line)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(nums[i])
	}

	res := 0
	l, r := 0, 1
	for r < n {
		if arr[r]-arr[l] > k {
			res += n - r
			l++
			continue
		}
		r++
	}

	fmt.Println(res)
	//l, r := 0, 1
	//for r < n {
	//	for arr[r]-arr[l] > k {
	//		res += n - r
	//		l++
	//	}
	//	r++
	//}

	fmt.Println(res)
}
