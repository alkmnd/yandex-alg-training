package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BinarySearchTree struct {
	Root *TreeNode
}

func (bst *BinarySearchTree) Search(value int) bool {
	return bst.Root != nil && bst.Root.search(value)
}

func (node *TreeNode) search(value int) bool {
	if node == nil {
		return false
	}
	if value < node.Value {
		return node.Left.search(value)
	} else if value > node.Value {
		return node.Right.search(value)
	} else {
		return true
	}
}

func (bst *BinarySearchTree) Insert(value int) bool {
	newNode := &TreeNode{Value: value}
	if bst.Root == nil {
		bst.Root = newNode
		return true
	} else {
		return bst.Root.insert(newNode)
	}
}

func (node *TreeNode) insert(newNode *TreeNode) bool {
	if node.Value == newNode.Value {
		return false
	}
	if newNode.Value < node.Value {
		if node.Left == nil {
			node.Left = newNode
			return true
		} else {
			return node.Left.insert(newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
			return true
		} else {
			return node.Right.insert(newNode)
		}
	}
}

func (bst *BinarySearchTree) Print() {
	bst.Root.print(0)

}

func (node *TreeNode) print(depth int) {
	if node == nil {
		return
	}
	node.Left.print(depth + 1)
	dots := ""
	for i := 0; i < depth; i++ {
		dots += "."
	}
	fmt.Printf("%s%d\n", dots, node.Value)
	node.Right.print(depth + 1)

}

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

func main() {
	tree := &BinarySearchTree{Root: nil}
	reader := bufio.NewReader(os.Stdin)

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		str = strings.TrimSpace(str)
		query := strings.Split(str, " ")

		if len(query) == 1 {
			if query[0] == "PRINTTREE" {
				tree.Print()
			} else {
				return
			}
		} else {
			value, _ := strconv.Atoi(query[1])
			if query[0] == "ADD" {
				res := tree.Insert(value)
				if res {
					fmt.Println("DONE")
				} else {
					fmt.Println("ALREADY")
				}
			} else if query[0] == "SEARCH" {
				res := tree.Search(value)
				if res {
					fmt.Println("YES")
				} else {
					fmt.Println("NO")
				}
			} else {
				return
			}
		}
	}
}
