package _50_Evaluate_Reverse_Polish_Notation

import (
	"fmt"
	"strconv"
)

//Time: O(n); Space: O(1)
func evalRPN(tokens []string) int {
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			a, _ := strconv.Atoi(tokens[i-2])
			b, _ := strconv.Atoi(tokens[i-1])
			tokens[i-2] = fmt.Sprintf("%d", a+b)
			tokens = append(tokens[:i-1], tokens[i+1:]...)
			i -= 2
		case "-":
			a, _ := strconv.Atoi(tokens[i-2])
			b, _ := strconv.Atoi(tokens[i-1])
			tokens[i-2] = fmt.Sprintf("%d", a-b)
			tokens = append(tokens[:i-1], tokens[i+1:]...)
			i -= 2
		case "*":
			a, _ := strconv.Atoi(tokens[i-2])
			b, _ := strconv.Atoi(tokens[i-1])
			tokens[i-2] = fmt.Sprintf("%d", a*b)
			tokens = append(tokens[:i-1], tokens[i+1:]...)
			i -= 2
		case "/":
			a, _ := strconv.Atoi(tokens[i-2])
			b, _ := strconv.Atoi(tokens[i-1])
			tokens[i-2] = fmt.Sprintf("%d", a/b)
			tokens = append(tokens[:i-1], tokens[i+1:]...)
			i -= 2
		}
	}
	result, _ := strconv.Atoi(tokens[0])
	return result
}

func evalRPN1(tokens []string) int {
	var result []int
	for i := 0; i < len(tokens); i++ {
		switch t := tokens[i]; t {
		case "+", "-", "*", "/":
			a, b := result[len(result)-2], result[len(result)-1]
			result = result[:len(result)-2]
			var res int
			if t == "+" {
				res = a + b
			} else if t == "-" {
				res = a - b
			} else if t == "*" {
				res = a * b
			} else {
				res = a / b
			}
			result = append(result, res)
		default:
			i, _ := strconv.Atoi(t)
			result = append(result, i)
		}
	}
	return result[0]
}