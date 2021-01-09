package main

import (
	"fmt"
)

var (
	Version string = "1.0.0"
	Build   string = "dev"
)

func main() {
	fmt.Println("Version: ", Version)
	fmt.Println("Build: ", Build)
}
