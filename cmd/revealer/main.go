package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/none-da/revealer/pkg/revealer/spec"
	flag "github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var secretsFile string
	flag.StringVarP(&secretsFile, "secrets-file", "f", "", "Secrets file to parse.")
	var printHelp bool
	flag.BoolVarP(&printHelp, "help", "h", false, "Prints this help content.")
	flag.Parse()
	if printHelp {
		flag.Usage()
		return
	}
	_, err := os.Stat(secretsFile)
	if (err != nil) && (os.IsNotExist(err)) {
		fmt.Println("file: ", secretsFile, " Does Not Exist")
	}

	var infraFileData spec.InfraFileSpec

	data, err := ioutil.ReadFile(secretsFile)
	fmt.Println("data: ", data)
	err = yaml.UnMarshal(data, &infraFileData)
	fmt.Println("infraFileData: ", infraFileData)
	panicOnError(err)
}
