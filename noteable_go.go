package main

import "github.com/nandarimansyah/noteable_go/commands"

func main() {
	command := commands.NewCommandEngine()
	command.Run()
}
