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
	n := 4244
	s := fmt.Sprintf("%d", n) //return string , so n convert to string type

	fmt.Printf("s = %q (type %T)\n", s, s) //%q will add the quotes around the string

	if len(s) == 4 && s[0] == s[len(s)-1] {
		for i := 0; i <= 4; i++ {
			continue
		}
		fmt.Println("An even ended number")
	} else {
		fmt.Println("Not an even ended number")
	}
}
