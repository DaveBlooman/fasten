package command

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/DaveBlooman/fasten/msg"
)

type AppStack struct {
	Cloud        string `yaml:"cloud"`
	Applications []Application
}

type Application struct {
	Lang string `yaml:"language"`
	Name string `yaml:"name"`
}

type InitCommand struct {
	Meta
}

func (c *InitCommand) Run(args []string) int {
	var stack AppStack
	fmt.Println("what kind of cloud are you interested in")
	cloud := msg.PromptCloud()
	stack.Cloud = cloud

	fmt.Println("How mang applications are you interested in running")
	applications := msg.PromptApps()
	apps := createAppStack(applications)
	stack.Applications = apps

	res, err := yaml.Marshal(stack)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(string(res))

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

func createAppStack(number int) []Application {
	languages := []string{"ruby", "nodejs", "python", "java", "golang", "go", "rust"}
	var apps []Application

	for i := 0; number > 0; i++ {
		fmt.Println("What is language of application " + strconv.Itoa(i+1))

		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("What is the name of application " + strconv.Itoa(i+1))

		appNameInput := bufio.NewReader(os.Stdin)
		appName, err := appNameInput.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		for _, c := range languages {

			if strings.EqualFold(c, strings.TrimSpace(text)) {
				app := Application{
					Lang: text,
					Name: appName,
				}
				apps = append(apps, app)
			}
		}

		number = number - 1
	}
	return apps
}
