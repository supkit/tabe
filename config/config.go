package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

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
	Client []struct {
		Name   string `yaml:"name"`
		Target string `yaml:"target"`
	} `yaml:"client"`
}

// Watch load config
func Watch() (config Config, err error) {
	// ...
	data, err := os.ReadFile("./tabe.yaml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return
	}

	return config, err
}
