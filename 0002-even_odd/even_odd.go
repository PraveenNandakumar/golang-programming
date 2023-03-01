// Integer - Even or odd

package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("The entered number is a even number.")
	} else {
		fmt.Println("The entered number is a odd number.")
	}
}
