package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

func init() {
	var secretsFile string
	flag.StringVarP(&secretsFile, "secrets-file", "f", "", "Secrets file to parse.")
	var printHelp bool
	flag.BoolVarP(&printHelp, "help", "h", false, "Prints this help content.")
	flag.Parse()
	if printHelp {
		flag.Usage()
		return
	}
}

func main() {
	fmt.Println("I read secrets and I reveal them.")
}
