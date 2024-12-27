package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	for name, _ := range dist {
		dist[name] = getDist(name, parents, dist)
	}

	keys := make([]string, 0, len(dist))
	for k := range dist {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, dist[k])
	}

}

func getDist(name string, parents map[string]string, dist map[string]int) int {
	if dist[name] == -1 {
		if _, ok := parents[name]; !ok {
			dist[name] = 0
		} else {
			dist[name] = getDist(parents[name], parents, dist) + 1
		}
	}
	return dist[name]
}
