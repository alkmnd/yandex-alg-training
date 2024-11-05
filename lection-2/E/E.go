package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumOfTripletProducts(arr []int) int {
	mod := 1000000007
	totalSum := 0

	n := len(arr)
	prefixSum := (arr[n-1] + arr[n-2]) % mod
	prefixProduct := (arr[n-1] * arr[n-2]) % mod
	totalSum += (arr[n-3] * arr[n-1] * arr[n-2]) % mod

	for i := n - 4; i >= 0; i-- {
		prefixProduct = (prefixProduct + (arr[i+1] * prefixSum)) % mod
		prefixSum = (prefixSum + arr[i+1]) % mod
		totalSum = (totalSum + prefixProduct*arr[i]) % mod
	}

	return totalSum % mod
}

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

	//mod := 1000000007
	result := sumOfTripletProducts(arr)
	fmt.Println(result)

}
