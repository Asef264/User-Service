package config

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v1"
)

type configs struct {
	Scylladb ScylladbConfigs `yaml:"scylladb"`
	User     UserConfig      `yaml:"user_config"`
}

var Configs configs

func Init() {
	// getting base address of running dir
	_, file, _, _ := runtime.Caller(0)
	BasePath := filepath.Dir(file)
	// adding config file address to the base address
	configPath := BasePath + "/file/config.yaml"

	// reading the config file
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("error on reading config.yaml file, err: %v", err)
	}
	// trying to unmarshal config data
	err = yaml.Unmarshal(yamlFile, &Configs)
	if err != nil {
		log.Fatalf("error on unmarshalling config file, error: %v", err)
	}
}
