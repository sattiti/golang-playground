package main

import (
	"fmt"

	"localhost.com/greetings/greetings"
)

// Create module
// ex: go mod init example.com/greetings
// go mod init MODULE_PATH_NAME

func main() {
	message := greetings.Hello("John.")
	fmt.Printf(message)
}
