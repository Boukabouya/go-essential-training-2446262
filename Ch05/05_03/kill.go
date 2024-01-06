package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

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

// sets up logging by opening a file for append mode
func setupLogging() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(out)
}

func killServer(pidFile string) error {
	setupLogging()
	// TODO
	id, err := ioutil.ReadFile(pidFile)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	pidStr := string(id) // ensure that the id is string
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		// Handle the error, e.g., return an error or log it
		fmt.Println("Error converting string to integer:", err)
	}
	// Simulate kill
	fmt.Printf("killing server with pid=%d\n", pid)
	return nil
}

func main() {
	if err := killServer("server.pid"); err != nil {
		log.Fatalf("error: %s", err)
	}
}
