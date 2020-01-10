package main

import (
	f "./Functions"
	nginx "./modules/nginx"
	php "./modules/php"
	"github.com/fatih/color"
	"os"
)

func help() {
f.Verbose("test")
}

func main() {
	_os := f.ClientOS()
	if _os == "mac" {
		color.Green("Your operating system is %s", _os)
	}
	command(os.Args)
}
func command(args []string) {
	has_command := f.DieIfEqual(1, args, "Please Enter the command.")
	if has_command {
		switch args[1] {
		case "nginx":
			nginx.Parse(args)
			break
		case "php":
			php.Parse(args)
			break

		}
	} else {
		help()
	}
}
