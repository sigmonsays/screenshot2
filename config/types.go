package config

// main configuration structure
type AppConfig struct {
	DataDir string
	Upload  *Upload
}

type Upload struct {
	CommandTemplate []string `yaml:"command_template"`
}
