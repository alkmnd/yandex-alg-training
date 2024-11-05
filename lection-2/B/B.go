package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	prefSumArr := make([]int, n+1)
	prefSumArr[0] = 0
	for i := 1; i < n+1; i++ {
		prefSumArr[i] = prefSumArr[i-1] + arr[i-1]
	}

	res := 0
	l, r := 0, 1
	for r < n+1 {
		if prefSumArr[r]-prefSumArr[l] == k {
			res++
			r++
		} else if prefSumArr[r]-prefSumArr[l] < k {
			r++
		} else {
			l++
		}
	}

	fmt.Println(res)

}
