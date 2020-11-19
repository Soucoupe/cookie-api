package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type (
	Config struct {
		Server    Server
		Database  Database
		Ratelimit Ratelimit
	}

	Server struct {
		Address string
	}

	Database struct {
		Addr     string
		Username string
		Password string
		Database string
	}

	Ratelimit struct {
		GenPerSecond float64 `json:"gen-per-second"`
	}
)

var (
	Conf Config
)

func LoadConfig() {
	log.Println("Loading config")

	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &Conf)
	if err != nil {
		panic(err)
	}
}
