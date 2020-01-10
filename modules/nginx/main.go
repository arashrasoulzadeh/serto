package Nginx

import (
	f "../../Functions"
	"fmt"
)

func ConfigFile() {
	path := "/etc/nginx/nginx.conf"
	if f.FileExists(path) {
		f.Verbose(fmt.Sprint("Nginx Conf Found!  ", path))
		//f.OpenFileForEdit(path)
	} else {
		f.ErrorAndDie("Cant find the config file in default location")
	}
}

func Parse(args []string) {
	f.SetProcessor("NGINX")
	f.DieIfEqual(2, args, "Please Enter sub command.")
	switch args[2] {
	case "config":
		ConfigFile()
	}
}
