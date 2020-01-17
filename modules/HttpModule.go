package modules

import (
	"strconv"

	"github.com/arashrasoulzadeh/serto.git/functions"
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
			free_port := functions.FreePort()
			serve(free_port.Port)
			break
		}
	}
}

/**
static serve current directory
*/
func serve(port int) {
	functions.NoJsonSupport()
	port_string := functions.GetArgOrDefault(4, strconv.Itoa(port))
	port, err := strconv.Atoi(port_string)
	if err != nil {
		functions.ErrorAndDie("invalid port ")
	}
	functions.ServeStaticHttp(port)
}

/**
get a free port
*/
func freePort() {
	port := functions.FreePort()
	if functions.IsJsonOutput() {
		functions.PrettyPrint(port)
	} else {
		functions.Verbose("open port : " + strconv.Itoa(port.Port))
	}
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
