package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Token       string `json:"token"`
	RateLimit   int    `json:"rate-limit"`
	HTTPTimeout int    `json:"http-timeout"`
}

func GetConfig() *Config {

	content, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	result := &Config{}

	err = json.Unmarshal(content, result)
	if err != nil {
		panic(err)
	}

	return result

	// if not found then generator for user
}
