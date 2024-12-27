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
	line, _ = reader.ReadString('\n')
	nums := strings.Fields(line)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(nums[i])
	}

	prefSums := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		prefSums[i] = prefSums[i-1] + arr[i-1]
	}

	sufSums := make([]int, n+1)
	for i := n - 1; i > -1; i-- {
		sufSums[i] = sufSums[i+1] + arr[i]
	}

	ans := 0

	for i := 1; i < n; i++ {
		ans += arr[i] * i
	}

	nowSum := ans
	for j := 1; j < n; j++ {
		nowSum += prefSums[j] - sufSums[j]
		ans = min(ans, nowSum)
	}

	fmt.Println(ans)

}
