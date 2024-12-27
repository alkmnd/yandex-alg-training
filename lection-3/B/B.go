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

	arr := make([]int, n)

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(parts[i])
		arr[i] = num
	}

	res := make([]int, n)

	for i := range res {
		res[i] = -1
	}

	type elem struct {
		ind   int
		value int
	}

	elems := make([]elem, 0, n)

	for i, v := range arr {
		if i != 0 {
			last := elems[len(elems)-1]
			for last.value > v {
				res[last.ind] = i
				elems = elems[:len(elems)-1]
				if len(elems) == 0 {
					break
				} else {
					last = elems[len(elems)-1]
				}
			}
		}
		elems = append(elems, elem{ind: i, value: v})

	}

	for i := range res {
		fmt.Print(strconv.Itoa(res[i]) + " ")
	}

}
