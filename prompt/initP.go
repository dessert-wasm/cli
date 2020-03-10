package prompt

import "github.com/AlecAivazis/survey/v2"

var initQ = []*survey.Question{
	{
		Name: "type",
		Prompt: &survey.Select{
			Message: promptProjectKind,
			Options: []string{"core", "connector"},
			Default: "core",
		},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
}

// InitA is the structure used for init Questions
type InitA struct {
	Type string
}

// InitPrompt for Dessert
func InitPrompt(initA *InitA) error {
	return survey.Ask(initQ, initA, survey.WithIcons(customPrompt))
}
