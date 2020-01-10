package main

import (
	"github.com/arashrasoulzadeh/serto.git/functions"
	"github.com/arashrasoulzadeh/serto.git/modules"
	"os"
)

func command(args []string) {
	has_command := functions.DieIfEqual(1, args, "Please Enter the command.")

	if has_command {
		command := args[1]
		switch command {
		case "http":
			loadHttpModule()
			break
		default:
			functions.ErrorAndDie("Invalid Module " + command)
		}
	}
}

func loadHttpModule() {
	modules.ParseHttpModule(os.Args)
}
func main() {
	//_os := functions.ClientOS()
	//color.Green("Your operating system is %s", _os)
	command(os.Args)
}
