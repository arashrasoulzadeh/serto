package modules

import (
	"github.com/arashrasoulzadeh/serto.git/functions"
)

/**
init the http module
*/
func ParseInfoModule(args []string) {
	functions.SetProcessor("Info")
	has_command := functions.DieIfEqual(2, args, "Please Enter Sub Command.")
	if has_command {
		command := args[2]
		switch command {
		case "ip":
			MyPublicIp()
			break
		}

	}
}

/**
show your public ip
*/
func MyPublicIp() {
	geo := functions.IpGeolocationStruct{}
	public_ip := functions.GetPublicIP()
	ip := functions.GetArgOrDefault(4, public_ip.IP)
	geo = functions.IpGeolocation(ip)
	if functions.IsJsonOutput() {
		functions.PrettyPrint(geo)
	} else {
		functions.Verbose("Country : " + geo.COUNTRY_NAME)
		functions.Verbose("Code    : " + geo.COUNTRY_CODE)
		functions.Verbose("Lat     : " + geo.LAT)
		functions.Verbose("Lon     : " + geo.LON)
	}
}
