package main

import (
	"os"

	"github.com/hiroakis/bb"
)

func main() {
	os.Exit(bb.Run(os.Args[1:]))
}
