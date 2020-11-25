package prompt

import "github.com/AlecAivazis/survey/v2"

var loginQ = []*survey.Question{
	{
		Name:      "username",
		Prompt:    &survey.Input{Message: promptLoginUsername},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "password",
		Prompt: &survey.Password{
			Message: promptLoginPassword,
		},
	},
}

// LoginA is the structure used for login questions
type LoginA struct {
	Username string
	Password string
}

// LoginPrompt is a custom login prompt for Dessert
func LoginPrompt(loginA *LoginA) error {
	return survey.Ask(loginQ, loginA, survey.WithIcons(customPrompt))
}
