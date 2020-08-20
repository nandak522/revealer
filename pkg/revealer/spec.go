package main

type InfraFileSpec struct {
	InfraSettings map[string]string `yaml:"infraSettings" validate:"required"`
}
