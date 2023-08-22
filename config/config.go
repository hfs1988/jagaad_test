package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ThirdParty ThirdParty `yaml:"THIRDPARTY"`
	Data       Data       `yaml:"DATA"`
}

type ThirdParty struct {
	URL string `yaml:"URL"`
}

type Data struct {
	UserFilename        string `yaml:"USERFILENAME"`
	UserTagsFilename    string `yaml:"USERTAGSFILENAME"`
	UserFriendsFilename string `yaml:"USERFRIENDSFILENAME"`
}

func GetConfig() Config {
	file, err := os.ReadFile("./app.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var conf Config

	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		log.Fatal(err)
	}

	return conf
}
