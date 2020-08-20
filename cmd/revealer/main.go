package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

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
	var printVersion bool
	flag.BoolVarP(&printVersion, "version", "v", false, "Prints the version of Revealer.")
	flag.Parse()
	if printHelp {
		flag.Usage()
		return
	}
	if printVersion {
		fmt.Println("v" + strings.Join(VERSION[:], "."))
		os.Exit(0)
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
	var settingsKeys []string
	for suppliedSettingsKey := range infraFileData.InfraSettings {
		settingsKeys = append(settingsKeys, suppliedSettingsKey)
	}
	sort.Strings(settingsKeys)

	for _, settingsKey := range settingsKeys {
		settingsValue := infraFileData.InfraSettings[settingsKey]
		decodedSettingsValue, err := base64.StdEncoding.DecodeString(settingsValue)
		if err != nil {
			fmt.Println("ERROR! value:", settingsValue, "is not a base64 encoded value. Please verify")
			os.Exit(1)
		}
		cleanedUpSettingsValue := base64.StdEncoding.EncodeToString([]byte(strings.TrimSpace(string(decodedSettingsValue))))
		if settingsValue != cleanedUpSettingsValue {
			fmt.Println("WARNING! Newlines are added in the supplied settingsValue:", settingsValue, "for settingsKey:", settingsKey, "Correct value should be:", cleanedUpSettingsValue)
			decodedSettingsValue, _ = base64.StdEncoding.DecodeString(cleanedUpSettingsValue)
		}

		fmt.Print(settingsKey, ": ", string(decodedSettingsValue))
		fmt.Println()
	}
	panicOnError(err)
}
