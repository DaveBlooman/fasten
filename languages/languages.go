package languages

type Install struct {
	Command    string `yaml:"command"`
	PreCommand string `yaml:"pre_command"`
}
