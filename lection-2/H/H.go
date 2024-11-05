package main

import (
	"bufio"
	"os"
	"sort"
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
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

}
