package main

import (
	"fmt"
	"github.com/brgv/sv/internal/version"
)

func main() {
	fmt.Println("Hey!")
	version.BuildPrintln()
}
