package main

import (
	"os"

	"github.com/arashrasoulzadeh/serto.git/functions"
	"github.com/arashrasoulzadeh/serto.git/modules"
)

func command(args []string) {
	has_command := functions.DieIfEqual(1, args, "Please Enter the command.")

	if has_command {
		command := args[1]
		switch command {
		case "info":
			loadInfoModule()
		case "http":
			loadHttpModule()
			break
		case "os":
			loadOsModule()
			break
		default:
			functions.ErrorAndDie("Invalid Module " + command)
		}
	}
}

/**
load http module
*/
func loadHttpModule() {
	modules.ParseHttpModule(os.Args)
}

/**
load os module
*/
func loadOsModule() {
	modules.ParseOSModule(os.Args)
}

/**
load info module
*/
func loadInfoModule() {
	modules.ParseInfoModule(os.Args)
}

func main() {
	//_os := functions.ClientOS()
	//color.Green("Your operating system is %s", _os)
	command(os.Args)
}
