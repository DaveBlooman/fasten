package output

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func FastenText() string {
	c := color.New(color.FgHiRed).SprintFunc()
	return c("[FASTEN]:")
}

// Banner outputs a summary of the command & flags
func Banner(name string, flags map[string]string) {
	fmt.Printf("%s : %s:\"%s\"\n", FastenText(), changeColor("Type", color.FgBlue), changeColor(name, color.FgWhite))
	for key, val := range flags {
		fmt.Printf("%s : %s: %s \n", FastenText(), changeColor(strings.Title(key), color.FgBlue), changeColor(val, color.FgWhite))
	}
}

// Standard output
func Standard(summary string) {
	fmt.Printf("%s : %s \n", FastenText(), changeColor(summary, color.FgBlue))
}

// Success output
func Success(summary string) {
	fmt.Printf("%s : %s \n", FastenText(), changeColor(summary, color.FgGreen))
}

// Install output
func Install(summary string) {
	fmt.Printf("%s : %s \n", FastenText(), changeColor(summary, color.FgWhite))
}

// Error ends the running process with a red error message
func Error(text string) {
	fmt.Printf("%s : %s \n", FastenText(), changeColor(text, color.FgRed))
	os.Exit(1)
}

// Selection Output
func Selection(items []string) {
	for i, choice := range items {
		fmt.Printf("%s : %v. %s: \n", FastenText(), i+1, changeColor(choice, color.FgWhite))
	}
}

func changeColor(text string, code color.Attribute) string {
	c := color.New(code).SprintFunc()
	return c(text)
}
