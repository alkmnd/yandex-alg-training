package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func minGroups(a []int, k int) int {
	sort.Ints(a)

	days := 0

	l, r := 0, 0
	curGr := 1
	for r < len(a) {
		curGr = curGr - 1
		for r < len(a) && a[r]-a[l] <= k {
			curGr++
			r++

		}
		if curGr > days {
			days = curGr
		}
		l += 1

	}

	return days

}

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

	result := minGroups(arr, k)
	fmt.Println(result)
}
