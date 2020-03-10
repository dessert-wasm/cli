package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2/core"
	"github.com/Netflix/go-expect"
)

// BINARY of Dessert
const BINARY = "./dessert-cli"

func fatalError(err error) {
	if err != nil {
		log.Panicf("Got fatal error, aborting: [%v]\n", err)
	}
}

func getConsole() *expect.Console {
	console, err := expect.NewConsole(expect.WithStdout(os.Stdout))
	fatalError(err)
	return console
}

func getCommand(console *expect.Console, name string, arg string) *exec.Cmd {
	cmd := exec.Command(name, arg)
	cmd.Stdin = console.Tty()
	cmd.Stdout = console.Tty()
	cmd.Stderr = console.Tty()
	return cmd
}

func runCommand(cmd *exec.Cmd) {
	err := cmd.Start()
	fatalError(err)
}

func waitCommand(cmd *exec.Cmd) {
	err := cmd.Wait()
	fatalError(err)
}

func testWrapper(name string, arg string, f func(*expect.Console)) {
	console := getConsole()
	defer console.Close()

	cmd := getCommand(console, name, arg)

	runCommand(cmd)
	f(console)
	go console.ExpectEOF()
	waitCommand(cmd)
}

func main() {
	core.DisableColor = true
	//testWrapper(BINARY, "init", dessertInitConnector)
	//testWrapper(BINARY, "init", dessertInitCore)
	testWrapper(BINARY, "login", dessertLogin)
}
