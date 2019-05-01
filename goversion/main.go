package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	flagVersion bool
	version     string
	buildTime   string
)

func main() {
	flag.BoolVar(&flagVersion, "v", false, "if true, print version and exit")
	flag.BoolVar(&flagVersion, "version", false, "if true, print version and exit")
	flag.Parse()
	if flagVersion {
		fmt.Println("version: ", version)
		fmt.Println("buildTime: ", buildTime)
		os.Exit(0)
	}

	fmt.Println("Hellow World")
	fmt.Println("First Update")
}
