package msg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v1"

	"github.com/DaveBlooman/fasten/files"
	"github.com/DaveBlooman/fasten/languages"
	"github.com/DaveBlooman/fasten/output"
	"github.com/fatih/color"
)

func PromptUntilYorN() bool {
	res, err := PromptUntil([]string{"y", "yes", "n", "no"})
	if err != nil {
		panic(err)
	}

	if res == "y" || res == "yes" {
		return true
	}

	return false
}

func PromptCloud() string {
	items := []string{"AWS", "GCP", "Azure"}

	output.Selection(items)

	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			output.Error("cannot read IP Address")
		}
		selection, err := strconv.Atoi(strings.TrimSpace(text))
		if err != nil {
			fmt.Println("please try again")
			continue
		}
		if selection > len(items) {
			fmt.Println("please select from the list above")
			continue
		}
		return items[selection-1]
	}

}

func PromptIP() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		output.Error("cannot read IP Address")
	}
	return strings.TrimSpace(text)
}

func PromptOS() (string, string) {
	definition, err := files.Asset("libraries/definition.yaml")

	var osType languages.OperatingSystem
	err = yaml.Unmarshal([]byte(definition), &osType)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	items := []string{"ubuntu1604", "Amazon Linux"}

	output.Selection(items)

	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			output.Error("OS selection not valid")
		}
		selection, err := strconv.Atoi(strings.TrimSpace(text))
		if err != nil {
			fmt.Println("please try again")
			continue
		}
		if selection > len(items) {
			fmt.Println("please select from the list above")
			continue
		}
		if items[selection-1] == "ubuntu1604" {
			return "ubuntu1604", osType.Ubuntu1604.Home
		} else {
			return "amzlinux", osType.Amzlinux.Home
		}
	}

}

func PromptApps() int {
	reader := bufio.NewReader(os.Stdin)

	for {

		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("enter a number")
		}

		text = strings.Replace(text, "\n", "", -1)
		i, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("enter a number between 1 and 10")
		}

		return i
	}
}

func PromptUntil(opts []string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		for _, c := range opts {
			if strings.EqualFold(c, strings.TrimSpace(text)) {
				return c, nil
			}
		}
	}
}

func PromptKeyPair() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		output.Error("cannot read path")
	}
	return strings.TrimSpace(text)
}

func PromptLang() string {

	items := []string{"ruby", "nodejs", "python", "java", "golang", "go", "rust"}

	output.Selection(items)

	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			output.Error("Language selection not valid")
		}
		selection, err := strconv.Atoi(strings.TrimSpace(text))
		if err != nil {
			fmt.Println("please try again")
			continue
		}
		if selection > len(items) {
			fmt.Println("please select from the list above")
			continue
		}
		return items[selection-1]
	}

}

func changeColor(text string, code color.Attribute) string {
	c := color.New(code).SprintFunc()
	return c(text)
}
