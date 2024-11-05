package main

import (
	"fmt"
)

func main() {
	var sw_x, sw_y int
	fmt.Scan(&sw_x, &sw_y)
	var ne_x, ne_y int
	fmt.Scan(&ne_x, &ne_y)
	var p_x, p_y int
	fmt.Scan(&p_x, &p_y)

	if p_x >= ne_x && p_y >= ne_y {
		fmt.Println("NE")
		return
	}
	if p_x >= ne_x && p_y <= sw_y {
		fmt.Println("SE")
		return
	}
	if p_x <= sw_x && p_y <= sw_y {
		fmt.Println("SW")
		return
	}
	if p_x <= sw_x && p_y >= ne_y {
		fmt.Println("NW")
		return
	}
	if p_y > sw_y && p_y < ne_y {
		if p_x > ne_x {
			fmt.Println("E")
		} else {
			fmt.Println("W")
		}
		return
	}
	if p_x > sw_x && p_x < ne_x {
		if p_y < ne_y {
			fmt.Println("S")
		} else {
			fmt.Println("N")
		}
		return
	}

}
