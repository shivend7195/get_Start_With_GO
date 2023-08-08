// fibonachhi upto n terms
package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	first := 0
	second := 1
	fmt.Print(first, ",", second)
	var nextTerm int
	for i := 0; i < n; i++ {
		nextTerm = first + second
		first = second
		second = nextTerm

		fmt.Print(",", nextTerm)
	}
}
