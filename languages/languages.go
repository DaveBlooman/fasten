package languages

type Install struct {
	Command    string `yaml:"command"`
	PreCommand string `yaml:"pre_command"`
	BuildLocal bool   `yaml:"build_local"`
}

type OperatingSystem struct {
	Amzlinux struct {
		Home    string `yaml:"home"`
		SSHUser string `yaml:"ssh_user"`
		Version string `yaml:"version"`
	} `yaml:"amzlinux"`
	Ubuntu1604 struct {
		Home    string `yaml:"home"`
		SSHUser string `yaml:"ssh_user"`
		Version string `yaml:"version"`
	} `yaml:"ubuntu1604"`
}

type OperatingSystemMeta struct {
	Home    string `yaml:"home"`
	SSHUser string `yaml:"ssh_user"`
	Version string `yaml:"version"`
}
