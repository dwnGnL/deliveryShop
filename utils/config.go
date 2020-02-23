package utils

import (
	"encoding/json"
	"io/ioutil"
)

func ReadConfig() *Config {
	var appParams Config
	doc, err := ioutil.ReadFile("./config.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(doc, &appParams)
	if err != nil {
		panic(err)
	}

	return &appParams
}
