package modules

import (
	"github.com/arashrasoulzadeh/serto.git/functions"
	"strconv"
)

/**
init the http module
*/
func ParseHttpModule(args []string) {
	functions.SetProcessor("Http")
	has_command := functions.DieIfEqual(2, args, "Please Enter Sub Command.")
	if has_command {
		command := args[2]
		switch command {
		case "ports":
			openPorts()
			break
		case "free":
			freePort()
			break
		case "serve":
			serve(functions.FreePort())
			break
		}
	}
}

/**
static serve current directory
*/
func serve(port int) {
	functions.ServeStaticHttp(port)
}

/**
get a free port
*/
func freePort() {
	functions.Verbose("open port : " + strconv.Itoa(functions.FreePort()))
}

/**
open os http ports
*/
func openPorts() {
	ports := []string{"23", "80", "443", "8080", "9090"}
	for i := 0; i < len(ports); i++ {
		functions.Verbose("port " + ports[i] + " is " + strconv.FormatBool(functions.CheckPort(ports[i])))
	}
}
