/*
An "even ended number" is a number whose first and last digit are the same.

You mission, should you choose to accept it, is to count how many "even ended numbers" are
there that are a multiplication of two 4 digit numbers.
*/

package main

import (
	"fmt"
)

func main() {
	/*n := 4244
	s := fmt.Sprintf("%d", n) //return string , so n convert to string type

	fmt.Printf("s = %q (type %T)\n", s, s) //%q will add the quotes around the string*/

	for i := 1000; i <= 9999; i++ {
		s := fmt.Sprintf("%d", i)
		if s[0] == s[len(s)-1] {
			for j := 0; j <= 4; j++ {
				continue
			}
			fmt.Println("The number", i, "is An even ended number")
		} else {
			fmt.Println("The number", i, "is not an even ended number")
		}
	}
}
