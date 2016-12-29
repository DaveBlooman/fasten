package msg

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	res, err := PromptUntil([]string{"aws"})
	if err != nil {
		panic(err)
	}

	if res == "aws" || res == "amazon" {
		return "aws"
	}

	return "invalid cloud provider"
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
