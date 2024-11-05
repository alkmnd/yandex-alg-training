package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	I = "#\n"
	O = "###\n#.#\n###\n"
	C = "##\n#.\n##\n"
	L = "#.\n##\n"
	H = "#.#\n###\n#.#\n"
	P = "###\n#.#\n###\n#..\n"
)

func main() {
	var n int
	fmt.Scan(&n)
	table := make([][]rune, 0)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var row []rune
		for _, char := range line {

			row = append(row, char)
		}

		table = append(table, row)
	}

	if len(table) == 0 {
		fmt.Println("X")
		return
	}
	var rowStack [][]rune
	// Удаляем не горящие строки по бокам.
	trimRowStart := 0
	for row := 0; row < n; row++ {
		if isGapRow(table, n, row) {
			trimRowStart++
		} else {
			break
		}
	}

	if trimRowStart == len(table) {
		fmt.Println("X")
		return
	}

	trimRowEnd := n
	for row := n - 1; row >= 0; row-- {
		if isGapRow(table, n, row) {
			trimRowEnd--
		} else {
			break
		}
	}

	trimColStart := 0
	// Удаляем не горящие столбцы по бокам.
	for col := 0; col < n; col++ {
		if isGapCol(table, trimRowStart, trimRowEnd, col) {
			trimColStart++
		} else {
			break
		}
	}

	trimColEnd := n
	for col := n - 1; col >= 0; col-- {
		if isGapCol(table, trimRowStart, trimRowEnd, col) {
			trimColEnd--
		} else {
			break
		}
	}

	// Удаляем дубликаты.
	for row := trimRowStart; row < trimRowEnd; row++ {
		if len(rowStack) != 0 {
			if isRowsEqual(rowStack[len(rowStack)-1], table[row]) {
				continue
			} else {
				rowStack = append(rowStack, table[row])
			}
		} else {
			rowStack = append(rowStack, table[row])
		}

	}

	trimColEnd = trimColEnd - trimColStart
	for i := range rowStack {
		if trimColStart < len(rowStack[i]) {
			rowStack[i] = rowStack[i][trimColStart:]

			rowStack[i] = rowStack[i][:trimColEnd]
		}
	}

	res := make([][]rune, len(rowStack))
	for row := 0; row < len(rowStack); row++ {
		res[row] = rowStack[row][:1]
	}
	step := 1
	for col := 0; col < len(rowStack[0])-1; {
		if col+step >= len(rowStack[0]) {
			break
		}
		if isColsEqual(rowStack, col, col+step) {
			step++
		} else {
			for row := 0; row < len(rowStack); row++ {
				res[row] = append(res[row], rowStack[row][col+step])
			}
			col = col + step
			step = 1
		}
	}

	var result string

	for _, row := range res {
		result += string(row) + "\n"
	}

	//fmt.Print(result)

	switch result {
	case I:
		fmt.Println("I")
	case O:
		fmt.Println("O")
	case L:
		fmt.Println("L")
	case H:
		fmt.Println("H")
	case P:
		fmt.Println("P")
	case C:
		fmt.Println("C")
	default:
		fmt.Println("X")

	}
}

func isRowsEqual(row1 []rune, row2 []rune) bool {
	n := len(row1)
	for i := 0; i < n; i++ {
		if row1[i] != row2[i] {
			return false
		}
	}
	return true
}

func isGapRow(table [][]rune, n int, row int) bool {

	for col := 0; col < n; col++ {
		if table[row][col] == '#' {
			return false
		}
	}

	return true
}

func isGapCol(table [][]rune, start, end, col int) bool {
	for row := start; row < end; row++ {
		if table[row][col] == '#' {
			return false
		}
	}

	return true
}

func isColsEqual(table [][]rune, col1, col2 int) bool {
	for row := 0; row < len(table); row++ {
		if table[row][col1] != table[row][col2] {
			return false
		}
	}
	return true
}
