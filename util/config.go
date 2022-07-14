package util

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type config struct {
	Port     int         `json:"port"`
	Accounts [][2]string `json:"accounts"`
	Proxy    []string    `json:"proxy"`
	// TODO: cors
}

func updatePasswords(pw [][2]string) {
	for i := range pw {
		sum := md5.Sum([]byte(pw[i][1]))
		pw[i][1] = fmt.Sprintf("%x", sum)
	}
}

func LoadCfg() *config {
	var cfg config
	bytes, err := ioutil.ReadFile("config.json")

	if err != nil {
		log.Println("Could not read config.json.")
		os.Exit(1)
	}

	if err := json.Unmarshal(bytes, &cfg); err != nil {
		log.Println("Could not read unmarshal config.json", err)
		os.Exit(1)
	}

	updatePasswords(cfg.Accounts)
	return &cfg
}
