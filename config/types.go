package config

// main configuration structure
type AppConfig struct {
	DataDir      string
	WriteUrlFile bool
	Capture      *Capture
	Upload       *Upload
	Clipboard    *Clipboard
	PostProcess  []*PostProcess
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
	Command   string   `yaml:"command"`
	Cmdline   []string `yaml:"cmdline"`
	KeepLocal bool     `yaml:"keep_local"`
}

type Clipboard struct {
	Interface string
	Command   string `yaml:"command"`
}
type PostProcess struct {
	Command string `yaml:"command"`
}
