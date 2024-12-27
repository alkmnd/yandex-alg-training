package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func isInteger(s string) (int, bool) {
	v, err := strconv.Atoi(s)
	return v, err == nil
}

// Проверка на правильность скобок
func validateBrackets(expr string) bool {
	stack := []rune{}
	for _, char := range expr {
		if char == '(' {
			stack = append(stack, char)
		} else if char == ')' {
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func validateOperationsBetweenNumbers(expr string) bool {
	parts := strings.Split(expr, " ")
	for i := 1; i < len(parts); i++ {
		if unicode.IsDigit(rune(parts[i-1][len(parts[i-1])-1])) && unicode.IsDigit(rune(parts[i][0])) ||
			rune(parts[i-1][len(parts[i-1])-1]) == ')' && unicode.IsDigit(rune(parts[i][0])) ||
			unicode.IsDigit(rune(parts[i-1][len(parts[i-1])-1])) && rune(parts[i][0]) == '(' {
			return false
		}
	}

	return true

}

func validateAllowedCharacters(expr string) bool {
	for i := 0; i < len(expr); i++ {
		if !unicode.IsDigit(rune(expr[i])) &&
			!isOperator(rune(expr[i])) &&
			!(rune(expr[i]) == '(') &&
			!(rune(expr[i]) == ')') &&
			!(rune(expr[i]) == ' ') {
			return false
		}
	}
	return true
}

// Проверка, является ли символ оператором
func isOperator(char rune) bool {
	return char == '+' || char == '-' || char == '*'
}

// Основная функция для проверки валидности выражения
func validateExpression(expr string) bool {
	return validateAllowedCharacters(expr) &&
		validateBrackets(expr) && validateOperationsBetweenNumbers(expr)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	expr, _ := reader.ReadString('\n')
	expr = strings.TrimSpace(expr)

	var parts []string
	if !validateExpression(expr) {
		fmt.Println("WRONG")

	} else {
		expr = strings.Replace(expr, " ", "", -1)
		var num string
		for i, v := range expr {
			if unicode.IsDigit(v) {
				num += string(v)
				if i == len(expr)-1 {
					parts = append(parts, num)
				}
				continue
			}
			if num != "" {
				parts = append(parts, num)
				num = ""
			}
			parts = append(parts, string(v))

		}
		var postfix []string
		stack := make([]string, 0, len(parts))
		for i := range parts {
			if parts[i] == "+" {
				for len(stack) != 0 && (stack[len(stack)-1] == "+" ||
					stack[len(stack)-1] == "*" ||
					stack[len(stack)-1] == "-") {
					postfix = append(postfix, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, parts[i])

			} else if parts[i] == "-" {
				for len(stack) != 0 && (stack[len(stack)-1] == "+" ||
					stack[len(stack)-1] == "*" ||
					stack[len(stack)-1] == "-") {
					postfix = append(postfix, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, parts[i])
			} else if parts[i] == "*" {
				for len(stack) != 0 && stack[len(stack)-1] == "*" {
					postfix = append(postfix, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, parts[i])
			} else if parts[i] == "(" {
				stack = append(stack, parts[i])
			} else if parts[i] == ")" {
				for len(stack) != 0 && stack[len(stack)-1] != "(" {
					postfix = append(postfix, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				stack = stack[:len(stack)-1]
			} else if _, err := strconv.Atoi(parts[i]); err == nil {
				postfix = append(postfix, parts[i])
			}
		}

		for j := len(stack) - 1; j >= 0; j-- {
			postfix = append(postfix, stack[j])
		}
		postStack := make([]int, 0, len(parts))
		for _, v := range postfix {
			if num, flag := isInteger(v); flag != false {
				postStack = append(postStack, num)
			} else {
				var val int
				n := len(postStack)
				if !(n-1 > -1 && n-2 > -1) {
					fmt.Println("WRONG")
					return
				}
				switch v {
				case "+":
					val = postStack[n-2] + postStack[n-1]
				case "-":
					val = postStack[n-2] - postStack[n-1]
				case "*":
					val = postStack[n-2] * postStack[n-1]
				}
				postStack = postStack[:n-2]
				postStack = append(postStack, val)
			}
		}
		fmt.Println(postStack[0])
	}

}
