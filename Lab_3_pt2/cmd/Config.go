package main

type Config struct {
	Files []struct {
		Filename  string `yaml:"filename"`
		Substring string `yaml:"substring"`
	} `yaml:"files"`
}
