package Nginx

import (
	f "../../Functions"
)

func PhpInfoCommand() []byte {
	return f.ExecuteCommand("php", "-info", "", "", false)
}

func PhpInfo(filter string) {
	f.GrepOutput(string(PhpInfoCommand()), filter)
}

func PhpConfig() {
	f.GrepOutput(string(PhpInfoCommand()), "php.ini")
}

func Parse(args []string) {
	f.SetProcessor("PHP")
	f.Verbose("Loaded")
	f.DieIfEqual(2, args, "Please Enter sub command.")
	switch args[2] {
	case "info":
		filter := ""
		if len(args) == 4 {
			filter = args[3]
		}
		PhpInfo(filter)
		break
	case "config":
		PhpConfig()
		break
	case "ini":
		PhpConfig()
		break
	}

}
