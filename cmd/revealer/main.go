package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	specs "github.com/none-da/revealer/pkg/specs"
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
	if secretsFile == "" {
		fmt.Println("Please supply secrets-file!!!")
		os.Exit(1)
	}
	_, err := os.Stat(secretsFile)
	if (err != nil) && (os.IsNotExist(err)) {
		fmt.Println("file", secretsFile, "Does Not Exist")
	}

	var infraFileData specs.InfraFileSpec

	data, err := ioutil.ReadFile(secretsFile)
	err = yaml.Unmarshal(data, &infraFileData)
	// fmt.Println("infraFileData: ", infraFileData)
	for settingsKey, settingsValue := range infraFileData.InfraSettings {
		decodedSettingsValue, err := base64.StdEncoding.DecodeString(settingsValue)
		if err != nil {
			fmt.Println("Error! value:", settingsValue, "is not a base64 encoded value. Please verify")
			os.Exit(1)
		}
		fmt.Println(settingsKey, "=>", string(decodedSettingsValue))
	}
	panicOnError(err)
}
