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
	values := []int{1, 2, 3, 4}
	doubleAt(values, 2)
	fmt.Println(values)

	val := 10
	double(val)
	fmt.Println(val)

	doublePtr(&val)
	fmt.Println(val)
}

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(n int) {
	n *= 2
}

func doublePtr(n *int) {
	*n *= 2
}
