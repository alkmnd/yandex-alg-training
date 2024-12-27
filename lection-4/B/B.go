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

	cnt := make(map[string]int)
	children := make(map[string][]string)

	for i := 0; i < N-1; i++ {
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		data := strings.Split(str, " ")
		child, parent := data[0], data[1]
		children[parent] = append(children[parent], child)
		cnt[child] = -1
		cnt[parent] = -1
	}

	for name, _ := range cnt {
		cnt[name] = getCnt(name, children, cnt)
	}

	keys := make([]string, 0, len(cnt))
	for k := range cnt {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, cnt[k])
	}

}

func getCnt(name string, children map[string][]string, cnt map[string]int) int {
	if cnt[name] == -1 {
		cnt[name] = 0
		if _, ok := children[name]; ok {
			for _, ch := range children[name] {
				cnt[name] += getCnt(ch, children, cnt) + 1
			}
		}
	}
	return cnt[name]
}
