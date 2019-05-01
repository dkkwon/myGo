package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	flagVersion        bool
	version, buildTime string
)

type goVersion struct {
	version, buildTime string
}

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

	var verInfo goVersion
	verInfo.version = version
	verInfo.buildTime = buildTime

	fmt.Println("version: ", verInfo.version)
	fmt.Println("buildTime: ", verInfo.buildTime)

}
