package utils

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func GetInputString(label string, required bool) (string, error) {
	var validate func(string) error
	if required {
		validate = func(input string) error {
			if input == "" {
				return fmt.Errorf("input cannot be null")
			}
			return nil
		}
	} else {
		validate = func(input string) error {
			return nil
		}
	}
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}
	return prompt.Run()
}
