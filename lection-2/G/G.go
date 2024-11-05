package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, c int
	fmt.Scan(&n, &c)

	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')

	l, r := 0, 0

	cntA := 0
	cntB := 0
	curR := 0

	res := 0

	for r < n {
		if s[r] == 'a' {
			cntA++
		} else if s[r] == 'b' {
			cntB++
			curR += cntA
		}
		if curR <= c {
			if res < r-l+1 {
				res = r - l + 1
			}
		}
		for curR > c {
			if s[l] == 'a' {
				cntA--
				curR -= cntB
			} else if s[l] == 'b' {
				cntB--
			}
			l++
		}
		r++
	}

	fmt.Println(res)
}
