package config

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

func Watch() {
	// ...
}
