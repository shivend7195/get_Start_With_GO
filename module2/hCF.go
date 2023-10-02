// finding the hcf(highest common factore) also known as GCD
// largest positive integer that can divide both the numbers

package main

import (
	"fmt"
)

func main() {
	var n1, n2 int
	fmt.Scanln(&n1, &n2)
	ans := 0

	for i := 1; i <= n1 && i <= n2; i++ {
		if n1%i == 0 && n2%i == 0 {
			ans = i
		}
	}
	fmt.Println(ans)
}
