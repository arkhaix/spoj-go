package main

import "fmt"
import "strings"

func push(self []string, value string) []string {
	return append(self, value)
}

func pop(self []string) ([]string, string) {
	lastIndex := len(self) - 1
	if lastIndex < 0 {
		return self, ""
	}

	value := self[lastIndex]
	return self[:lastIndex], value
}

func top(self []string) string {
	i := len(self) - 1
	if i < 0 {
		return ""
	}
	return self[i]
}

func convert(infix string) {
	precedence := "()+-*/^"
	operands := make([]string, 0)
	var val string

	for _, r := range infix {
		token := string(r)
		if strings.Contains(precedence, token) == false {
			fmt.Printf("%s", token)
		} else if token == "(" {
			operands = push(operands, token)
		} else if token == ")" {
			operands, val = pop(operands)
			for val != "(" && len(operands) > 0 {
				fmt.Printf("%s", val)
				operands, val = pop(operands)
			}
		} else {
			for len(operands) > 0 && strings.Index(precedence, top(operands)) >= strings.Index(precedence, token) {
				operands, val = pop(operands)
				if val != "(" {
					fmt.Printf("%s", val)
				}
			}
			operands = push(operands, token)
		}
	}

	for len(operands) > 0 {
		operands, val = pop(operands)
		fmt.Printf("%s", val)
	}

	fmt.Println("")
}

func main() {
	numTests := 0
	fmt.Scanln(&numTests)
	for i := 0; i < numTests; i++ {
		var infix string
		fmt.Scanln(&infix)
		convert(infix)
	}
}
