package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Scan(&num)
	arr := make([]int, num)
	for i := 0; i < num; i++ {
		fmt.Scan(&arr[i])
	}

	prefSumArr := make([]int, num+1)
	prefSumArr[0] = 0
	for i := 1; i < num+1; i++ {
		prefSumArr[i] = prefSumArr[i-1] + arr[i-1]
	}

	for j := 1; j < num+1; j++ {
		fmt.Print(strconv.Itoa(prefSumArr[j]) + " ")
	}

}
