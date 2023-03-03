package main

import (
	"flag"
	"fmt"
)

const VERSION = "1.0.0"

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "查看版本")

	flag.Parse()

	if showVersion {
		fmt.Println(VERSION)
	}
}
