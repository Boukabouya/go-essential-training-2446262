package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/* In this challenge, you are going to implement the io.Writer interface,
which has a one method of Write, which has a single method called Write
 that gets a slice of bytes and returns how many bytes were written and if there was an error.
 Here we have our Capper, which is turning everything into uppercase.
 If you're going to wrap Capper over the standard output and then print into the Capper,
 we should see "Hello there" printed all caps to the standard output.*/

// Capper implements io.Writer and turns everything to uppercase

type Capper struct {
	// TODO
	write io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	// Convert the input to uppercase
	upperCaseData := []byte(strings.ToUpper(string(p)))

	// Write the uppercase data to the underlying writer
	n, err = c.write.Write(upperCaseData)
	return n, err
}
func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello there")
}
