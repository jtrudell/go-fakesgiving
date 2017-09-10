package config

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const envFile = ".env"

type Env struct {
	Port   string
	DBName string
	DBUser string
}

// Setup opens and reads .env file in root directory and parses lines
func Setup() Env {
	file, err := os.Open(envFile)
	if err != nil {
		panic("Failed to open " + envFile)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic("Failed to read content from " + envFile)
	}
	file.Close()

	envLines := strings.Split(string(content), "\n")
	return setEnv(envLines)
}

// setEnv sets environment variables
func setEnv(envLines []string) Env {
	for _, line := range envLines {
		values := strings.Split(line, "=")
		if len(values) > 1 {
			err := os.Setenv(values[0], values[1])
			if err != nil {
				log.Fatalln("Could not set environment variable" + values[0])
			}
		}
	}

	return Env{
		Port:   os.Getenv("PORT"),
		DBName: os.Getenv("DBNAME"),
		DBUser: os.Getenv("DBUSER"),
	}
}
