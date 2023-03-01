// Factorial of integer

package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	result := factorial(num)
	fmt.Printf("Factorial of %d is %d", num, result)
}

func factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}
