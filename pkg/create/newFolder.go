package create

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/manifoldco/promptui"
)

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
		return createNewDir()
	} else {
		fmt.Println("Moving forward without create a new folder.")
	}
	return err
}

func createNewDir() error {
	projectNamePrompt := promptui.Prompt{
		Label:    "Please provide project name.",
		Validate: validateProjectName,
	}

	result, err := projectNamePrompt.Run()

	if err != nil {
		fmt.Printf("Invalid project name %v\n", err)
		return err
	}

	fmt.Printf("Creating new folder with the name %s ", result)

	cmd := exec.Command("mkdir", result)

	_, err = cmd.Output()

	if err != nil {
		fmt.Println("Error in creating the folder")
		return err
	}

	return nil
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
