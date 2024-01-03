// Package main is the entry point for the executable program.
package main

import (
	"fmt" // The fmt package provides formatted I/O functions like Println.
)

// The main function is the entry point of the program.
func main() {
	// fmt.Println is used to print the specified message to the console.
	fmt.Println("Welcome Gophers ☺")
}

/*- Let's have a look at the go program. Our code start with the package main.
 Every go code lives inside the package and all the files in the same folder on the disc should be in the same package.
  Package Main is a special package that will make your code compile to an executable and not a library. After that, we have import.
  Go is a simple language and most of the functionality is inside other packages we need to import. In this case,
  we are importing the FMT package for formatted output. On line seven, we define main, which is a function. We use the func keyword to define the function,
  and then the name of the function. This function gets no parameters,
  and then the function body start with the curly braces. Do not try to put the curly braces anywhere else.
  The go code is automatically formatted with an external tool and every go code looks exactly the same. On line eight, we have a function call.
  We call the print ln from the FMT package. We always prefix the package name before the function we're calling. So FMT dot, and then we have print ln.
  Print ln start with a capital T, meaning it's an exported symbol. Every symbol or function you're going to use
  from another package is going to start with an upper case. We do print ln, and then we open parenthesis for the arguments.
  And we have a single argument which is string, starting and ending with double quotes, saying welcome gophers, and we have the smiley sign.
   Strings in go are UTF eight. And we have the closing parenthesis for closing the arguments for the function code. And we don't have a semicolon.
    Most go code will not have any semicolon in it. And finally we have the closing curly braces for the function body.
	To run the code, you can go to the run menu, and do run without debugging. If you don't see the output, click on view and debug console.
	 Their output is from the ID and my output is the one in blue. Welcome gophers.*/
