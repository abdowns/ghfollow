package main

import (
	"encoding/json"
	"io"
	"os"
	"strings"
)

type Config struct {
	Token       string `json:"token"`
	RateLimit   int    `json:"rate-limit"`
	HTTPTimeout int    `json:"http-timeout"`
}

func GetConfig() *Config {

	content, err := os.ReadFile("config.json")
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			GenConfig()
			panic("config.json not found, generated for you.\nPlease put your github token with the 'user:follow' scope into the 'token' json field in config.json")
		}
		panic(err)
	}

	result := &Config{}

	err = json.Unmarshal(content, result)
	if err != nil {
		panic(err)
	}

	return result
}

func GenConfig() {
	source, err := os.Open("config.example.json")
	if err != nil {
		panic(err)
	}
	defer source.Close()

	dest, err := os.Create("config.json")
	if err != nil {
		panic(err)
	}
	defer dest.Close()

	_, err = io.Copy(dest, source)
	if err != nil {
		panic(err)
	}
}
