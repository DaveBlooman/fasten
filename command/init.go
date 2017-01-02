package command

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/DaveBlooman/fasten/appmeta"
	"github.com/DaveBlooman/fasten/msg"
	"github.com/DaveBlooman/fasten/output"
)

var apps []appmeta.Application

type InitCommand struct {
	Meta
}

func (c *InitCommand) Run(args []string) int {
	var stack appmeta.AppStack
	output.Standard("What kind of cloud are you interested in?  AWS, GCP or Azure")
	cloud := msg.PromptCloud()
	stack.Cloud = cloud

	output.Standard("What is the full path to your key pair")
	keypair := msg.PromptKeyPair()
	stack.KeyPair = keypair

	output.Standard("What is your server IP address?")
	ipAddress := msg.PromptIP()
	stack.IP = ipAddress

	output.Standard("What Operating System are you using ?")
	operatingSystem := msg.PromptOS()
	stack.OS = operatingSystem

	output.Standard("How many applications are you interested in running?")
	applications := msg.PromptApps()
	apps := createAppStack(applications)
	stack.Applications = apps

	res, err := yaml.Marshal(stack)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	file, err := os.Create("fasten.yaml")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, string(res))

	return 0
}

func (c *InitCommand) Synopsis() string {
	return ""
}

func (c *InitCommand) Help() string {
	helpText := `Init help

`
	return strings.TrimSpace(helpText)
}

func createAppStack(number int) []appmeta.Application {
	languages := []string{"ruby", "nodejs", "python", "java", "golang", "go", "rust"}

	for i := 0; number > 0; i++ {
		m := map[string]string{"What is language of application ": strconv.Itoa(i + 1)}
		output.Banner("Language", m)

		language := strings.TrimSpace(msg.PromptLang())

		m = map[string]string{"What is the name of application ": strconv.Itoa(i + 1)}
		output.Banner("Name", m)

		appNameInput := bufio.NewReader(os.Stdin)
		appName, err := appNameInput.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		appName = strings.TrimSpace(appName)

		m = map[string]string{"What is the full path of": strings.TrimSpace(appName)}
		output.Banner("Path to Application", m)

		appPathInput := bufio.NewReader(os.Stdin)
		appPath, err := appPathInput.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		appPath = strings.TrimSpace(appPath)

		for _, c := range languages {

			if strings.EqualFold(c, language) {
				app := appmeta.Application{
					Lang: language,
					Name: appName,
					Path: appPath,
				}
				apps = append(apps, app)
			}
		}

		number = number - 1
	}
	return apps
}
