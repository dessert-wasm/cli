package main

import (
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/Netflix/go-expect"
)

func dessertInitConnector(console *expect.Console) {
	console.SendLine(string(terminal.KeyArrowDown))
	console.SendLine(string(terminal.KeyEnter))
	_, err := console.ExpectString("making your folder dessert-ready")
	fatalError(err)
	_, err = console.ExpectString("project set to 'connector'")
	fatalError(err)
	_, err = console.ExpectString("success, happy dessert !")
	fatalError(err)
}

func dessertInitCore(console *expect.Console) {
	console.SendLine(string(terminal.KeyEnter))
	_, err := console.ExpectString("making your folder dessert-ready")
	fatalError(err)
	_, err = console.ExpectString("project set to 'core'")
	fatalError(err)
	_, err = console.ExpectString("success, happy dessert !")
	fatalError(err)
}

func dessertLogin(console *expect.Console) {
}
