// String reversal

package main

import "fmt"

func main() {

	var str string

	fmt.Print("Enter an input string: ")
	fmt.Scan(&str)

	reversed := reverse(str)
	fmt.Printf("Reversed string is: %s", reversed)

}

func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
