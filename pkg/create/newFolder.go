package create

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
)

var templates = &promptui.PromptTemplates{
	Prompt:  "{{ . }} ",
	Valid:   "{{ . | green }} ",
	Invalid: "{{ . | red }} ",
	Success: "{{ . | bold }} ",
}

func CreateNewFolder() error {
	prompt := promptui.Select{
		Label: "Do you need to create new folder for your project?",
		Items: []string{"Yes", "No"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Println("Hehe I got an error: ", err)
	}

	if result == "Yes" {
		fmt.Println("Creating new folder for the project. Please provide a valid name to this project.")
		projectNamePrompt := promptui.Prompt{
			Label:     "Please provide project name.",
			Templates: templates,
			Validate:  validateProjectName,
		}

		result, err := projectNamePrompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

		fmt.Printf("You answered %s\n", result)
	} else {
		fmt.Println("Moving forward without create a new folder.")
	}
	return err
}

func validateProjectName(input string) error {
	if input == "" {
		return errors.New("Project name can't be empty.")
	}

	if len(input) < 4 {
		return errors.New("Project name should be of at-least 4 characters.")
	}

	return nil
}
