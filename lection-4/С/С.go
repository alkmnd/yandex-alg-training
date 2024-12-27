package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)
	reader := bufio.NewReader(os.Stdin)

	dist := make(map[string]int)
	parents := make(map[string]string)

	for i := 0; i < N-1; i++ {
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		data := strings.Split(str, " ")
		child, parent := data[0], data[1]
		parents[child] = parent
		dist[child] = -1
		dist[parent] = -1
	}

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		str = strings.TrimSpace(str)
		query := strings.Split(str, " ")
		name1, name2 := query[0], query[1]
		
		ancestors := make(map[string]struct{})
		ancestors[name1] = struct{}{}

		for {
			if parent, ok := parents[name1]; ok {
				ancestors[parent] = struct{}{}
				name1 = parent
			} else {
				break
			}
		}

		for {
			if parent, ok := parents[name2]; ok {
				if _, found := ancestors[name2]; found {
					fmt.Println(name2)
					break
				}
				name2 = parent
			} else {
				fmt.Println(name2)
				break
			}
		}
	}

}
