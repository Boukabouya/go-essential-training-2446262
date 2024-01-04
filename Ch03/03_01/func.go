// Basic function definition
package main

import (
	"fmt"
)

func main() {
	/*. Go is passing parameters by value to function, 
	meaning it's creating a copy of the integer and passing it to the double function. 
	A slice, like the one we are passing to doubleAt, behaves like a pointer,
	 meaning we create a copy of the slice but it points to the same location in memory,
	  so the effect is reflected outside of the function. So how can we make double work? By using */
	val := add(1, 2)
	fmt.Println(val)

	div, mod := divmod(7, 2)
	fmt.Printf("div=%d, mod=%d\n", div, mod)
}

// add adds a to b
func add(a int, b int) int {
	return a + b
}

// divmod returns quotient and reminder
func divmod(a int, b int) (int, int) {
	return a / b, a % b
}
