package languages

type Install struct {
	Command    string `yaml:"command"`
	PreCommand string `yaml:"pre_command"`
}

type OperatingSystem struct {
	Amzlinux struct {
		Home string `yaml:"home"`
	} `yaml:"amzlinux"`
	Ubuntu1604 struct {
		Home string `yaml:"home"`
	} `yaml:"ubuntu1604"`
}
