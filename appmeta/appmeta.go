package appmeta

type AppStack struct {
	IP           string `yaml:"ip_address"`
	SSHUser      string `yaml:"ssh_user"`
	OS           string `yaml:"operating_system"`
	KeyPair      string `yaml:"key_pair"`
	InstallDir   string `yaml:"install_dir"`
	Applications []Application
}

type Application struct {
	Lang       string `yaml:"language"`
	Name       string `yaml:"name"`
	Path       string `yaml:"path"`
	PreCommand string `yaml:"pre_command"`
	RunCommand string `yaml:"run_command"`
}
