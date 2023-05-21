package main

import "fmt"

func main() {
	var lfNum int
	var operator string
	var rgNum int

	fmt.Print("> ")
	fmt.Scan(&lfNum, &operator, &rgNum)
	switch operator {
	case "+":
		fmt.Println("Result", lfNum+rgNum)
	case "-":
		fmt.Println("Result", lfNum-rgNum)
	case "*":
		fmt.Println("Result", lfNum*rgNum)
	case "/":
		fmt.Println("Result", lfNum/rgNum)
	case "%":
		fmt.Println("Result", lfNum%rgNum)
	default:
		fmt.Println("Please insert correct value!")
	}
}
