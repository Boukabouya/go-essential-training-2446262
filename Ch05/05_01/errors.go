// pkg/errors example
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

// Config holds configuration
type Config struct {
	// configuration fields go here (redacted)
}
// read the error configue file 
func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "can't open configuration file")
	}
	defer file.Close()

	cfg := &Config{}
	// Parse file here (redacted)
	return cfg, nil

}
//sets up logging by opening a file for append mode
func setupLogging() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(out)
}

func main() {
	setupLogging()
	cfg, err := readConfig("/path/to/config.toml")
	if err != nil {
		//os.Stderr: os.Stderr is an io.Writer that represents the standard error stream. Anything written to it will be displayed as an error message in the console.
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		log.Printf("error: %+v", err)
		os.Exit(1)
	}

	// Normal operation (redacted)
	fmt.Println(cfg)
}
