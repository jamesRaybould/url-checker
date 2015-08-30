//This is a self-loading config, anytime you add it as an import
//it will reach down and grab a copy of the config.yml and populate Settings automatically
package model

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	ApiUrl  string `yaml:"api_url"`
	BaseUrl string `yaml:"base_url"`
}

var Settings Config

func init() {
	log.Println("Loading config.yml")

	source, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(source, &Settings)
	if err != nil {
		panic(err)
	}
}
