package main

import "fmt"

func main() {
	var blueTops, redTops, blueSocks, redSocks int
	fmt.Scan(&blueTops, &redTops, &blueSocks, &redSocks)
	if blueTops == 0 {
		if blueSocks == 0 {
			fmt.Print(1, 1)
		} else {
			fmt.Print(1, blueSocks+1)
		}
		return
	}
	if redTops == 0 {
		if redSocks == 0 {
			fmt.Print(1, 1)
		} else {
			fmt.Print(1, redSocks+1)
		}
		return
	}
	if blueSocks == 0 {
		fmt.Print(blueTops+1, 1)
		return
	}
	if redSocks == 0 {
		fmt.Print(redTops+1, 1)
		return
	}
	m1, n1 := max(blueTops, redTops)+1, 1
	m2, n2 := 1, max(blueSocks, redSocks)+1
	m3, n3 := blueTops+1, blueSocks+1
	m4, n4 := redTops+1, redSocks+1

	minM, minN := m1, n1

	updateMin := func(m, n int) {
		if m+n < minM+minN || (m+n == minM+minN && (m < minM || (m == minM && n < minN))) {
			minM, minN = m, n
		}
	}
	updateMin(m3, n3)
	updateMin(m4, n4)
	updateMin(m1, n1)
	updateMin(m2, n2)

	fmt.Println(minM, minN)

}
