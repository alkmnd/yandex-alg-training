package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)
	reader := bufio.NewReader(os.Stdin)
	graph := make(map[int][]int)
	des := make([]int, N+1)
	for i := range des {
		des[i] = -1
	}

	for i := 0; i < N-1; i++ {
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		data := strings.Split(str, " ")
		aS, bS := data[0], data[1]
		a, _ := strconv.Atoi(aS)
		b, _ := strconv.Atoi(bS)
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	cnt(1, des, graph)

	for _, i := range des[1:] {
		fmt.Printf("%d ", i)
	}
}

func cnt(num int, des []int, graph map[int][]int) int {
	des[num] = 1
	for _, s := range graph[num] {
		if des[s] == -1 {
			des[num] += cnt(s, des, graph)
		}
	}

	return des[num]

}
