package main

import "fmt"

func main() {
	var operator1, operator2 int
	var operation string
	fmt.Println("Please enter 1st operator")
	fmt.Scanf("%d ", &operator1)
	fmt.Println("Please enter 2nd operator")
	fmt.Scanf("%d ", &operator2)
	fmt.Println("Please choose the operation (+-*/) ")
	fmt.Scanf("%s ", &operation)
	switch operation {
	case "+":
		result := operator1 + operator2
		fmt.Println(result)
	case "-":
		result := operator1 - operator2
		fmt.Println(result)
	case "*":
		result := operator1 * operator2
		fmt.Println(result)
	}

}
