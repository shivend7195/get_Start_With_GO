// lcm of two numbers (lowest common multiple)

package main

import (
	"fmt"
	"math"
)

func main() {
	var n1, n2 int
	fmt.Scanln(&n1, &n2)
	ans := 0
	max := int(math.Max(float64(n1), float64(n2)))
	for i := max; i <= (n1 * n2); i++ {
		if i%n1 == 0 && i%n2 == 0 {
			ans = i
			break
		}
	}
	fmt.Println(ans)
}
