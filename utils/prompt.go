package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func GetPathFromUserPrompt() string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New("You must enter a path")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Enter the path where you want to start renaming",
		Validate: validate,
		Templates: &promptui.PromptTemplates{
			Prompt:  "{{ . }}: ",
			Valid:   "{{ . | green }}: ",
			Invalid: "{{ . | blue }}: ",
			Success: "{{ . | bold }}: ",
		},
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Looking in path: %s\n", result)

	return result
}

func GetConfirmationFromUserPrompt() bool {
	items := []string{"Yes!", "No!"}

	prompt := promptui.Select{
		Label: "Are you sure you want to rename the files?",
		Items: items,
		Templates: &promptui.SelectTemplates{
			Selected: `{{ "✔" | green | bold }} {{ "Recipe" | bold }}: {{ .Title | cyan }}`,
		},
	}
	idx, _, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return idx == 0
}

func PrintList(initialList []string, newList []string) {
	fmt.Println()
	for index, item := range initialList {
		fmt.Printf("📄 - %s 👉 %s \n", item, newList[index])
	}
}
