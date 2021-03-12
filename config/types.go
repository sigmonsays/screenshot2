package config

// main configuration structure
type AppConfig struct {
	DataDir string
	Capture *Capture
	Upload  *Upload
}

type Upload struct {
	Interface       string
	CommandTemplate []string `yaml:"command_template"`
	Bucket          string   `yaml:"bucket"`
	AccessKey       string   `yaml:"access_key"`
	SecretKey       string   `yaml:"secret_key"`
}
type Capture struct {
	// "import"
	Command string `yaml:"command"`
}
