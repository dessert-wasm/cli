package prompt

import "github.com/AlecAivazis/survey/v2"

func customPrompt(icons *survey.IconSet) {
	icons.Question.Text = ">"
}
