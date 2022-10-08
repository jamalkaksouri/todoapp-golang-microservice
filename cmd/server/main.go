package main

import (
	"fmt"
	"github.com/jamalkaksouri/todoapp-golang-microservice/pkg/cmd"
	"os"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
