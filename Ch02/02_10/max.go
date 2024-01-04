// Calculate maximal value in a slice
package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	max := nums[0]
	for _, name := range nums {

		if max > name {
			continue
		} else {
			max = name
		}
	}
	fmt.Println(max)
}
