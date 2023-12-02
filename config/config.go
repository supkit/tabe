package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// config 配置
var config Config

// Config 结构
type Config struct {
	Env   string `yaml:"env"`
	Debug bool   `yaml:"debug"`
	Log   struct {
		FileName      string `yaml:"fileName"`
		MaxAgeDay     int    `yaml:"maxAgeDay"`
		FileFormat    string `yaml:"fileFormat"`
		DistWarnLevel bool   `yaml:"distWarnLevel"`
		CallerSkip    int    `yaml:"callerSkip"`
	} `yaml:"log"`
	Client []ConfigClient `yaml:"client"`
}

// ConfigClient client
type ConfigClient struct {
	Name    string `yaml:"name"`
	Target  string `yaml:"target"`
	Timeout int64  `yaml:"timeout"`
}

// Watch load config
func Watch(path string) (err error) {
	// ...
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return
	}

	return err
}

// GetClientByName get client
func GetClientByName(name string) (client ConfigClient, err error) {
	for _, v := range config.Client {
		if v.Name == name {
			client = v
			break
		}
	}

	if len(client.Name) == 0 {
		err = errors.New(fmt.Sprintf("invalid client name"))
	}

	return
}
