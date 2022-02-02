package config

// main configuration structure
type AppConfig struct {
	DataDir   string
	Capture   *Capture
	Upload    *Upload
	Clipboard *Clipboard
}

type Upload struct {
	Interface       string
	CommandTemplate []string `yaml:"command_template"`
	Bucket          string   `yaml:"bucket"`
	AccessKey       string   `yaml:"access_key"`
	SecretKey       string   `yaml:"secret_key"`
}
type Capture struct {
	Interface string
	// "import"
	Command   string `yaml:"command"`
	KeepLocal bool `yaml:"keep_local"`
}

type Clipboard struct {
	Interface string
	Command   string `yaml:"command"`
}
