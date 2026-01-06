package main

import (
	"fmt"

	"example.com/utils"
)

func getNumbers(a int, b int) (int, int) {
	sum := a + b
	mul := a * b

	return sum, mul
}
func main() {
	fmt.Println("Hello world")
	var res int = 10
	fmt.Println(res)

	// switch
	age := 100
	switch age {
	case 10:
		fmt.Println("You are too samll")
		fmt.Println("Please register when you are 18")
	case 15, 100:
		fmt.Println("Just three years wait.....")
	default:
		fmt.Println("Not reachable....")
	}

	// return two
	add, mul := getNumbers(10, 20)
	fmt.Println(add)
	fmt.Println(mul)

	myApp := 1

	if myApp == 1 {
		ans := 20
		fmt.Println(ans)
	}
	// package scope
	ans := sum(10, 20) // called sum.go -> run: go run main.go sum.go
	fmt.Println(ans)

	// using utils package
	sumAAndB := utils.Add(1, 2)
	fmt.Println(sumAAndB)
}
