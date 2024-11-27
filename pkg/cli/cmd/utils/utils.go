package utils

import "github.com/manifoldco/promptui"

func GetInputString(label string) (string, error) {
	validate := func(input string) error {
		return nil
	}
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}
	return prompt.Run()
}
